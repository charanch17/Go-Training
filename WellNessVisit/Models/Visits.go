package models

import (
	"time"

	"gorm.io/gorm"
)

type Visits struct {
	gorm.Model
	PatientID uint      `gorm:"not null" json:"patientid"`
	DoctorID  uint      `gorm:"not null" json:"doctorid"`
	VisitDate time.Time `json:"visitdate"`
	Patient   User      `gorm:"foreignKey:PatientID;references:ID"`
	Doctor    User      `gorm:"foreignKey:DoctorID;references:ID"`
}

func (visit *Visits) CheckIsValid() bool {
	return visit.PatientID != 0 && visit.DoctorID != 0 && visit.VisitDate.UnixNano() >= time.Now().Add(time.Hour*24).UnixNano()
}
