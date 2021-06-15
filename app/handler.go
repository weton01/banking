package app

import (
	"encoding/json"
	"net/http"

	"github.com/weton01/banking/go_project/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
