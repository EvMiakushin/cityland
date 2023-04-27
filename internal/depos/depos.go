package depos

import (
	"bufio"
	"cityland/internal/city"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CityBase struct {
	MapCity map[int]*city.City
}

func NewCityBase(filer string) *CityBase {

	fmt.Printf("\n### Read file: %s ###\n\n", filer)
	file, _ := os.Open(filer)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	ncity := new(CityBase)
	ncity.MapCity = make(map[int]*city.City)

	var str []string
	var text string

	for fileScanner.Scan() {
		text = fileScanner.Text()
		str = strings.Split(text, ",")

		id, _ := strconv.Atoi(str[0])
		name := str[1]
		region := str[2]
		district := str[3]
		population, _ := strconv.Atoi(str[4])
		foundation, _ := strconv.Atoi(str[5])

		ncity.MapCity[id] = &city.City{
			Id:         id,
			Name:       name,
			Region:     region,
			District:   district,
			Population: population,
			Foundation: foundation,
		}
	}
	return ncity
}

func (cb *CityBase) SaveCityBase(filer string) {
	csvFile, err := os.Create(filer)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, usance := range cb.MapCity {
		var row []string
		row = append(row, strconv.Itoa(usance.Id))
		row = append(row, usance.Name)
		row = append(row, usance.Region)
		row = append(row, usance.District)
		row = append(row, strconv.Itoa(usance.Population))
		row = append(row, strconv.Itoa(usance.Foundation))
		writer.Write(row)
	}
	writer.Flush()

}
