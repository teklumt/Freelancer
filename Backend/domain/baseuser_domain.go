package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseUser struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username        string             `json:"username" bson:"username"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	ProfilePicture  string             `json:"profile_picture" bson:"profile_picture"`
	Role            string             `json:"role" bson:"role"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	ActivationToken string             `bson:"activation_token,omitempty" json:"activation_token,omitempty"`
	TokenCreatedAt  time.Time          `bson:"token_created_at"`
	RefreshTokens   []RefreshToken     `bson:"refresh_tokens" json:"refresh_tokens"`

	GoogleID           string `bson:"google_id,omitempty" json:"google_id,omitempty"`
	PasswordResetToken string `bson:"password_reset_token,omitempty" json:"password_reset_token,omitempty"`
}