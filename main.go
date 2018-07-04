package main

import (
	"github.com/spf13/pflag"
	"apiserver/config"
	"github.com/gin-gonic/gin"
	"apiserver/router"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"errors"
)

var (
	cfg = pflag.StringP("config","c","","apiserver config file path.")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	//init db
	//model.DB.Init()
	//defer model.DB.Close()

	g := gin.New()

	middleware := []gin.HandlerFunc{}

	router.Load(
		g,
		middleware...,
	)
	go func() {
		if err := pingServer();err != nil{
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0;i < viper.GetInt("max_ping_count"); i++ {
		resp,err := http.Get(viper.GetString("url")+"/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}