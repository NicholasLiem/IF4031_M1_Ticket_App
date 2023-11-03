package middleware

import (
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"net/http"
	"os"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		clientAuthHeader := "Bearer " + os.Getenv("CLIENT_API_KEY")
		paymentAuthHeader := "Bearer " + os.Getenv("PAYMENT_API_KEY")
		if authHeader != "" && (authHeader == clientAuthHeader || authHeader == paymentAuthHeader) {
			next.ServeHTTP(w, r)
		} else {
			response.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized access")
			return
		}
	})
}
