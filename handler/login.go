package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type LoginHandler struct{
	svc *service.UserService
}

func NewLoginHandler(svc *service.UserService) *LoginHandler {
	return &LoginHandler{
		svc: svc,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var res model.ReadUserResponse

		params := r.URL.Query()
		name := params.Get("username")
		if name == "" {
			http.Error(w, "username is missing", http.StatusUnauthorized)
			log.Println("Login: Failed to get parameters")
			return
		}

		user, err := h.svc.ReadUser(r.Context(), name)
		if err != nil {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			log.Println("Login: Authentification filed, err = ", err)
			return
		}

		res.UserId = user.UserId
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}
	}
}
