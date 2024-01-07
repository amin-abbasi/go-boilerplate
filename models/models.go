package models

// User struct definition and database operations
// ...

var (
	DB         = make(map[string]User)
	SECRET_KEY = []byte("your_secret_key") // Replace with your secret key
)

// User struct to represent a user
type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
