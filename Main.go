package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
import "github.com/TarsCloud/TarsGo/tars"
func main() {
	r := gin.Default()
	GetM4a()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

	//r.GET("/benchmark", MyBenchLogger())
	mux := &tars.TarsHttpMux{}
	//mux.HandleFunc("/test1", HttpRootHandler)
	//mux.HandleFunc("/test2", HttpRootHandler2)
	cfg := tars.GetServerConfig()                               //Get Config File Object
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".GoHttpObj")  //Register http server

	tars.Run()

}

func GetM4a() {
	res, err := http.Get("http://openod.sign.qingting.fm/vod/00/00/0000000000000000000025794354_24.m4a?sign=04c527dce510ee8a3db279af1ce7bf64&t=5c520210")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error
		fmt.Println("err is:",err.Error())
	}
	fmt.Println("filesize is:",len(body))


}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("before middleware")
	}
}

