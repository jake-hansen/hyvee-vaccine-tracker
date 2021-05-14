package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const HYVEE_URL = "https://www.hy-vee.com"

type HyVeeAPI struct {
	Client *http.Client
}

func (h *HyVeeAPI) GetPharmacies(variables Variables) ([]Pharmacy, error) {
	reqURL := HYVEE_URL + "/my-pharmacy/api/graphql"

	graphReq := &GraphQLRequest{
		OperationName: "SearchPharmaciesNearPointWithCovidVaccineAvailability",
		Variables:     variables,
		Query:         "query SearchPharmaciesNearPointWithCovidVaccineAvailability($latitude: Float!, $longitude: Float!, $radius: Int! = 10) {\n  searchPharmaciesNearPoint(latitude: $latitude, longitude: $longitude, radius: $radius) {\n    distance\n    location {\n      locationId\n      name\n      nickname\n      phoneNumber\n      businessCode\n      isCovidVaccineAvailable\n      covidVaccineEligibilityTerms\n      address {\n        line1\n        line2\n        city\n        state\n        zip\n        latitude\n        longitude\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
	}
	
	requestBody, err := json.Marshal(graphReq)
	if err != nil {
		return nil, NewRequestCreationError(reqURL, err)
	}

	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest(http.MethodPost, reqURL, buffer)

	res, err := h.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(reqURL, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("an error occured closing the body for request with url %s: %s\n", reqURL, err.Error())
		}
	}(req.Body)

	type ResponseWrapper struct {
		Data Data `json:"data"`
	}

	var responseList ResponseWrapper
	err = json.NewDecoder(res.Body).Decode(&responseList)
	if err != nil {
		return nil, NewResponseDecodingError(reqURL, err)
	}

	return responseList.Data.SearchPharmaciesNearPoint, nil
}
