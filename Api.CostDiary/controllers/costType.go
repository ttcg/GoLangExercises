package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/ttcg/GoLangExercies/Api.CostDiary/managers/costtypemgr"
)

type costTypeController struct {
	costTypeIDPattern *regexp.Regexp
}

func (ctrl costTypeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if strings.EqualFold(r.URL.Path, "/costtypes") || strings.EqualFold(r.URL.Path, "/costtypes/") {
		switch r.Method {
		case http.MethodGet:
			ctrl.getAll(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := ctrl.costTypeIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := uuid.Parse(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			ctrl.getByID(id, w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func newCostTypeController() *costTypeController {
	return &costTypeController{
		costTypeIDPattern: regexp.MustCompile(`^/costtypes/([0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12})/?`),
	}
}

func (ctrl *costTypeController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(costtypemgr.GetCostTypes(), w)
}

func (ctrl *costTypeController) getByID(ID uuid.UUID, w http.ResponseWriter, r *http.Request) {
	costType, err := costtypemgr.GetCostTypeByID(ID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		writeTextResponse(w, err.Error())
		return
	}

	encodeResponseAsJSON(costType, w)
}
