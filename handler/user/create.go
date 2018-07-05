package user

import (
	"github.com/lexkong/log"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"github.com/lexkong/log/lager"
	"apiserver/util"
	"apiserver/model"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("user create function called.",lager.Data{"X-Request-Id":util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username:r.Username,
		Password:r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c,errno.ErrValidation,nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c,errno.ErrEncrypt,nil)
	}

	if err := u.Create();err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
	}

	rsp := CreateResponse{
		Username:r.Username,
	}

	SendResponse(c,nil,rsp)


}


func (r *CreateRequest) checkParam() error{
	if r.Username == "" {
		return errno.New(errno.ErrValidation,nil).Add("username is empty")
	}
	if r.Password == "" {
		return errno.New(errno.ErrValidation,nil).Add("password is empty")
	}

	return nil
}