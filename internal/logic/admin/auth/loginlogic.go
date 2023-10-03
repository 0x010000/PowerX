package auth

import (
	"PowerX/internal/model/option"
	"PowerX/internal/types/errorx"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"strings"
	"time"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	if err != nil {
		panic(err)
	}
	opt := option.EmployeeLoginOption{
		Account:     req.UserName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	employee, err := l.svcCtx.PowerX.Organization.FindOneEmployeeByLoginOption(l.ctx, &opt)
	if err != nil {
		return nil, errorx.WithCause(errorx.ErrBadRequest, "账户或密码错误")
	}

	if !l.svcCtx.PowerX.Organization.VerifyPassword(employee.Password, req.Password) {
		return nil, errorx.WithCause(errorx.ErrBadRequest, "账户或密码错误")
	}

	roles, _ := l.svcCtx.PowerX.AdminAuthorization.Casbin.GetRolesForUser(employee.Account)
	var app []string
	if employee.Position != nil {
		app = strings.Split(employee.Position.App, `,`)
	}
	claims := types.TokenClaims{
		UID:     employee.Id,
		Account: employee.Account,
		App:     app,
		Roles:   roles,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "powerx",
			Subject:   employee.Account,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(l.svcCtx.Config.JWT.JWTSecret))
	if err != nil {
		return nil, errors.Wrap(err, "sign token failed")
	}

	return &types.LoginReply{
		Token: signedToken,
	}, nil
}
