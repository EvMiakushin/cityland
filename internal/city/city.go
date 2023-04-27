package city

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population int    `json:"population"`
	Foundation int    `json:"foundation"`
}

type CityJson struct {
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population int    `json:"population"`
	Foundation int    `json:"foundation"`
}

type PopJson struct {
	From int `json:"from"`
	To   int `json:"to"`
}
