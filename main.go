package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"swkGin/sGin"
	"time"
)
func onlyForV2() sGin.HandlerFunc {
	return func(c *sGin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
type student struct {
	Name string
	Age  int8
}

func main() {
	r:=sGin.New()
	r.Static("/assets", "./static")

	r.Use(sGin.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")


	stu1 := &student{Name: "ben", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.GET("/", func(c *sGin.Context) {
		c.HTML(http.StatusOK,"css.tmpl",nil)

	})
	r.GET("/students", func(c *sGin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", sGin.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *sGin.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", sGin.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	v1:=r.Group("/v1")
	{

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


	r.POST("/login", func(c *sGin.Context) {
		c.JSON(http.StatusOK,sGin.H{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
