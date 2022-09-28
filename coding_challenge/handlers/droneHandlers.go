package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lazyspell/drone_test/models"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/imports"

	"github.com/kelvins/geocoder"
)

var ctx = context.Background()
var starting = 2

//Initialize data coming from the CSV file.  Creating and returning a dataframe.
func runs() *dataframe.DataFrame {
	csvfile, err := os.Open("./data/drone_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	df, err := imports.LoadFromCSV(ctx, csvfile)
	if err != nil {
		log.Fatal(err)
	}

	s2 := dataframe.NewSeriesFloat64("completed", nil, 0, 0, 0, 0)

	dff := dataframe.NewDataFrame(df.Series[0], df.Series[1], s2)

	return dff

}

//Will send out coordinates based off the address given in the dataframe.
//The most optimal path that i decided was run based off a questack.  First order that comes in should be the should be the first order to be completed and delivered.
func NewDestination(w http.ResponseWriter, r *http.Request) {
	geocoder.ApiKey = "Google API Key Here"
	var inputAddress models.Address
	df := runs()

	fmt.Println(starting)
	current := fmt.Sprintf("%v", df.Series[2].Value(starting))
	current_order, _ := strconv.Atoi(current)

	if current_order == 1 {
		fmt.Println("went in here")
		fmt.Println(starting)
		starting = starting + 1
	}

	address := fmt.Sprintf("%v", df.Series[1].Value(starting))

	arrayAddress := strings.Split(address, ",")

	stateZip := strings.Split(arrayAddress[2], " ")
	numberStreet := strings.Split(arrayAddress[0], " ")

	street_name := ""
	for _, j := range numberStreet[1:] {

		street_name = street_name + j + " "
	}

	inputAddress.Number, _ = strconv.Atoi(numberStreet[0])
	inputAddress.Street = street_name
	inputAddress.City = arrayAddress[1]
	inputAddress.State = stateZip[1]
	inputAddress.ZipCode = stateZip[2]

	geo_address := geocoder.Address{
		Street:  inputAddress.Street,
		Number:  inputAddress.Number,
		City:    inputAddress.City,
		State:   inputAddress.State,
		Country: "United States",
	}

	fmt.Println(inputAddress)

	location, err := geocoder.Geocoding(geo_address)
	if err != nil {
		log.Fatal(err)
	}

	df.UpdateRow(starting, nil, map[string]interface{}{
		"completed": 1,
	})

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)

}
