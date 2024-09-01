package domain

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)


type FreelancerUser struct {
	BaseUser
	Sex            string               `json:"sex"   bson:"sex"`
	Phone          string               `json:"phone" bson:"phone"`
	Bio            string               `json:"bio" bson:"bio"`
	Skills         []string             `json:"skills" bson:"skills"`
	Experience     []primitive.ObjectID `json:"experience" bson:"experience"` // Array of job references
	EducationIDs   []primitive.ObjectID `json:"education_ids" bson:"education_ids"` // Array of education references
	Rating         float64              `json:"rating" bson:"rating"`
	Reviews        []primitive.ObjectID `json:"reviews" bson:"reviews"` // Array of review references
	Location       string               `json:"location" bson:"location"`
	Languages      []string             `json:"languages" bson:"languages"`
	HourlyRate     float64              `json:"hourly_rate" bson:"hourly_rate"`
	Availability   string               `json:"availability" bson:"availability"`
	MemberSince    primitive.Timestamp  `json:"member_since" bson:"member_since"`
	LastActive     primitive.Timestamp  `json:"last_active" bson:"last_active"`
	Portfolios     []Portfolio          `json:"portfolios" bson:"portfolios"`
}

type Job struct {
	Title       string    `json:"title" bson:"title"`                 // Job title
	Description string    `json:"description" bson:"description"`     // Brief description of the job
	Company     string    `json:"company" bson:"company"`             // Company or client name
	StartDate   time.Time `json:"start_date" bson:"start_date"`       // Start date of the job
	EndDate     time.Time `json:"end_date" bson:"end_date"`           // End date of the job (if applicable)
}

type Review struct {
	ReviewerName string    `json:"reviewer_name" bson:"reviewer_name"`   // Name of the reviewer
	Comment      string    `json:"comment" bson:"comment"`               // Review comment
	Rating       float64   `json:"rating" bson:"rating"`                 // Rating given by the reviewer
	Date         time.Time `json:"date" bson:"date"`                     // Date when the review was written
}

type Education struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`               // Reference to the User ID
	School       string             `json:"school" bson:"school"`
	Degree       string             `json:"degree" bson:"degree"`
	FieldOfStudy string             `json:"field_of_study" bson:"field_of_study"`
	StartDate    time.Time          `json:"start_date" bson:"start_date"`
	EndDate      time.Time          `json:"end_date" bson:"end_date"`
}

type Portfolio struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`               // Reference to the User ID
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Images      []string           `json:"images" bson:"images"`
	Link        string             `json:"link" bson:"link"`
}

type ReturnUser struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string              `bson:"username" json:"username"`
	Email     string              `bson:"email" json:"email"`
	Role      string              `bson:"role" json:"role"`
	CreatedAt primitive.Timestamp `bson:"createdAt" json:"createdAt"`

	FirstName      string               `json:"first_name" bson:"first_name"`
	LastName       string               `json:"last_name" bson:"last_name"`
	ProfilePicture string               `json:"profile_picture" bson:"profile_picture"`

	Phone          string               `json:"phone" bson:"phone"`
	Bio            string               `json:"bio" bson:"bio"`
	Skills         []string             `json:"skills" bson:"skills"`
}

func (u *FreelancerUser) MarshalJSON() ([]byte, error) {
	type Alias FreelancerUser
	return json.Marshal(&struct {
		*Alias
		Username      string               `json:"username,omitempty"`
		Image         string               `json:"image,omitempty"`
		MemberSince   primitive.Timestamp  `json:"member_since"`
		LastActive    primitive.Timestamp  `json:"last_active"`
		Skills        []string             `json:"skills"`
		Rating        float64              `json:"rating"`
		Location      string               `json:"location"`
		Languages     []string             `json:"languages"`
		HourlyRate    float64              `json:"hourly_rate"`
		Availability  string               `json:"availability"`
		Phone         string               `json:"phone,omitempty"`
		Bio           string               `json:"bio,omitempty"`
	}{
		Alias:         (*Alias)(u),
		Username:      u.Username,
		Image:         u.ProfilePicture,
		MemberSince:   u.MemberSince,
		LastActive:    u.LastActive,
		Skills:        u.Skills,
		Rating:        u.Rating,
		Location:      u.Location,
		Languages:     u.Languages,
		HourlyRate:    u.HourlyRate,
		Availability:  u.Availability,
		Phone:         u.Phone,
		Bio:           u.Bio,
	})
}

type OAuthProvider string

const (
	Google OAuthProvider = "google"
)

type RefreshToken struct {
	Token     string    `bson:"token" json:"token"`
	DeviceID  string    `bson:"device_id" json:"device_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type LogInResponse struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         ReturnUser  `json:"user"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type TokenGenerator interface {
	GenerateToken(user BaseUser) (string, error)
	GenerateRefreshToken(user BaseUser) (string, error)
	RefreshToken(token string) (string, error)
}

type TokenVerifier interface {
	VerifyToken(token string) (*BaseUser, error)
	VerifyRefreshToken(token string) (*BaseUser, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type FreelancerUsecase interface {

	// for every user
	Register(user FreelancerUser) ErrorResponse
	ActivateAccount(email, otp string) ErrorResponse

}


type FreelancerRepository interface {
	// for every user

	Register(user FreelancerUser) error
	GetUserByUsernameOrEmail(username, email string) (FreelancerUser, error)
	ActivateAccount(email string) error
	GetUserByEmail(email string) (FreelancerUser, error)
	// GetUserByUsernameOrEmail(username, email string) (User, error)
	// AccountActivation(email string) error
	// GetUserByEmail(email string) (User, error)

	// Login(user *User) (*User, error)
	// UpdateUser(user *User) error
	// DeleteRefreshToken(user *User, refreshToken string) error
	// DeleteAllRefreshTokens(user *User) error

	// GetUserByID(id string) (User, error)
	
	// GetUserByResetToken(token string) (User, error)

	// // // ActivateAccountMe(Email string) error

	// // // // for user profile
	// GetMyProfile(userID string) (User, error)
	// GetUsers(byName, limit , page string) ([]User, error)
	// DeleteUser(userID string) (User, error)

	

}
