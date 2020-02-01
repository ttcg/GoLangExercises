package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers : Register controllers
func RegisterControllers() {
	costTypeCtrl := newCostTypeController()

	http.Handle("/costtype", *costTypeCtrl)
	http.Handle("/costtypes/", *costTypeCtrl)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
