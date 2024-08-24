// Util : request from user --> to unmarshal or decode data
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func CheckPasswordHash(hash, password string) bool {
	fmt.Println(hash, password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//panic(err)
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
