package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUser struct {
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
	Salt     string `json:"-" bson:"salt"`
}

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Salt     string             `json:"salt" bson:"salt"`
}
