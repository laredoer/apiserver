package user

import (
	"github.com/lexkong/log"
	"apiserver/pkg/errno"
	"fmt"
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

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
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