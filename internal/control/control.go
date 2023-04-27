package control

import (
	"cityland/internal/city"
	"cityland/internal/depos"
)

type CitHand interface {
	GetCityId(id int) (*city.City, error)
	CreateCity(nc city.CityJson) (int, error)
	DelCityId(int) error
	SetIdPop(int, int) (*city.City, error)
	GetDistrict(district string) ([]string, error)
	GetRegion(region string) ([]string, error)
	PopRange(int, int) ([]string, error)
	FoundRange(int, int) ([]string, error)
}

type CityHandler struct {
	CitHand
}

func NewCityHand(repos *depos.CityStore) *CityHandler {
	return &CityHandler{
		CitHand: NewCityControl(repos.Cities),
	}
}
