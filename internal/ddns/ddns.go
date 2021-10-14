package ddns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/gin-gonic/gin"
	"net/http"
	log "unknwon.dev/clog/v2"
)

type InfoModel struct {
	AccessKeyId  string `json:"accessKeyId"`
	AccessSecret string `json:"accessSecret"`
	Domain       string `json:"domain"`
	Record       string `json:"record"`
	Ip           string `json:"ip"`
}

func GetRecordId(info *InfoModel) (string, error) {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", info.AccessKeyId, info.AccessSecret)

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = info.Domain

	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	records := response.DomainRecords.Record
	for _, value := range records {
		if value.RR == info.Record {
			return value.RecordId, nil
		}
	}
	return "", err
}

func UpdateDomainRecord(c *gin.Context) (int, int, interface{}) {
	info := new(InfoModel)
	err := c.ShouldBindJSON(info)
	if err != nil {
		log.Warn("can't bind json")
		return http.StatusBadRequest, 40000, "can't bind json"
	}

	recordId, err := GetRecordId(info)
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", info.AccessKeyId, info.AccessSecret)

	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = recordId
	request.RR = info.Record
	request.Type = "A"
	request.Value = info.Ip

	if _, err = client.UpdateDomainRecord(request); err != nil {
		log.Error(err.Error())
		return http.StatusInternalServerError, 50000, "update record failed"
	}
	return http.StatusOK, 20000, "update record " + info.Record + " successful to " + info.Ip
}
