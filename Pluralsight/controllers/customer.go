package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ttcg/golangExercises/pluralsight/models"
)

type customerController struct {
	customerIDPattern *regexp.Regexp
}

func (customerCtrl customerController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path == "/customers" || r.URL.Path == "/customers/" {
		switch r.Method {
		case http.MethodGet:
			customerCtrl.getAll(w, r)
		case http.MethodPost:
			customerCtrl.addCustomer(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	} else {
		matches := customerCtrl.customerIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			customerCtrl.getCustomerByID(id, w)
		case http.MethodPut:
			customerCtrl.updateCustomerByID(id, w, r)
		case http.MethodDelete:
			customerCtrl.deleteCustomer(id, w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func newCustomerController() *customerController {
	return &customerController{
		customerIDPattern: regexp.MustCompile(`^/customers/(\d+)/?`),
	}
}

func (customerCtrl *customerController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetCustomers(), w)
}

func (customerCtrl *customerController) getCustomerByID(id int, w http.ResponseWriter) {
	customer, err := models.GetCustomerByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errStr := fmt.Sprint("Could not find a customer: ", id)
		writeTextResponse(w, errStr)
		return
	}

	encodeResponseAsJSON(customer, w)
}

func (customerCtrl *customerController) addCustomer(w http.ResponseWriter, r *http.Request) {
	customer, err := customerCtrl.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprint("Could not parse Customer Object: ", err)
		writeTextResponse(w, errStr)
		return
	}

	addedCustomer, err := models.AddCustomer(customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprint("Could not add a customer: ", err)
		writeTextResponse(w, errStr)
		return
	}

	encodeResponseAsJSON(addedCustomer, w)
}

func (customerCtrl *customerController) updateCustomerByID(id int, w http.ResponseWriter, r *http.Request) {
	customer, err := customerCtrl.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprint("Could not parse Customer Object: ", err)
		writeTextResponse(w, errStr)
		return
	}

	if id != customer.ID {
		w.WriteHeader(http.StatusInternalServerError)
		writeTextResponse(w, "The Id in URL and CustomerID are not the same")
		return
	}

	_, err = models.GetCustomerByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errStr := fmt.Sprint("Could not find a customer: ", id)
		writeTextResponse(w, errStr)
		return
	}

	_, err = models.UpdateCustomerByID(customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprint("Could not update a customer: ", err)
		writeTextResponse(w, errStr)
		return
	}

	encodeResponseAsJSON(customer, w)
}

func (customerCtrl *customerController) deleteCustomer(id int, w http.ResponseWriter, r *http.Request) {

	err := models.DeleteCustomerByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprint("Could not delete a customer: ", err)
		writeTextResponse(w, errStr)
		return
	}
}

func (customerCtrl *customerController) parseRequest(r *http.Request) (models.Customer, error) {
	dec := json.NewDecoder(r.Body)
	var customer models.Customer
	err := dec.Decode(&customer)

	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func writeTextResponse(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}
