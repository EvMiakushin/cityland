package control

import (
	"cityland/internal/city"
	"cityland/internal/depos"
)

type CityControl struct {
	repo depos.Cities
}

func NewCityControl(repo depos.Cities) *CityControl {
	return &CityControl{repo: repo}
}

func (c CityControl) GetCityId(id int) (*city.City, error) {
	return c.repo.GetCityId(id)
}

func (c CityControl) CreateCity(nc city.CityJson) (int, error) {
	return c.repo.CreateCity(nc)
}

func (c CityControl) DelCityId(id int) error {
	return c.repo.DelCityId(id)
}

func (c CityControl) SetIdPop(id int, pop int) (*city.City, error) {
	return c.repo.SetIdPop(id, pop)
}

func (c CityControl) GetDistrict(district string) ([]string, error) {
	return c.repo.GetDistrict(district)
}

func (c CityControl) GetRegion(region string) ([]string, error) {
	return c.repo.GetRegion(region)
}

func (c CityControl) PopRange(from int, to int) ([]string, error) {
	return c.repo.PopRange(from, to)
}

func (c CityControl) FoundRange(from int, to int) ([]string, error) {
	return c.repo.FoundRange(from, to)
}
