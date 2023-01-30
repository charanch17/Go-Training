package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "Privia.com/charan/WellNessVisits/Models"
)

func (h *Handler) AddWellNessVisit(w http.ResponseWriter, r *http.Request) {
	var NewVisit, tmp models.Visits
	var Patient models.User
	var Doctor models.User
	err := json.NewDecoder(r.Body).Decode(&NewVisit)
	if err != nil {
		fmt.Println(err, h.activeUser.Email, NewVisit.VisitDate)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "failed to parse data"})
		return
	}
	err1 := h.db.First(&Patient, NewVisit.PatientID).Error
	err2 := h.db.First(&Doctor, NewVisit.DoctorID).Error
	if !NewVisit.CheckIsValid() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid visit info"})
		return
	}
	h.db.Where("patient_id = ? ", NewVisit.PatientID).Where("doctor_id = ?", NewVisit.DoctorID).First(&tmp, "date(visit_date) = ?", NewVisit.VisitDate.Format("2006-01-02"))
	if tmp.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "entry exist for the same day"})
		return
	}
	if NewVisit.PatientID != NewVisit.DoctorID && err1 == nil && err2 == nil && Patient.Email != Doctor.Email {
		err = h.db.Create(&NewVisit).Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		}

		json.NewEncoder(w).Encode(NewVisit)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	// json.NewEncoder(w).Encode(NewVisit)

	json.NewEncoder(w).Encode(map[string]string{"message": "doctor and patient are same person"})

}
