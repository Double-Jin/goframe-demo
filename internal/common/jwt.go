package common

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe/internal/consts"
	"goframe/internal/model/entity"
	"time"
)

// jwt鉴权
type JWT struct{}

var Jwt = new(JWT)

func (component *JWT) GenerateAdminLoginToken(ctx context.Context, user *entity.Admin, isRefresh bool) (interface{}, error) {

	jwtVersion, _ := g.Cfg().Get(ctx, "jwt.version", "1.0")
	jwtSign, _ := g.Cfg().Get(ctx, "jwt.sign", "jin")
	Expires, _ := g.Cfg().Get(ctx, "jwt.expires", "10000")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.Id,
		"login_name":  user.LoginName,
		"admin_name":    user.AdminName,
		"admin_img":      user.AdminImg,
		"password":       user.Password,
		"sex":      user.Sex,
		"id_card":   user.IdCard,
		"phone":     user.Phone,
		"is_refresh":  isRefresh,
		"jwt_version": jwtVersion.String(),
	})

	tokenString, err := token.SignedString(jwtSign.Bytes())
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}

	tokenStringMd5 := gmd5.MustEncryptString(tokenString)

	// TODO  绑定登录token
	cache := Cache.New()
	key := consts.RedisJwtToken + tokenStringMd5

	// TODO  将有效期转为持续时间，单位：秒
	expires, _ := time.ParseDuration(fmt.Sprintf("+%vs", Expires))

	err = cache.Set(ctx, key, tokenString, expires)
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}
	_ = cache.Set(ctx, consts.RedisJwtUserBind+":"+gconv.String(user.Id), key, expires)

	return tokenString, err
}

func (component *JWT) ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	if tokenString == "" {
		err := gerror.New("token 为空")
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if token == nil {
		err := gerror.New("token不存在")
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}


func (component *JWT) GetAuthorization(r *ghttp.Request) string {

	// TODO  默认从请求头获取
	var authorization = r.Header.Get("Authorization")

	// TODO  如果请求头不存在则从get参数获取
	if authorization == "" {
		return r.Get("authorization").String()
	}

	return gstr.Replace(authorization, "Bearer ", "")
}


func (component *JWT) Layout(adminUserId int, tokenString string) {
	if tokenString == "" {
		return
	}
	//g.Redis().Do("HDEL", "VerifyLoginToken", gmd5.MustEncryptString(tokenString))
	//// 删除
	//g.Redis().Do("HDEL", "VerifyLoginTokenAdminUserId", adminUserId)
}
