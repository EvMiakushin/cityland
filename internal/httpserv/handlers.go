package httpserv

import (
	"cityland/internal/city"
	"cityland/pkg/mod/github.com/go-chi/chi/v5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetCityIdHandler(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idS)

	city, err := h.sv.CitHand.GetCityId(id)
	if err != nil {
		h.message(w, err.Error(), 404)
		return
	}

	line := fmt.Sprintf("City %v %s %s %s %v %v",
		id, city.Name, city.Region, city.District, city.Population, city.Foundation)
	h.message(w, line, 200)

}

func (h *Handler) CreateCityHandler(w http.ResponseWriter, r *http.Request) {
	var nc city.CityJson
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&nc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, _ := h.sv.CreateCity(nc)

	line := fmt.Sprintf("New City %v created: Name %s Region %s District %s Population %v Foundation %v",
		id, nc.Name, nc.Region, nc.District, nc.Population, nc.Foundation)
	h.message(w, line, http.StatusCreated)
}

func (h *Handler) DeleteCityHandler(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		h.message(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.sv.CitHand.DelCityId(id)
	if err != nil {
		h.message(w, err.Error(), http.StatusInternalServerError)
		return
	}

	line := fmt.Sprintf("City ID: %v  have been DELETED",
		id)
	h.message(w, line, http.StatusOK)

}

func (h *Handler) SetIdPopulationHandler(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		h.message(w, err.Error(), http.StatusBadRequest)
		return
	}

	var nc city.CityJson
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&nc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	city, err := h.sv.CitHand.SetIdPop(id, nc.Population)
	if err != nil {
		h.message(w, err.Error(), 404)
		return
	}

	line := fmt.Sprintf("City %v %s %s %s %v %v",
		id, city.Name, city.Region, city.District, city.Population, city.Foundation)
	h.message(w, line, 200)

}

func (h *Handler) GetDistrictHandler(w http.ResponseWriter, r *http.Request) {
	district := chi.URLParam(r, "district")
	district = firstUp(district)

	dst, err := h.sv.CitHand.GetDistrict(district)
	if err != nil {
		h.message(w, err.Error(), 404)
		return
	}
	line := fmt.Sprintf("This district << %s >> includes the following cities: %v", district, dst)
	h.message(w, line, 200)

}

func (h *Handler) GetRegionHandler(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	region = firstUp(region)
	dst, err := h.sv.CitHand.GetRegion(region)
	if err != nil {
		h.message(w, err.Error(), 404)
		return
	}
	line := fmt.Sprintf("This region << %s >> includes the following cities: %v", region, dst)
	h.message(w, line, 200)

}

func (h *Handler) PopulationRangeHandler(w http.ResponseWriter, r *http.Request) {
	var pops city.PopJson
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&pops); err != nil {
		fmt.Println("StatusBadRequest1", err)
		http.Error(w, err.Error(), 200)
		return
	}

	pop, err := h.sv.PopRange(pops.From, pops.To)
	if err != nil {
		h.message(w, err.Error(), 404)
		return
	}
	line := fmt.Sprintf("The following cities have populations in this range: %v", pop)
	h.message(w, line, 200)

}

func (h *Handler) FoundationRangeHandler(w http.ResponseWriter, r *http.Request) {
	var founds city.PopJson
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&founds); err != nil {
		fmt.Println("StatusBadRequest1", err)
		http.Error(w, err.Error(), 200)
		return
	}

	found, err := h.sv.FoundRange(founds.From, founds.To)
	if err != nil {
		h.message(w, err.Error(), 404)
		return
	}
	line := fmt.Sprintf("The following cities were founded during this period.: %v", found)
	h.message(w, line, 200)

}

func (h *Handler) message(w http.ResponseWriter, line string, code int) {
	w.WriteHeader(code)
	log.Println(line)
	w.Write([]byte(line))
}

func firstUp(text string) (textUp string) {
	textUp = fmt.Sprintf(strings.Title(strings.ToLower(text)))
	return
}
