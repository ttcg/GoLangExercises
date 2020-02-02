package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/ttcg/GoLangExercies/Api.CostDiary/managers/costitemmgr"
)

type costItemController struct {
	costItemIDPattern *regexp.Regexp
}

func (ctrl costItemController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if strings.EqualFold(r.URL.Path, "/costitems") || strings.EqualFold(r.URL.Path, "/costitems/") {
		switch r.Method {
		case http.MethodGet:
			ctrl.getAll(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := ctrl.costItemIDPattern.FindStringSubmatch(r.URL.Path)
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

func newCostItemController() *costItemController {
	return &costItemController{
		costItemIDPattern: regexp.MustCompile(`^/costitems/([0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12})/?`),
	}
}

func (ctrl *costItemController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(costitemmgr.GetCostItems(), w)
}

func (ctrl *costItemController) getByID(ID uuid.UUID, w http.ResponseWriter, r *http.Request) {
	costType, err := costitemmgr.GetCostItemByID(ID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		writeTextResponse(w, err.Error())
		return
	}

	encodeResponseAsJSON(costType, w)
}
