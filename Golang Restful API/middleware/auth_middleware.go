package middleware

import (
	"encoding/json"
	"golang-restful-api/helper"
	"golang-restful-api/model/dto"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := dto.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   nil,
		}

		encoder := json.NewEncoder(writer)
		err := encoder.Encode(webResponse)
		helper.PanicIfError(err)
	}
}
