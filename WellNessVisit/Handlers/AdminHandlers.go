package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "Privia.com/charan/WellNessVisits/Models"
	"github.com/gorilla/mux"
)

func (h *Handler) AddRole(w http.ResponseWriter, r *http.Request) {
	var newRole, tmp models.Role
	json.NewDecoder(r.Body).Decode(&newRole)
	h.db.First(&tmp, "role_name = ? ", newRole.RoleName)
	if newRole.RoleName != "" && tmp.RoleName == "" {
		err := h.db.Create(&newRole).Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "Inserted Record"})
		return

	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"message": "invalid rolename or role already exists"})

}

func (h *Handler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	var tmp models.Role
	params := mux.Vars(r)
	fmt.Println(params["rolename"])
	h.db.First(&tmp, "role_name = ?", params["rolename"])
	if params["rolename"] != "" && tmp.RoleName != "" {
		err := h.db.Delete(&tmp).Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "Deleted Record"})
		return

	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"message": "invalid rolename or role not found"})

}
