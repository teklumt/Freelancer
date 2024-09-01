package usecase

import (
	"felagi/domain"
	"felagi/infrastracture"
	"fmt"
	"time"
)

type FreelancerUsecase struct {
	FreelancerRepo    domain.FreelancerRepository
	TokenGen    domain.TokenGenerator
	PasswordSvc domain.PasswordService
}

func NewFreelancerUsecase(freelancerRepo domain.FreelancerRepository, tokenGen domain.TokenGenerator, passwordSvc domain.PasswordService) domain.FreelancerUsecase {
	return &FreelancerUsecase{
		FreelancerRepo:    freelancerRepo,
		TokenGen:    tokenGen,
		PasswordSvc: passwordSvc,
	}
}



func (u *FreelancerUsecase) Register(user domain.FreelancerUser) domain.ErrorResponse {
	fmt.Println("Registering user*******", user)
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "Missing required fields",
		}
	}

	if !infrastracture.IsValidEmail(user.Email) {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "Invalid email",
		}
	}

	if !infrastracture.IsValidPassword(user.Password) {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "Invalid password",
		}
	}

	_, err := u.FreelancerRepo.GetUserByEmail( user.Email)
	if err == nil {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "User already exists",
		}
	}

	user.Role = "freelancer"

	// Hash password
	hashedPassword, err := u.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.ErrorResponse{
			StatusCode: 500,
			Message:    "Internal server error",
		}
	}

	token, err := infrastracture.GenerateOTP()
	if err != nil {
		return domain.ErrorResponse{
			StatusCode: 500,
			Message:    "Internal server error",
		}
	}

	user.Password = hashedPassword
	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()

	// Create user account in the database
	err = u.FreelancerRepo.Register(user)
	if err != nil {
		return domain.ErrorResponse{
			StatusCode: 500,
			Message:    "Internal server error",
		}
	}

	// Send activation email or link to the user
	err = infrastracture.SendActivationEmail(user.Email, token)
	if err != nil {
		return domain.ErrorResponse{
			StatusCode: 500,
			Message:    "Failed to send email",
		}
	}

	return domain.ErrorResponse{}
}


func (u *FreelancerUsecase) ActivateAccount(email, otp string) domain.ErrorResponse {
	fmt.Println("Activating account*******", email, otp)
	user, err := u.FreelancerRepo.GetUserByUsernameOrEmail("", email)
	if err != nil {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "User not found",
		}
	}

	if user.ActivationToken != otp {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "Invalid OTP",
		}
	}

	if time.Since(user.TokenCreatedAt) > 30*time.Minute {
		return domain.ErrorResponse{
			StatusCode: 400,
			Message:    "OTP expired",
		}
	}

	err = u.FreelancerRepo.ActivateAccount(email)
	if err != nil {
		return domain.ErrorResponse{
			StatusCode: 500,
			Message:    "Internal server error",
		}
	}

	return domain.ErrorResponse{}
}