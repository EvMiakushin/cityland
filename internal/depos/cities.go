package depos

import (
	"cityland/internal/city"
)

type Cities interface {
	GetCityId(int) (*city.City, error)
	CreateCity(nc city.CityJson) (id int, err error)
	DelCityId(int) error
	SetIdPop(int, int) (*city.City, error)
	GetDistrict(district string) ([]string, error)
	GetRegion(region string) ([]string, error)
	PopRange(int, int) ([]string, error)
	FoundRange(int, int) ([]string, error)
}

type CityStore struct {
	Cities
}

func NewCityStor(cb *CityBase) *CityStore {
	return &CityStore{
		Cities: NewCityListDB(cb),
	}
}
