package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers : Register controllers
func RegisterControllers() {
	costTypeCtrl := newCostTypeController()
	costItemCtrl := newCostItemController()

	http.Handle("/costtype", *costTypeCtrl)
	http.Handle("/costtypes/", *costTypeCtrl)

	http.Handle("/costitem", *costItemCtrl)
	http.Handle("/costitems/", *costItemCtrl)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func writeTextResponse(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}
