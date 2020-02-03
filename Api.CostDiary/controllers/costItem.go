package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/ttcg/GoLangExercies/Api.CostDiary/models"

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
		case http.MethodPost:
			ctrl.addCostItem(w, r)
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

func (ctrl *costItemController) addCostItem(w http.ResponseWriter, r *http.Request) {
	costItem, err := ctrl.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprint("Could not parse Cost Item Object: ", err)
		writeTextResponse(w, errStr)
		return
	}

	costItem, err = costitemmgr.AddCostItem(costItem)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeTextResponse(w, err.Error())
		return
	}

	encodeResponseAsJSON(costItem, w)
}

func (ctrl *costItemController) parseRequest(r *http.Request) (models.CostItem, error) {
	dec := json.NewDecoder(r.Body)
	var costItem models.CostItem
	err := dec.Decode(&costItem)

	if err != nil {
		return models.CostItem{}, err
	}

	return costItem, nil
}
