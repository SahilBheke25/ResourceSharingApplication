package login

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var secretKey = []byte("hmacSampleSecret")

// func encode(username string) (string, error) {

// 	loginExiration := time.Now().Add(time.Hour * 24)

// 	loginClaims := jwt.MapClaims{
// 		"username": username,
// 		"exp":      loginExiration.Unix(),
// 		"iat":      time.Now().Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, loginClaims)

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println(tokenString, err)

// 	return tokenString, nil
// }

// func decode(tokenString string) (string, error){

// 	Dectoken, err := jwt.Parse(tokenString, func(Dectoken *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := Dectoken.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", Dectoken.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return secretKey, nil
// 	})
// 	if err != nil {
// 		log.Fatal("error")
// 	}

// 	if claims, ok := Dectoken.Claims.(jwt.MapClaims); ok {
// 		fmt.Println(claims["foo"], claims["nbf"])
// 	} else {
// 		fmt.Println(err)
// 	}
// }

// // func generateToken(length int) (string, error) {
// // 	bytes := make([]byte, length)
// // 	if _, err := rand.Read(bytes); err != nil {
// // 		err = fmt.Errorf("Failed to generate token: %v", err)
// // 		return "", err
// // 	}
// // 	return base64.URLEncoding.EncodeToString(bytes), nil
// // }

// // var AuthError = errors.New("Unauthorized")

// // func Authorize(r *http.Request) error {
// // 	username := r.F
// // }
