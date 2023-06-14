package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
	"go-web/pkg/app"
	"go-web/pkg/config"
	"go-web/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type Jwt struct {
	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte
	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

type JwtCustomClaims struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.StandardClaims
}

func NewJwt() *Jwt {
	return &Jwt{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.Max_refresh_time")) * time.Minute,
	}
}

// IssueToken 生成 token
func (j *Jwt) IssueToken(userId, userName string) string {
	// 构造用户 claims 信息
	expireTime := j.expireTime()
	claims := JwtCustomClaims{
		userId,
		userName,
		expireTime,
		jwtpkg.StandardClaims{
			NotBefore: app.TImeNowInTimezone().Unix(), //签名生效时间
			IssuedAt:  app.TImeNowInTimezone().Unix(), // 首次签名时间 后续刷新token 不会更新
			ExpiresAt: expireTime,                     // 签名过期时间
			Issuer:    config.GetString("app.name"),   //签名颁发者
		},
	}

	token, err := j.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// ParserToken 解析token 中间件调用
func (j *Jwt) ParserToken(c *gin.Context) (*JwtCustomClaims, error) {

	// 解析token数据
	token, err := j.getToken(c)
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrHeaderMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, err
	}

	// 将token中的claims信息解析出来和 JwtCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 刷新token
func (j *Jwt) RefreshToken(c *gin.Context) (string, error) {
	//
	token, err := j.getToken(c)
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		// 满足 refresh 的条件，只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 解析JwtCustomClaims的数据
	claims := token.Claims.(*JwtCustomClaims)

	// 检查是否过了最大允许刷新时间
	x := app.TImeNowInTimezone().Add(-j.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = j.expireTime()
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// getToken 从 header 中获取 token 值并解析
func (j *Jwt) getToken(c *gin.Context) (*jwtpkg.Token, error) {
	tokenString, err := j.getTokenFromHerder(c)
	if err != nil {
		return nil, err
	}

	return j.parseTokenString(tokenString)
}

// getTokenFromHerder 获取token的内容
func (j *Jwt) getTokenFromHerder(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, "", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrTokenMalformed
	}
	return parts[1], nil
}

// parseTokenString 使用jwt.ParseWithClaims解析token数据
func (j *Jwt) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

// expireTime 过期时间
func (j *Jwt) expireTime() int64 {
	timeNow := app.TImeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timeNow.Add(expire).Unix()
}

// createToken 创建 token
func (j *Jwt) createToken(claims JwtCustomClaims) (string, error) {
	// 使用 HS256 算法进行 token 生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}
