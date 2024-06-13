package biz

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	. "vhrgo/pkg/common"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	apiweb "vhrgo/api/frontend/web/v1"
)

type AuthRepo interface {
	GetIdByCid(ctx context.Context, cid string) (string, error)
	SetIdByCid(ctx context.Context, cid, id string) error
	SetToken(ctx context.Context, id, token string) error
	GetToken(ctx context.Context, id string) (string, error)
}

type AuthUseCase struct {
	userRepo    UserRepo
	authRepo    AuthRepo
	captchaRepo CaptchaRepo
	log         *log.Helper
	securityKey string
}

func NewAuthUseCase(logger log.Logger,
	userRepo UserRepo,
	authRepo AuthRepo,
	captchaRepo CaptchaRepo,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo:    authRepo,
		userRepo:    userRepo,
		captchaRepo: captchaRepo,
		log:         log.NewHelper(log.With(logger, "module", "biz/login")),
		// securityKey: conf.SecurityKey,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, req *apiweb.LoginRequest) (string, error) {
	cid := ctx.Value(ContextCidKey{}).(string)
	// 是否验证过验证码
	err := uc.captchaRepo.GetSuccessCaptcha(ctx, cid)
	if err != nil {
		return "", err
	}
	d, _ := json.Marshal("123456")
	tokenData := jwt.New(jwt.SigningMethodHS256) // 访问 https://golang-jwt.github.io/jwt/usage/create/
	token, err := tokenData.SignedString(d)
	if err != nil {
		uc.log.Error("token", err)
	}
	user, err := uc.userRepo.GetPassword(ctx, req.Account)
	if err != nil {
		return "", errors.New("账号错误")
	}
	if req.Password != user.Password.String {
		return "", errors.New("密码错误")
	}

	err = uc.authRepo.SetIdByCid(ctx, cid, strconv.Itoa(int(user.ID)))
	if err != nil {
		return "", err
	}
	err = uc.authRepo.SetToken(ctx, strconv.Itoa(int(user.ID)), token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *AuthUseCase) Logout(ctx context.Context) error {
	// cid := ctx.Value(constants.ContextCidKey{}).(string)
	// result, err := uc.repo.Logout(ctx, &csV1.LogoutRequest{
	// 	Cid: cid,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &v1.LogoutReply{
	// 	Code:    result.Code,
	// 	Message: result.Message,
	// }, nil
	return nil
}
