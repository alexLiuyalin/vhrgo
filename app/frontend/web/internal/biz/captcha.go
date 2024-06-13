package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	apiweb "vhrgo/api/frontend/web/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wenlng/go-captcha/captcha"
)

type CaptchaRepo interface {
	// 缓存验证码
	CacheCaptcha(ctx context.Context, cid string, captcha []byte) error
	// 获取验证码
	GetCacheCaptcha(ctx context.Context, cid string) ([]byte, error)
	// 缓存成功验证码
	CacheSuccessCaptcha(ctx context.Context, cid string) error
	// 获取成功验证码
	GetSuccessCaptcha(ctx context.Context, cid string) error
}
type CaptchaUseCase struct {
	repo CaptchaRepo
	log  *log.Helper
}

func NewCaptchaUseCase(repo CaptchaRepo, logger log.Logger) *CaptchaUseCase {
	return &CaptchaUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "web", "biz/captcha")),
	}
}

func (uc *CaptchaUseCase) GetCaptcha(ctx context.Context) (*apiweb.GetCaptchaReply, error) {
	capt := captcha.GetCaptcha()
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return &apiweb.GetCaptchaReply{}, err
	}
	// uc.writeCache(dots, key) // 缓存在本地

	bt, err := json.Marshal(dots)
	if err != nil {
		return nil, err
	}

	err = uc.repo.CacheCaptcha(ctx, key, bt)
	if err != nil {
		return &apiweb.GetCaptchaReply{}, err
	}
	return &apiweb.GetCaptchaReply{
		B64:  b64,
		Tb64: tb64,
		Cid:  key,
	}, nil
}

func (uc *CaptchaUseCase) VerifyCaptcha(ctx context.Context,
	req *apiweb.VerifyCaptcha) error {
	cacheData, err := uc.repo.GetCacheCaptcha(ctx, req.Cid)
	if err != nil {
		return nil
	}
	src := strings.Split(req.VerifyValue, ",")
	var dct map[int]captcha.CharDot
	if err := json.Unmarshal(cacheData, &dct); err != nil {
		return nil
	}
	chkRet := false
	if (len(dct) * 2) == len(src) {
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)

			// 检测点位置
			// chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))

			// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
			// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
			chkRet = captcha.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5)
			if !chkRet {
				break
			}
		}
	}
	if !chkRet {
		return errors.New("验证不通过")
	}

	// 验证通过
	err = uc.repo.CacheSuccessCaptcha(ctx, req.Cid)
	if err != nil {
		return err
	}
	return nil
}
