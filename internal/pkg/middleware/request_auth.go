package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
)

func VerifyIncomingRequest(w http.ResponseWriter, r *http.Request) error {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Printf("Middleware: error authorization header empty")
		utils.ErrorResponse(r.Context(), w, http.StatusUnauthorized, apperrors.ErrHeaderMissing)
		return apperrors.ErrHeaderMissing
	}

	auth := NewAuthService()
	_, err := auth.VerifyToken(strings.TrimPrefix(authHeader, "Bearer "))
	if err != nil {
		log.Println("Middleware: error in auth check, err: ", err)
		utils.ErrorResponse(r.Context(), w, http.StatusUnauthorized, apperrors.ErrInvalidToken)
		return apperrors.ErrInvalidToken
	}
	// log.Println("tokenID: ", tokenID, " ", "userParam: ", userId)

	// Ensure the token user ID matches the request user ID
	// if tokenID != userId {
	// 	utils.ErrorResponse(r.Context(), w, http.StatusForbidden, apperrors.ErrIdMissmatch)
	// 	return apperrors.ErrIdMissmatch
	// }

	return nil
}
