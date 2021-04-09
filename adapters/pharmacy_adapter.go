package adapters

import (
	"github.com/jake-hansen/hyvee-vaccine-search/api"
	"github.com/jake-hansen/hyvee-vaccine-search/domain"
	"strconv"
)

func HyVeePharmacyToDomainPharmacy(pharmacy api.Pharmacy) domain.Pharmacy {
	dPharmacy := domain.Pharmacy{
		Name:                  pharmacy.Location.Nickname,
		Address:               HyVeeAddressToDomainAddress(pharmacy.Location.Address),
		PhoneNumber:           domain.PhoneNumber(pharmacy.Location.PhoneNumber),
		VaccinationsAvailable: pharmacy.Location.IsCovidVaccineAvailable,
	}

	return dPharmacy
}

func HyVeeAddressToDomainAddress(address api.Address) domain.Address {
	zip, _ := strconv.Atoi(address.Zip)

	dAddress := domain.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City:  address.City,
		State: address.State,
		Zip:   zip,
	}
	return dAddress
}