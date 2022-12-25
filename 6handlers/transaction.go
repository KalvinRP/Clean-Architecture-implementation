package handlers

import (
	models "dewetour/1models"
	repositories "dewetour/4repositories"
	dto "dewetour/5dto/result"
	transactiondto "dewetour/5dto/transaction"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trans, err := h.TransactionRepository.FindTrans()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trans}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trans, err := h.TransactionRepository.GetTrans(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransResponse(trans)}
	json.NewEncoder(w).Encode(response)
}

func convertTransResponse(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		TripsID:  u.TripsID,
		Trips:    u.Trips,
		Users:    u.Users,
		TotalPrc: u.TotalPrc,
		Status:   u.Status,
		// Name:           u.Name,
		// Desc:           u.Desc,
		// Price:          u.Price,
		// Eat:            u.Eat,
		// Quota:          u.Quota,
		// Image:          u.Image,
		// Country:        u.Country,
	}
}

func (h *handlerTransaction) MakeTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	request := new(transactiondto.TransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trans := models.Transaction{
		TripsID:  request.TripsID,
		UsersID:  request.UsersID,
		TotalPrc: request.TotalPrc,
		Status:   request.Status,
	}

	transaction, err := h.TransactionRepository.MakeTrans(trans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransResponse(transaction)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) EditTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(transactiondto.UpdateTransactionRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ID, _ := strconv.Atoi(mux.Vars(r)["id"])

	trans, err := h.TransactionRepository.GetTrans(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// trans := models.Transaction{}

	if request.TripsID != 0 {
		trans.TripsID = request.TripsID
	}

	if request.UsersID != 0 {
		trans.UsersID = request.UsersID
	}

	if request.TotalPrc != 0 {
		trans.TotalPrc = request.TotalPrc
	}

	if request.Status != "" {
		trans.TotalPrc = request.TotalPrc
	}

	data, err := h.TransactionRepository.EditTrans(trans, ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) DeleteTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trans, err := h.TransactionRepository.GetTrans(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransactionRepository.DeleteTrans(trans, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransResponse(data)}
	json.NewEncoder(w).Encode(response)
}
