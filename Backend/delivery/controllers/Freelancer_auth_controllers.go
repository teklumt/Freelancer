package controllers

import (
	"felagi/domain"

	"github.com/gin-gonic/gin"
)

type FreelancerController struct {
	FreelancerUsecase domain.FreelancerUsecase
}

func NewFreelancerController(freelancerUsecase domain.FreelancerUsecase) *FreelancerController {
	return &FreelancerController{FreelancerUsecase: freelancerUsecase}
}


func (c *FreelancerController) Register(ctx *gin.Context) {
	var user domain.FreelancerUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"error": err.Error()})
		return
	}

	err := c.FreelancerUsecase.Register(user)
	if err.Message != "" {
		ctx.JSON(err.StatusCode, err)
		return
	}

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Your account has been created successfully see your email to activate your account",
	})	
}


func (c *FreelancerController) ActivateAccount(ctx *gin.Context) {
	type ActivateAccountRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	var req ActivateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"error": err.Error()})
		return
	}

	err := c.FreelancerUsecase.ActivateAccount( req.Email, req.OTP)
	if err.Message != "" {
		ctx.JSON(err.StatusCode, err)
		return
	}

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Your account has been activated successfully",
	})
}