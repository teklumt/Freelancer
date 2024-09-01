package routers

import (
	"felagi/config/db"
	"felagi/delivery/controllers"
	"felagi/infrastracture"
	"felagi/repository"
	"felagi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpFreelancerAuth(router *gin.Engine) {

	freelancerRepo := repository.NewfreelancerRepositoryImpl(db.FreelancerCollection)

	tokenGen := infrastracture.NewTokenGenerator()
	passwordSvc := infrastracture.NewPasswordService()


	freelancerUsecase := usecase.NewFreelancerUsecase(freelancerRepo, tokenGen, passwordSvc)


	freelancerAuthController := controllers.NewFreelancerController(freelancerUsecase)
	// controllers.NewFreelancerController(freelancerUsecase)

	auth := router.Group("/freelancer")
	{
		// auth.POST("/login", authController.Login)
		auth.POST("/register", freelancerAuthController.Register)
		auth.POST("/activate", freelancerAuthController.ActivateAccount)

		// // OAuth routes
		// auth.GET("/login/google", authController.HandleGoogleLogin)
		// auth.GET("/callback", authController.HandleGoogleCallback)

		// // Password reset routes
		// auth.POST("/reset-password", authController.SendPasswordResetLink)
		// auth.POST("/reset-password/:token", authController.ResetPassword)
	}
}
