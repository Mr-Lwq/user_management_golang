package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	pb "user_management_golang/protoc/user_service"
	"user_management_golang/service"
	"user_management_golang/src"
	"user_management_golang/utils"
)

type RestService struct{}

func NewRestService() *RestService {
	return &RestService{}
}

func (r *RestService) Register(c *gin.Context) {
	var req pb.RegisterReq
	// 使用 ShouldBindJSON 将请求体绑定到 req 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := []byte(req.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPasswordStr := base64.StdEncoding.EncodeToString(hashedPassword)
	sessionToken, _ := utils.GenerateSessionToken(req.Username)
	account := src.Account{
		UserId:         req.Username,
		Username:       req.Username,
		Password:       hashedPasswordStr,
		Email:          req.Email,
		Phone:          req.Phone,
		FullName:       req.FullName,
		ProfilePicture: req.ProfilePicture,
		Status:         "activate",
		SessionToken:   sessionToken,
	}
	success, err := service.UserRegister(account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "OK."})
	} else {
		if success {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "OK."})
		} else {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "OK."})
		}
	}

}

func (r *RestService) Login(c *gin.Context) {}
