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

	v1:=r.Group("/v1")
	{
		v1.GET("/", func(c *sGin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *sGin.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *sGin.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *sGin.Context) {
			c.JSON(http.StatusOK, sGin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

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
