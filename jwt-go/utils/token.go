package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 包含生成和刷新令牌的逻辑。

var MySigningKey = []byte("AllYourBase")

// MyCustomClaims 自定义声明
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

/*
在 jwt-go 包中，Claims 接口定义了一些方法，任何实现了这些方法的类型都可以作为 JWT 的声明，jwt.StandardClaims 已经实现了 Claims 接口。
通过在 MyClaims 中嵌入 jwt.StandardClaims，MyClaims 结构体自动继承了 jwt.StandardClaims 的所有字段和方法，因此也实现了 Claims 接口。
*/

// GenerateToken 生成 JWT
func GenerateToken(username string) (string, string, error) {
	accessClaims := MyCustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(), // 15分钟有效期
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(MySigningKey)
	if err != nil {
		return "", "", err
	}

	refreshClaims := MyCustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), // 7天有效期
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(MySigningKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

/*
func GenerateToken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
		// 这种简写形式实际上是可以的，但为了代码的可读性和清晰度，建议明确写出字段名。
		// 这不仅有助于代码的自文档化，还可以避免在添加或修改字段时引入错误。
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建了一个 jwt.Token 对象，并将头部和载荷设置为指定的值。
	// - jwt.SigningMethodHS256：指定签名算法为HS256（HMAC SHA256）。
	// - claims：指定JWT的载荷部分，即声明（claims）。
	return token.SignedString(mySigningKey)
}
*/

func RefreshToken(refreshToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})

	if err != nil || !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return "", "", err
	}

	return GenerateToken(claims.Username)
}
