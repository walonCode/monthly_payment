package main

import (
	"errors"
	"log"
	"net/http"
	"encoding/json"
)

const RATE float64 = 12

func calculateInterest(amount, year float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}

	interest := (amount * RATE * year) / 100
	return interest, nil
}

func getMonthlyPayment(amount, year float64) (float64, error) {
	interest, err := calculateInterest(amount, year)
	if err != nil {
		return 0, err
	}

	totalAmount := amount + interest
	payment := totalAmount / (year * 12)
	return payment, nil
}

type PaymentRequest struct{
	Amount float64 `json:"amount"`
	Year float64  `json:"year"`
}

type Message struct{
	Ok bool `json:"ok"`
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body := Message{
		Ok:true,
		Message:"Welcome to my server",
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&body)
}

func monthlyPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//getting the value from a json
	var req PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil{
		http.Error(w, "invalid data input", http.StatusBadRequest)
		return
	}


	payment, err := getMonthlyPayment(req.Amount, req.Year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	type Message struct {
		Value any
	}

	type PaymentMessage struct {
		Ok bool `json:"ok"`
		Message string `json:"message"`
		Data  Message `json:"data"`
	}

	body := PaymentMessage{
		Ok: true,
		Message: "monthly payment",
		Data : Message{
			Value: payment,
		},
	}
	json.NewEncoder(w).Encode(&body)
}



func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/monthly-payment", monthlyPaymentHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
