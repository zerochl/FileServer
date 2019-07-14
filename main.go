package main

import (
	"FileServer/server"
	"FileServer/server/entity"
	"github.com/gogather/com/log"
)

type StartCallBackTest struct {

}

func (scbt StartCallBackTest) OnError(err error) {
	log.Println("error:", err.Error())
}

func (scbt StartCallBackTest) OnSuccess(files []*entity.File) {
	log.Println("on success:", len(files))
}

func main()  {
	conf := &entity.FileServerConf{
		Port: "8081",
		Route: "/staticfile/",
		FilePath: "./staticfile",
	}
	startCallBacktest := StartCallBackTest{}
	fileserver.StartFileServer(conf, startCallBacktest)
}
