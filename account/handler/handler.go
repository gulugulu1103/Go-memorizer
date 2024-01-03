package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Handler struct holds required services for the handler to function
// Handler 结构体包含了 handler 层所需的服务
type Handler struct {
}

// Config will hold repositories that will eventually be injected into this
// handler layer on handler initialization
// Config 会保存 repositories，这些 repositories 最终会在初始化 handler 层的时候被注入到这里
type Config struct {
	R *gin.Engine
	//US *model.UserService
}

// NewHandler initializes the handler with required injections
// NewHandler 用于初始化 handler 层
func NewHandler(c *Config) {
	// 通过 gin.Default() 创建一个默认的路由引擎，使用了 Logger、Recovery 中间件
	h := &Handler{}

	// 路由组: https://www.liwenzhou.com/posts/Go/gin_route/
	g := c.R.Group(os.Getenv("ACCOUNT_API_URL")) // 路由组，可以定义一些公共的前缀

	// 路由组中间件: https://www.liwenzhou.com/posts/Go/gin_middleware/
	g.GET("/me", h.Me) // h.将会调用 Handler 结构体中的 Me 方法
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.GET("/details", h.Details)

	// 路由 GET 请求，第一个参数是路径，第二个参数是处理这个请求的函数
	// 函数的要求是 func(c *gin.Context)，gin.Context 封装了 request 和 response
	// 这里返回一个 JSON，JSON 是一个 map[string]interface{} 的实例
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // gin.H 是 map[string]interface{} 的一种快捷方式
			"message": "hello",
		})
	})

}

// Me handler calls services for getting a user's profile
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's me!",
	})
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's signup!",
	})
}

// Signin handler
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's signin!",
	})
}

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's signout!",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's tokens!",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's image!",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's delete image!",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's details!",
	})
}
