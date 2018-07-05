package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"apiserver/model"
	"apiserver/handler"
	"apiserver/pkg/errno"
)

func Delete(c *gin.Context)  {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	handler.SendResponse(c,nil,nil)
}