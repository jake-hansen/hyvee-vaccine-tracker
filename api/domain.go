package api

type Variables struct {
	Radius int `json:"radius"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Pharmacy struct {
	Distance float64 `json:"distance"`
	Location	Location `json:"location"`
}

type Location struct {
	Nickname	string `json:"nickname"`
	PhoneNumber	string `json:"phoneNumber"`
	IsCovidVaccineAvailable	bool `json:"isCovidVaccineAvailable"`
	Address	Address `json:"address"`
}

type Address struct {
	Line1	string `json:"line1"`
	Line2	string `json:"line2"`
	City	string `json:"city"`
	State	string `json:"state"`
	Zip		string `json:"zip"`
}

type Data struct {
	SearchPharmaciesNearPoint []Pharmacy `json:"searchPharmaciesNearPoint"`
}

type ResponseWrapper struct {
	Data Data `json:"data"`
}