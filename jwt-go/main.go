package main

import (
	"github.com/gin-gonic/gin"
	"jwt-go/controllers"
	"jwt-go/middleware"
)

func main() {

	//mySigningKey := []byte("AllYourBase")
	//
	//c := MyClaims{
	//	Username: "admin",
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(3 * time.Hour).Unix(), // 表示 JWT 的过期时间（expiration time）
	//		Issuer:    "admin",                              // 表示 JWT 的签发者（issuer）
	//		NotBefore: time.Now().Unix() - 60,               // 表示 JWT 的生效时间（not before time）
	//	},
	//	// 当一个结构体嵌入到另一个结构体中时，嵌入的结构体的名字会成为外层结构体的一个字段名。
	//	// 这种字段名是匿名的，但仍然可以通过嵌入的结构体名来访问其字段和方法。
	//}
	//t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//fmt.Println(t)
	//s, e := t.SignedString(mySigningKey)
	//if e != nil {
	//	fmt.Printf("Couldn't sign token: %s\n", e)
	//}
	//fmt.Println(s)
	//
	//token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return mySigningKey, nil
	//})
	//if err != nil {
	//	fmt.Printf("Couldn't parse token: %s\n", err)
	//}
	//fmt.Println(token)
	//fmt.Println(token.Claims)
	//fmt.Println(reflect.TypeOf(token))        // *jwt.Token
	//fmt.Println(reflect.TypeOf(token.Claims)) // *main.MyClaims

	/*
		r := gin.Default()

		// 公共路由，无需认证
		r.POST("/login", func(c *gin.Context) {
			var json struct {
				Username string `json:"username"`
			}

			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			token, err := middleware.GenerateToken(json.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": token})
		})

		// 受保护的路由，需要 JWT 验证
		protected := r.Group("/protected")
		protected.Use(middleware.JWTMideware())
		{
			protected.GET("/profile", func(c *gin.Context) {
				username, _ := c.Get("username")
				c.JSON(http.StatusOK, gin.H{"username": username})
			})
		}

		r.Run(":8080")
	*/
	r := gin.Default()

	// 公共路由，无需认证
	r.POST("/login", controllers.Login)

	// 刷新令牌路由
	r.POST("/refresh", controllers.RefreshToken)

	// 受保护的路由，需要 JWT 验证
	protected := r.Group("/protected")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/profile", controllers.Profile)
	}

	r.Run(":8080")
}

/*
在 JWT（JSON Web Token）中，时间字段都是以 Unix 时间戳的形式存储的。这是因为 Unix 时间戳是一种通用且精确的时间表示方法，能够跨越不同的系统和编程语言。

什么是 Unix 时间戳？
Unix 时间戳是指从 1970 年 1 月 1 日 00:00:00 UTC（称为 Unix 纪元）开始所经过的秒数。它是一个整数值，可以方便地用于表示和比较时间。

为什么使用 Unix 时间戳？
通用性：Unix 时间戳在所有编程语言和操作系统中都是一致的，可以方便地进行跨系统时间比较和处理。
精确性：Unix 时间戳以秒为单位，非常适合表示精确的时间点。
简单性：与字符串格式的时间表示相比，Unix 时间戳更简单且节省空间，便于存储和传输。
*/
