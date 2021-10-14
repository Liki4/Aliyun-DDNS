package web

import (
	"fmt"
	"github.com/Liki4/Aliyun-DDNS/internal/conf"
	"github.com/Liki4/Aliyun-DDNS/internal/ddns"
	"github.com/Liki4/Aliyun-DDNS/toolkit"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.POST("/updateDomainRecord", toolkit.Entry(ddns.UpdateDomainRecord))

	err := r.Run(fmt.Sprintf(":%d", conf.Get().Site.Port))
	if err != nil {
		panic(err)
	}
}
