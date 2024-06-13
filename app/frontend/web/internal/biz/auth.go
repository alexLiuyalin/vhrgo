package biz

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

//	type AuthRepo interface {
//		Login(ctx context.Context, req *csV1.LoginPcRequest) (*csV1.LoginPcReply, error)
//		VerifyToken(ctx context.Context, req *csV1.VerifyTokenRequest) (*csV1.VerifyTokenReply, error)
//		Logout(ctx context.Context, req *csV1.LogoutRequest) (*csV1.LogoutReply, error)
//	}
type AuthUseCase struct {
	hrRepo      HrRepo
	log         *log.Helper
	securityKey string
}

func NewAuthUseCase(logger log.Logger,
	hrRepo HrRepo,
) *AuthUseCase {
	return &AuthUseCase{
		hrRepo: hrRepo,
		log:    log.NewHelper(log.With(logger, "module", "biz/login")),
		// securityKey: conf.SecurityKey,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, req *apiweb.LoginRequest) (string, model.HR, error) {
	tokenData := jwt.New(jwt.SigningMethodES512)
	token, _ := tokenData.SignedString("123456")

	res, err := uc.hrRepo.Get(ctx, req.Account)
	if req.Password != res.Password.String {
		return "", model.HR{}, err
	}
	return token, res, nil
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
