package handlers

import (
	models "dewetour/1models"
	repositories "dewetour/4repositories"
	dto "dewetour/5dto/result"
	tripsdto "dewetour/5dto/trips"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTrips struct {
	TripsRepository repositories.TripsRepository
}

func HandlerTrips(TripsRepository repositories.TripsRepository) *handlerTrips {
	return &handlerTrips{TripsRepository}
}

func (h *handlerTrips) FindTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trips, err := h.TripsRepository.FindTrips()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrips) GetTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var trips models.Trips
	trips, err := h.TripsRepository.GetTrips(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrips(trips)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrips) MakeTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(tripsdto.TripsRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	trips := models.Trips{
		Name:           request.Name,
		Desc:           request.Desc,
		Price:          request.Price,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Duration:       request.Duration,
		DateTrip:       request.DateTrip,
		Qty:            request.Qty,
		Image:          request.Image,
		UserID:         request.UserID,
	}

	trips, err = h.TripsRepository.MakeTrips(trips)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trips, _ = h.TripsRepository.GetTrips(trips.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTrips(u models.Trips) models.TripsResponse {
	return models.TripsResponse{
		Name:           u.Name,
		Desc:           u.Desc,
		Price:          u.Price,
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Duration:       u.Duration,
		DateTrip:       u.DateTrip,
		Qty:            u.Qty,
		Image:          u.Image,
		User:           u.User,
		Country:        u.Country,
	}
}

func (h *handlerTrips) EditTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(tripsdto.UpdateTripsRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ID, _ := strconv.Atoi(mux.Vars(r)["id"])

	user := models.Trips{}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Desc != "" {
		user.Desc = request.Desc
	}

	if request.Accomodation != "" {
		user.Accomodation = request.Accomodation
	}

	if request.Transportation != "" {
		user.Transportation = request.Transportation
	}

	if request.Eat != "" {
		user.Eat = request.Eat
	}

	if request.Duration != "" {
		user.Duration = request.Duration
	}

	if request.DateTrip != "" {
		user.DateTrip = request.DateTrip
	}

	// if request.Qty !=  {
	// 	user.Password = request.Qty
	// }

	if request.Image != "" {
		user.Image = request.Image
	}

	if request.Price != "" {
		user.Price = request.Price
	}

	data, err := h.TripsRepository.EditTrips(user, ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrips(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrips) DeleteTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.TripsRepository.GetTrips(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripsRepository.DeleteTrips(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrips(data)}
	json.NewEncoder(w).Encode(response)
}
