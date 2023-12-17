package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"user_management_golang/core"
	pb "user_management_golang/protoc/user_service"
	"user_management_golang/service"
)

type RestController struct {
	server *service.Server
}

func NewRestController(db *service.Server) *RestController {
	return &RestController{db}
}

func (r *RestController) Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "1.0.0"})
}

func (r *RestController) Register(c *gin.Context) {
	var req pb.RegisterReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	account := core.Account{
		UserId:         req.Username,
		Username:       req.Username,
		Password:       req.Password,
		Email:          req.Email,
		Phone:          req.Phone,
		FullName:       req.FullName,
		ProfilePicture: req.ProfilePicture,
		Status:         "activate",
	}
	success, err := r.server.UserRegister(account)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"Message": err.Error()})
	} else {
		if success {
			c.JSON(http.StatusOK, gin.H{"Message": "registered successfully."})
		} else {
			c.JSON(500, gin.H{"Message": "server unknown error."})
		}
	}
}

func (r *RestController) Login(c *gin.Context) {
	var req pb.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	account := core.Account{
		UserId:   req.Username,
		Username: req.Username,
		Password: req.Password,
	}
	token, err := r.server.Login(account)
	if err != nil {
		c.JSON(401, gin.H{
			"Message":      err.Error(),
			"SessionToken": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message":      "login  successful.",
			"SessionToken": token,
		})
	}
}

func (r *RestController) Logout(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "" && password != "" {
		account := core.Account{
			UserId:   username,
			Password: password,
		}
		err := r.server.LogoutByCredentials(account)
		if err != nil {
			c.JSON(401, gin.H{"Message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "logout successful.",
		})
		return
	}

	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.LogoutByToken(*account, token)
		if err != nil {
			c.JSON(401, gin.H{"Message": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "logout  successful.",
			})
		}
	})
}

func (r *RestController) CheckTokenValid(c *gin.Context) {
	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Message": "the token is invalid"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "the token is valid"})
		return
	})
}

func (r *RestController) SearchRole(c *gin.Context) {
	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Message": "the token is invalid"})
			return
		}
		roleStr, err := r.server.SearchRole(account)
		if err != nil {
			c.JSON(500, gin.H{"Message": "server internal error"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Message": roleStr})
			return
		}
	})
}

func (r *RestController) SearchGroup(c *gin.Context) {
	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Message": "the token is invalid"})
			return
		}
		groupStr, err := r.server.SearchGroup(account)
		if err != nil {
			c.JSON(500, gin.H{"Message": "server internal error"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Message": groupStr})
			return
		}
	})
}

func (r *RestController) SearchPermission(c *gin.Context) {
	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Message": "the token is invalid"})
			return
		}
		groupStr, err := r.server.SearchGroup(account)
		if err != nil {
			c.JSON(500, gin.H{"Message": "server internal error"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Message": groupStr})
			return
		}
	})
}

func (r *RestController) Edit(c *gin.Context) {
	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Message": "the token has expired"})
		} else {
			r.server.Edit(account)
			c.JSON(200, gin.H{"Message": "ok"})
		}
		return
	})
}

func (r *RestController) GetUserId(c *gin.Context) {
	verifyToken(c, func(account *core.Account, token string) {
		err := r.server.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Message": "the token has expired"})
			return
		} else {
			c.JSON(401, gin.H{"Message": account.UserId})
			return
		}
	})
}

// CreateRole
//func (r *RestController) CreateRole(c *gin.Context) {
//	verifyToken(c, func(account *core.Account, token string) {
//		r.server.
//	})
//}

func verifyToken(c *gin.Context, callFunc func(account *core.Account, token string)) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"Message": "missing Authorization header"})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(401, gin.H{"Message": "invalid Authorization header format"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := service.GetUserIdFromToken(token)
	if err != nil {
		c.JSON(401, gin.H{"Message": err.Error()})
		return
	}

	success, err := service.IsTokenValid(token)
	if err != nil {
		c.JSON(401, gin.H{"Message": err.Error()})
		return
	}

	if !success {
		c.JSON(401, gin.H{"Message": "the token has expired"})
		return
	}

	account := &core.Account{UserId: userId}
	callFunc(account, token)
}
