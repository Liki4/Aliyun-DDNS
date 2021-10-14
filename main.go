package main

import (
	"github.com/Liki4/Aliyun-DDNS/internal/conf"
	"github.com/Liki4/Aliyun-DDNS/internal/web"
	log "unknwon.dev/clog/v2"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	if err = conf.Load(); err != nil {
		log.Fatal("Failed to load config: %v", err)
	}
	web.Run()
}
