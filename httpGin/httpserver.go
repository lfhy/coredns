package httpGin

import (
	"dns/controller"
	dhtml "dns/html"
	"dns/resource"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StartHttp(port string) {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	dhtml.WriteHTMLTemplate()
	r.LoadHTMLGlob(os.TempDir() + "/views/*")
	// os.RemoveAll(os.TempDir() + "/views")
	r.StaticFS("/static", resource.GetStaticFS())
	admin := r.Group("admin")
	admin.Use(controller.AuthRequired())
	{
		admin.GET("", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/admin/dns")
		})
		admin.GET("dns", controller.Dnslist2())
		//admin.GET("dns2",controller.Dnslist2())
		admin.GET("deldns", controller.DelDns())
		admin.GET("etcdlist", controller.DnsApiList())
		admin.GET("dnsedit", controller.DnsEditGet())
		admin.POST("dnsedit", controller.DnsEditPost())
		//admin.GET("ws",controller.WsHandler())
	}
	r.GET("ws", controller.WsHandler())
	r.GET("/login", controller.Login_get())
	r.POST("/login", controller.Login_post())
	r.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin/dns")
	})
	r.Run(":" + port)
}
