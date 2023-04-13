package router

import (
	"net/http"
	"os"

	"github.com/MochamadAkbar/dimy-practical-test/common"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if errRecover := recover(); errRecover != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				errResponse := common.ErrResponse{
					Code:    http.StatusInternalServerError,
					Status:  http.StatusText(http.StatusInternalServerError),
					Message: errTranslate(errRecover),
				}

				common.SerializeWriter(w, http.StatusInternalServerError, errResponse)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func errTranslate(err interface{}) interface{} {
	isDev := os.Getenv("APP_ENV")

	if isDev == "development" {
		return err
	} else {
		return "There was an internal server error"
	}
}