package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/pkg/errno"
	"github.com/lexkong/log"
	"apiserver/handler"
	"github.com/lexkong/log/lager"
	"apiserver/util"
	"apiserver/model"
	"strconv"
	"apiserver/service"
	"apiserver/pkg/auth"
	"apiserver/pkg/token"
)

type CreateRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct{
	Username string `json:"username"`
}

type ListRequest struct{
	Username	string		`json:"username"`
	Offset		uint		`json:"offset"`
	Limit		uint		`json:"limit"`
}

type ListResponse struct{
	TotalCount	uint64	`json:"totalCount"`
	UserList	[]*model.UserInfo	`json:"userList"`
}

func Create(c *gin.Context){
	log.Info("User Create function called.",lager.Data{"X-Request-Id":util.GetReqID})
	var r CreateRequest
	if err := c.Bind(&r); err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}

	u := model.UserModel{
		Username:r.Username,
		Password:r.Password,
	}

	r.checkParam()

	if err := u.Validate();err !=nil{
		handler.SendResponse(c,errno.ErrValidation,nil)
		return
	}

	if err := u.Encrypt();err !=nil{
		handler.SendResponse(c,errno.ErrEncrypt,nil)
		return
	}

	if err := u.Create();err != nil{
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	rsp := CreateResponse{Username:r.Username}
	handler.SendResponse(c,nil,rsp)
}

func(r *CreateRequest) checkParam() (err error){

	if r.Username == ""{
		return errno.New(errno.ErrValidation,nil).Add("username is empty.")
	}

	if r.Password == ""{
		return errno.New(errno.ErrValidation,nil).Add("password is empty.")

	}
	return nil
}



func Delete(c *gin.Context){
	userId,_ := strconv.Atoi(c.Param("id"))
	if err := model.Delete(uint64(userId));err!=nil{
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	handler.SendResponse(c,nil,nil)
}

func Update(c *gin.Context){
	log.Info("Update function called.",lager.Data{"X-Request-Id":util.GetReqID(c)})
	userId,_ := strconv.Atoi(c.Param("id"))
	var u model.UserModel
	if err := c.Bind(&u);err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}

	u.Id = uint64(userId)

	if err := u.Validate();err!=nil{
		handler.SendResponse(c,errno.ErrValidation,nil)
		return
	}

	if err := u.Encrypt(); err != nil{
		handler.SendResponse(c,errno.ErrEncrypt,nil)
		return
	}

	if err := u.Update(); err !=nil{
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	handler.SendResponse(c,nil,nil)
}

func Get(c *gin.Context){
	username := c.Param("username")
	user,err :=model.GetUser(username)
	if err != nil{
		handler.SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}
	handler.SendResponse(c,nil,user)
}

func List(c *gin.Context){
	var r ListRequest
	if err := c.Bind(&r); err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}
	infos,count,err := service.ListUser(r.Username,r.Offset,r.Limit)
	if err!=nil{
		handler.SendResponse(c,err,nil)
		return
	}
	handler.SendResponse(c,nil,ListResponse{
		TotalCount:count,
		UserList:infos,
	})
}

func Login(c *gin.Context){
	var u  model.UserModel
	if err :=c.Bind(&u);err!=nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}

	d,err :=model.GetUser(u.Username)
	if err != nil{
		handler.SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	if err:= auth.Compare(d.Password,u.Password);err!=nil{
		handler.SendResponse(c,errno.ErrPasswordIncorrect,nil)
		return
	}
	t,err := token.Sign(c,token.Context{ID:d.Id,Username:d.Username},"")
	if err != nil{
		handler.SendResponse(c,errno.ErrToken,nil)
		return
	}
	handler.SendResponse(c,nil,model.Token{Token:t})

}