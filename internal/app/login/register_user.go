package login

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/Repository"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/handle"
)

// const (
// 	EncryptionKey = "thisis32bitlongpassphrase!"
// )

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Pincode    int    `json:"pincode"`
	Uid        int    `json:"uid"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Ensure body gets closed
	defer r.Body.Close()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
		return
	}

	_, err = validateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	// ecrypting password
	// user.Password, err = Encrypt(user.Password, EncryptionKey)
	// if err != nil {
	// 	err = fmt.Errorf("Internal Server Error while decrypting password: %v", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// create user & checks if user already exist.
	err = Repository.CreateUser(user.Username, user.Password, user.First_name, user.Last_name, user.Email, user.Phone, user.Address, user.Pincode, user.Uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handle.HandleResponse(w, "User Registered Successfully ", r)
}

func validateUser(user User) (bool, error) {
	if len(user.Username) < 2 {

		return false, fmt.Errorf("Username too short")

	} else if len(user.First_name) < 2 {

		return false, fmt.Errorf("First Name too short")

	} else if len(user.Last_name) < 2 {

		return false, fmt.Errorf("Last Name too short")

	} else if len(user.Phone) != 10 {

		return false, fmt.Errorf("Invalid phone number")

	} else if len(strconv.Itoa(user.Uid)) != 12 {

		return false, fmt.Errorf("Invalid UID")

	} else if len(user.Password) < 8 {

		return false, fmt.Errorf("Password too short")

	}

	return true, nil
}

// func hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 2)
// 	return string(bytes), err
// }

// Encrypt encrypts a string using AES-GCM
func Encrypt(password, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(password), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts a string encrypted with AES-GCM
func Decrypt(encrypted, key string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// func handleResponse(w http.ResponseWriter, message any, r *http.Request) {
// 	res, err := json.Marshal(message)
// 	if err != nil {
// 		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// "username":"sahilbheke",
// "password" : "Aim@1045",
// "firstname": "Ritesh",
// "lastname":"bheke",
// "email":"sahil@example1.com",
// "phone":"8856026645",
// "address": "402, T3, Ace Aurum3, Ravet",
// "pincode": 412101,
// "uid":123456789123
