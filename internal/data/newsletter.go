package data

import (
	"canvas/validator"
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"time"
)

type Newsletter struct {
    ID     int  `json:"id"`
    Email  string `json:"email"`
    token  string `json:"token"`
    confirmed bool `json:"confirm"`
    active  bool `json:"active"`
    createdAt string `json:"created_at"`
    updatedAt string `json:"updated_at"`
}

type NewsletterModel struct {
    DB *sql.DB
}


func (m NewsletterModel) Insert(email string) (string,error) {

    token, err := generateToken()
    if err != nil {
        return "", err
    }

    stmt:= `INSERT INTO newsletter_subscribers(email,token) 
            VALUES ($1,$2) ON CONFLICT(email) DO UPDATE SET 
            token = EXCLUDED.token, updated_at = now()`

   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()

   _,err = m.DB.ExecContext(ctx,stmt,email,token)

   return token,err

}

func generateToken() (string, error) {
    secret:= make([]byte, 32)

    _, err := rand.Read(secret)

    if err != nil {
        return "", err
    }
    token := fmt.Sprintf("%x", secret)
    return token, nil
}


func ValidateEmail(v *validator.Validator, email string) {
    v.Check(email != "", "email", "must be provided")
    v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

// func ValidateUser(v *validator.Validator, user *User) {
// 	v.Check(user.Name != "", "name", "must be provided")
// 	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

// 	// Call the standalone ValidateEmail() helper.
// 	ValidateEmail(v, user.Email)

// }