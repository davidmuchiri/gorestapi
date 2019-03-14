package users

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dinobambino7/gorestapi/db"
	"github.com/dinobambino7/gorestapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//Token is a token structure
type Token struct {
	UserID interface{}
	jwt.StandardClaims
}

// Account is a structure for a user
type Account struct {
	UserID       interface{} `json:"_id" bson:"_id"`
	UserName     string      `json:"username" bson:"username"`
	UserEmail    string      `json:"email" bson:"email"`
	UserPassword string      `json:"userpassword" bson:"userpassword"`
	UserIMGURL   string      `json:"userimgurl" bson:"userimgurl"`
	Token        string      `json:"token" bson:"token"`
}

var client = db.Client
var collection = client.Database("gotest").Collection("accounts")
var ctx = context.Background()

//Validate is a method on Account struct validation when user is signing up
func (account *Account) Validate() (map[string]interface{}, bool) {
	// if email is invalid
	if !strings.Contains(account.UserEmail, "@") {
		return utils.Message(false, "email address is required", nil), false
	}
	if len(account.UserPassword) < 8 {
		return utils.Message(false, "Password is required", nil), false
	}

	temp := &Account{}

	filter := bson.M{"email": account.UserEmail}

	err := collection.FindOne(context.TODO(), filter).Decode(temp)

	if err != nil {
		fmt.Println("register err:", err)
	}

	if temp.UserEmail != "" {
		return utils.Message(false, "Email address already in use by another user", nil), false
	}

	return utils.Message(true, "Requirements passed", nil), true

}

//CreateAccount Creates a user account
func (account *Account) CreateAccount() map[string]interface{} {
	if res, ok := account.Validate(); !ok {
		return res
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.UserPassword), bcrypt.DefaultCost)

	log.Println("hashingerror", err)
	account.UserPassword = string(hashedPassword)

	res, err := collection.InsertOne(ctx, bson.M{
		"username":     account.UserName,
		"email":        account.UserEmail,
		"userpassword": account.UserPassword,
		"userimgurl":   account.UserIMGURL,
	})

	if res.InsertedID == nil || res.InsertedID == "" || err != nil {
		return utils.Message(false, "Failed to create account, check connection", nil)
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserID: res.InsertedID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.UserPassword = ""
	account.UserID = res.InsertedID

	response := utils.Message(true, "Account has been created", nil)
	response["account"] = account
	return response
}

//Login is a function that logs a registered user into the system
func Login(email, password string) map[string]interface{} {
	filter := bson.M{"email": email}
	var acc Account
	err := collection.FindOne(context.TODO(), filter).Decode(&acc)
	if err != nil {
		return utils.Message(true, "user email account not found", nil)
	}

	bcrypterr := bcrypt.CompareHashAndPassword([]byte(acc.UserPassword), []byte(password))

	// password and hash do not match
	if bcrypterr != nil && bcrypterr == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(true, "Invalid login credentials. please check password and email and try again", nil)
	}

	// worked Logged in

	acc.UserPassword = ""

	tk := &Token{UserID: acc.UserID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString

	return utils.Message(true, "Logged in", acc)

}

//GetUser is a func that gets user information
func GetUser(id string) map[string]interface{} {
	acc := &Account{}
	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(acc)

	if err != nil {
		return utils.Message(false, "could find user, connection error", nil)

	}

	if acc.UserEmail == "" {
		return utils.Message(true, "could find user in db", nil)
	}

	acc.UserPassword = ""

	return utils.Message(true, "user found!", acc)

}
