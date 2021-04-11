package main

import (
	"fmt"
	"net/http"
	"swkGin/sGin"
)

func main() {
	r:=sGin.New()
	r.GET("/", func(c *sGin.Context) {
		fmt.Fprintf(c.Writer, "<h1>hello user</h1>")
	})
	r.GET("/hello", func(c *sGin.Context) {
		c.String(http.StatusOK,"hello %s you're at %s\n",c.Query("name"),c.Path)
	})
	r.GET("/hello/:name", func(c *sGin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *sGin.Context) {
		c.JSON(http.StatusOK, sGin.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *sGin.Context) {
		c.JSON(http.StatusOK,sGin.H{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
