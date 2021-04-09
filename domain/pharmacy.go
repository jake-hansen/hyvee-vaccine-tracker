package domain

type PhoneNumber string
type PharmacyID	string

type Pharmacy struct {
	ID	PharmacyID
	Name string
	Address Address
	PhoneNumber PhoneNumber
	VaccinationsAvailable	bool
}

type Address struct {
	Line1	string
	Line2	string
	City	string
	State	string
	Zip		int
}
