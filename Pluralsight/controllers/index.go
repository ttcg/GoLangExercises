package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	customerCtrl := newCustomerController()

	http.Handle("/customers", *customerCtrl)
	http.Handle("/customers/", *customerCtrl)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
