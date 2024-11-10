package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    string             `json:"first_name" validate:"required, min=2, max=50"`
	Last_name     string             `json:"last_name" validate:"required, min=2, max=50"`
	Password      string             `json:"password" validate:"required, min=6"`
	Phone         string             `json:"phone" validate:"required, max=10"`
	Token         string             `json:"token"`
	Refresh_token string             `json:"refresh_token"`
	User_type     string             `json:"user_type" validate:"required, eq=ADMIN | USER"`
	Email         string             `json:"email" validate:"email,required"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
