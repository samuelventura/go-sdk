package main

import (
	"embed"
	"io/ioutil"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//several http methods available
//GET, POST, PUT, DELETE, ...

//three form to pass params
// 1. Path
// 2. Query
// 3. Headers
// 4. Body
func main() {
	gin.SetMode(gin.ReleaseMode) //remove debug warning
	router := gin.New()          //remove gin.Default() logging
	router.Use(gin.Recovery())   //catch exceptions
	router.Use(static)           //serve static webfiles
	router.GET("/api/:path_param", func(c *gin.Context) {
		//PATH
		//curl http://localhost:8080/api/12345
		path_param, ok := c.Params.Get("path_param")
		log.Println("path_param", path_param, ok)
		//QUERY
		//curl "http://localhost:8080/api/12345?query_param=abcdef"
		query_param, ok := c.GetQuery("query_param")
		log.Println("query_param", query_param, ok)
		//HEADERS
		//curl "http://localhost:8080/api/12345?query_param=abcdef" -H "header_param: xyx123"
		header_param := c.GetHeader("header_param")
		log.Println("header_param", header_param)
		//BODY
		//curl "http://localhost:8080/api/12345?query_param=abcdef" -X GET -d "987abc" -H "header_param: xyx123"
		body_bytes, err := ioutil.ReadAll(c.Request.Body)
		log.Println("body", string(body_bytes), err)
		//something happens before responding
		c.String(200, "ok")
	})
	//curl http://localhost:8080/ping -v
	//fetch("/ping").then(r => console.log(r))
	//fetch("/ping").then(r => r.json()).then(m => console.log(m))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	endpoint := "0.0.0.0:8080"
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Addr:    endpoint,
		Handler: router,
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err = server.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}

//go:embed webfiles/*
var webfiles embed.FS

func static(c *gin.Context) {
	path := "webfiles" + c.Request.URL.Path
	if path == "webfiles/" {
		path = "webfiles/index.html"
	}
	//log.Println(path)
	data, err := webfiles.ReadFile(path)
	if err != nil {
		c.Next()
	} else {
		ext := filepath.Ext(path)
		ct := mime.TypeByExtension(ext)
		c.Data(http.StatusOK, ct, data)
	}
}
