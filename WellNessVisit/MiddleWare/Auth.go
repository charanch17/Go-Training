package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "Privia.com/charan/WellNessVisits/DB"
	handlers "Privia.com/charan/WellNessVisits/Handlers"
	models "Privia.com/charan/WellNessVisits/Models"
	utils "Privia.com/charan/WellNessVisits/Utils"
)

func AuthenticateAndCheckPermission(handler http.HandlerFunc, AllowedRoles []string, handlerInstance *handlers.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("AuthenticationToken")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Unable to Authorize/Token not found"})
			return
		}
		claims, err := utils.ValidateJwtToken(tokenCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid Token Please Login again"})
			return
		}
		var user models.User
		err = db.DbInst.Preload("Role").First(&user, claims.Uid).Error
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Unable to Authorize or fetch user info"})
			return
		}
		var hasPermission bool = false
		for _, role := range AllowedRoles {
			if role == user.Role.RoleName {
				hasPermission = true
				break
			}
		}
		if hasPermission {
			handlerInstance.SetactiveUser(&user)
			fmt.Println(handlerInstance)
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorized access"})

		}
	}
}
