package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"apiserver/util"
	"strconv"
	"apiserver/model"
	. "apiserver/handler"
	"apiserver/pkg/errno"
)

func Update(c *gin.Context) {
	log.Info("update function called",lager.Data{"X-Request-Id":util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}
	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		SendResponse(c,errno.ErrValidation,nil)
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c,errno.ErrEncrypt,nil)
		return
	}
	if err := u.Update(); err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
	}

	SendResponse(c,nil,nil)

}
