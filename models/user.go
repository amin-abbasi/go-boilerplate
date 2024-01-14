package models

import (
	"context"
	"fmt"
	"log"

	"github.com/amin4193/go-boilerplate/services"
	"github.com/amin4193/go-boilerplate/configs"
	"go.mongodb.org/mongo-driver/bson"
)

// User struct definition and database operations
// ...

// User struct to represent a user
type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getCtx() (context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TIME_OUT_DURATION)
	defer cancel()
	return ctx
}

func (newUser User) Create() (User, error) {
	collection := services.DB.Collection("users")
	result, err := collection.InsertOne(getCtx(), newUser)
	if err != nil {
		return User{}, err
	}

	fmt.Println(">>> insert user result: ", result)

	// return result, err
	return newUser, nil
}

func (u User) List() ([]User, error) {
	collection := services.DB.Collection("users")
	// Implementation to retrieve list of users...
	ctx := getCtx()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetByUsername(username string) (*User, error) {
	collection := services.DB.Collection("users")
	ctx := getCtx()

	var user User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}