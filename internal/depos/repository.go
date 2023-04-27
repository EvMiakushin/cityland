package depos

import (
	"cityland/internal/city"
	"fmt"
	"math/rand"
	"regexp"
)

type CityListCB struct {
	cb *CityBase
}

func NewCityListDB(cb *CityBase) *CityListCB {
	return &CityListCB{cb: cb}
}

func (c CityListCB) GetCityId(id int) (*city.City, error) {
	t, ok := c.cb.MapCity[id]
	if ok {
		return t, nil
	} else {
		return nil, fmt.Errorf("City with id= %v not found\n", id)
	}
}

func (c CityListCB) CreateCity(nc city.CityJson) (int, error) {
	id := c.newId()

	c.cb.MapCity[id] = &city.City{
		Id:         id,
		Name:       nc.Name,
		Region:     nc.Region,
		District:   nc.District,
		Population: nc.Population,
		Foundation: nc.Foundation,
	}
	return id, nil
}

func (c CityListCB) DelCityId(id int) error {

	if _, ok := c.cb.MapCity[id]; !ok {
		return fmt.Errorf("City with id= %v not found\n", id)
	}
	delete(c.cb.MapCity, id)
	return nil
}

func (c CityListCB) SetIdPop(id int, pop int) (*city.City, error) {
	t, ok := c.cb.MapCity[id]
	if ok {
		c.cb.MapCity[id].Population = pop
		return t, nil
	} else {
		return nil, fmt.Errorf("City with id= %v not found\n", id)
	}

}

func (c CityListCB) GetDistrict(district string) ([]string, error) {
	cf := make([]string, 0)
	for _, v := range c.cb.MapCity {
		if v.District == district {
			cf = append(cf, v.Name)
		}
	}
	if len(cf) == 0 {
		return nil, fmt.Errorf("This District << %s >> not found\n", district)
	}
	return cf, nil
}

func (c CityListCB) GetRegion(region string) ([]string, error) {
	cf := make([]string, 0)
	rg := regexp.MustCompile(region)

	for _, v := range c.cb.MapCity {
		if rg.MatchString(v.Region) {
			cf = append(cf, v.Name)
		}
	}
	if len(cf) == 0 {
		return nil, fmt.Errorf("This region << %s >> not found\n", region)
	}
	return cf, nil
}

func (c CityListCB) PopRange(from int, to int) ([]string, error) {
	if from > to {
		from, to = to, from
	}
	cf := make([]string, 0)
	for _, v := range c.cb.MapCity {
		if v.Population >= from && v.Population <= to {
			cf = append(cf, v.Name)
		}
	}
	if len(cf) == 0 {
		return nil, fmt.Errorf("There is no city with such a population.\n")
	}
	return cf, nil
}

func (c CityListCB) FoundRange(from int, to int) ([]string, error) {
	if from > to {
		from, to = to, from
	}
	cf := make([]string, 0)
	for _, v := range c.cb.MapCity {
		if v.Foundation >= from && v.Foundation <= to {
			cf = append(cf, v.Name)
		}
	}
	if len(cf) == 0 {
		return nil, fmt.Errorf("There are no cities with such years of foundation.\n")
	}
	return cf, nil
}

func (c CityListCB) newId() int {
	for {
	loop:
		id := (rand.Intn(9998)) + 1
		_, ok := c.cb.MapCity[id]
		if ok {
			goto loop
		} else {
			return id
		}
	}
}
