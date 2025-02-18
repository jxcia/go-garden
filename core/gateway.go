package core

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jxcia/go-garden/core/log"
)

func (g *Garden) gateway(c *gin.Context) {
	// get openTracing span
	span := GetSpan(c)
	// get request datatype
	request := GetRequest(c)
	//debug
	body := getBody(c)
	service := c.Param("service")
	action := c.Param("action")

	// request service
	code, data, header, err := g.callServiceapi(span, service, action, request, body, nil, nil)
	if err != nil {
		c.JSON(code, gatewayFail(data))
		log.Error("callService", err)
		span.SetTag("CallService", err)
		return
	}
	var result MapData
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		c.JSON(httpFail, gatewayFail(infoServerError))
		log.Error("returnInvalidFormat", err)
		span.SetTag("ReturnInvalidFormat", err)
		return
	}

	for k, v := range header {
		if k != "Content-Type" && k != "Date" && k != "Content-Length" {
			c.Header(k, v[0])
		}
	}
	c.JSON(code, gatewaySuccess(result))
}

func gatewaySuccess(data MapData) MapData {
	response := MapData{
		"status": true,
	}
	for k, v := range data {
		response[k] = v
	}
	return response
}

func gatewayFail(message string) MapData {
	response := MapData{
		"status": false,
		"msg":    message,
	}
	return response
}
