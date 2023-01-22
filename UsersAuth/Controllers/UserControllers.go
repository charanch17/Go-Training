package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "auth.com/charan/authenticate/Models"
	utils "auth.com/charan/authenticate/Utils"
	"golang.org/x/crypto/bcrypt"
)

func (c *controller) Signup(w http.ResponseWriter, r *http.Request) {
	var user, tmp models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Add("Content-Type", "application/json")
	if !user.CheckIsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "enter non empty email and password"})
		return
	}
	c.db.First(&tmp, "email = ?", user.Email)
	if tmp.CheckIsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "user already exists please login"})
		return

	}
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
	}
	user.Password = string(PasswordHash)
	c.db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	var user, StoredUser models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Add("Content-Type", "application/json")
	if !user.CheckIsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "enter non empty email and password"})
		return
	}
	c.db.First(&StoredUser, "email = ?", user.Email)
	if StoredUser.CheckIsValid() {
		err := bcrypt.CompareHashAndPassword([]byte(StoredUser.Password), []byte(user.Password))
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid username and password"})
			return

		}
		token, err := utils.GenerateJwtToken(user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "error while generating token"})
			fmt.Println(err)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "AuthorizationKey", Value: token})
		json.NewEncoder(w).Encode(map[string]string{"message": "logged in"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "invalid username and password"})

}

func (c *controller) Validate(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("AuthorizationKey")
	if err != nil {
		fmt.Println(err)
	}
	claims, valid, _ := utils.ValidateJwtToken(token.Value)
	if !valid {
		fmt.Println(valid)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Token is expired"})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)

}

func (c *controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("AuthorizationKey")
	tkn, _ := utils.RefreshJwtToken(token.Value)
	http.SetCookie(w, &http.Cookie{Name: "AuthorizationKey", Value: tkn})
	json.NewEncoder(w).Encode(map[string]string{"message": "Token is Refreshed"})

}
