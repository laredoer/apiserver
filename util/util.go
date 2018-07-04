package util

import (
	"github.com/teris-io/shortid"
	"github.com/gin-gonic/gin"
)

func GetShortId() (string,error){
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v,ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId,ok := v.(string); ok{
		return requestId
	}
	return ""
}
