package httpserv

import (
	"cityland/internal/city"
	"cityland/internal/control"
	"cityland/pkg/mod/github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

type Handler struct {
	sv *control.CityHandler
}

func NewHandler(sv *control.CityHandler) *Handler {
	return &Handler{sv: sv}
}

func (h *Handler) Https() *chi.Mux {

	r := chi.NewRouter()

	r.HandleFunc("/", h.hp)
	r.Get("/{id}", h.GetCityIdHandler)
	r.Post("/", h.CreateCityHandler)
	r.Delete("/{id}", h.DeleteCityHandler)
	r.Patch("/population/{id}", h.SetIdPopulationHandler)
	r.Get("/district/{district}", h.GetDistrictHandler)
	r.Get("/region/{region}", h.GetRegionHandler)
	r.Options("/population/range", h.PopulationRangeHandler)
	r.Options("/foundation/range", h.FoundationRangeHandler)

	return r
}

func (h *Handler) hp(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("temp/hp.html")

	data := city.City{Name: "Monreal", Region: "Canada"}

	tmpl.Execute(w, data)
}
