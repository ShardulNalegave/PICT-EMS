package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ShardulNalegave/PICT-EMS/database/models"
	"github.com/ShardulNalegave/PICT-EMS/utils"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func initPeopleRoutes(r *chi.Mux) {
	r.Get("/people/students/{id}", getStudent)
	r.Get("/people/students", getStudents)
	r.Get("/people/staff-members/{id}", getStaffMember)
	r.Get("/people/staff-members", getStaffMembers)
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)

	var students []models.Student
	if err := db.Find(&students).Error; err != nil {
		http.Error(w, "Couldn't get all student records", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(students)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)

	var student models.Student
	if err := db.First(&student, "registration_id = ?", chi.URLParam(r, "id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Student record doesn't exist", http.StatusNotFound)
		} else {
			http.Error(w, "Couldn't get student record", http.StatusInternalServerError)
		}
		return
	}

	data, _ := json.Marshal(student)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getStaffMembers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)

	var staffMembers []models.StaffMember
	if err := db.Find(&staffMembers).Error; err != nil {
		http.Error(w, "Couldn't get all staff-member records", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(staffMembers)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getStaffMember(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)

	var staffMember models.StaffMember
	if err := db.First(&staffMember, "registration_id = ?", chi.URLParam(r, "id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Staff-member record doesn't exist", http.StatusNotFound)
		} else {
			http.Error(w, "Couldn't get staff-member record", http.StatusInternalServerError)
		}
		return
	}

	data, _ := json.Marshal(staffMember)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
