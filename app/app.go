package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weton01/banking/domain"
	"github.com/weton01/banking/service"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post new user")
}
