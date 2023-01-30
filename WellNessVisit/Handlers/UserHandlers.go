package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "Privia.com/charan/WellNessVisits/Models"
	utils "Privia.com/charan/WellNessVisits/Utils"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) GetWellNessVisits(w http.ResponseWriter, r *http.Request) {
	fmt.Println(h.activeUser)
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.(string)})

		}

	}()
	if h.activeUser == nil || !h.activeUser.CheckIsValid() {
		panic("no active user")
	}
	QueryMap := r.URL.Query()
	var visits []models.Visits
	userIdFilter := fmt.Sprintf("%s_id = ?", h.activeUser.Role.RoleName)
	if QueryMap.Get("visits") == "Today" {
		err := h.db.Preload("Patient").Preload("Doctor").Where(userIdFilter, h.activeUser.ID).Find(&visits, "date(visit_date) = ?", time.Now().Format("2006-01-02")).Error
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(visits)
		return
	}
	if QueryMap.Get("visits") == "Upcoming" {
		err := h.db.Preload("Patient").Preload("Doctor").Where(userIdFilter, h.activeUser.ID).Find(&visits, "date(visit_date) >= ?", time.Now().Format("2006-01-02")).Error
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(visits)
		return

	}
	if QueryMap.Get("visits") == "Past" {
		err := h.db.Preload("Patient").Preload("Doctor").Where(userIdFilter, h.activeUser.ID).Find(&visits, "date(visit_date) < ?", time.Now().Format("2006-01-02")).Error
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(visits)
		return
	}
	err := h.db.Preload("Patient").Preload("Doctor").Find(&visits, userIdFilter, h.activeUser.ID).Error
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	json.NewEncoder(w).Encode(visits)

}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var user, tmp models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "unable to read request data"})
		return
	}
	if !user.CheckIsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Enter proper password and emailid"})
		return

	}
	h.db.First(&tmp, "email = ?", user.Email)
	if tmp.CheckIsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User already registered"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
		return
	}
	user.Password = string(hashedPassword)
	err = h.db.Create(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return

	}

	json.NewEncoder(w).Encode(user)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.(string)})
		}
	}()
	var user, tmp models.User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" || user.Password == "" {
		fmt.Println(user.Email, user.Password)
		panic("invalid credentials")

	}
	h.db.First(&tmp, "email = ?", user.Email)
	if tmp.Email == "" {
		panic("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(tmp.Password), []byte(user.Password))
	if err != nil {
		panic("invalid credentials")
	}
	tkn, err := utils.GenerateJwtToken(tmp.ID)
	if err != nil {
		fmt.Println(err)
		panic("unable to generate token")
	}
	http.SetCookie(w, &http.Cookie{Name: "AuthenticationToken", Value: tkn})
	json.NewEncoder(w).Encode(tmp)

}
