package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Users []interface{} `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Points   int    `json:"points"`
}

type Users struct {
	Users []interface{} `json:"users"`
}

var userData Users

func main() {
	// initial limit and offset
	limit := 500000 // a higher number will result in a 500 internal server error on the external API
	offset := 0     // the offset is the starting point of the data to be retrieved

	// total number of users
	total := 1453578 // Retrieved from https://api.streamelements.com/kappa/v2/points/592b0d5610b3f73ce98ace4c/alltime. Date retrieved: 2022-11-28

	userDataRetrieved := 0 // keeps track of the number of users retrieved

	//header
	fmt.Print("(c) 1Nukez 2022\n")
	fmt.Print("Retrieving data...\n")
	for userDataRetrieved < total {

		body := Get(limit, offset)
		var data Response
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
		}

		AddUsers(data.Users)

		userDataRetrieved += len(data.Users)
		offset += limit
	}

	fmt.Printf("\n%d entries have been retrieved!\n", userDataRetrieved)

	file, _ := json.MarshalIndent(userData, "", " ")

	//write to file
	_ = ioutil.WriteFile("users.json", file, 0644)

	//keep application alive until user input
	fmt.Print("Press ENTER to exit...")
	fmt.Scanln()

}

// loops through the users in the interface and appends them to the Users struct
func AddUsers(data []interface{}) {
	for _, user := range data {
		userData.Users = append(userData.Users, User{
			Username: user.(map[string]interface{})["username"].(string),
			Points:   int(user.(map[string]interface{})["points"].(float64)),
		})
	}
}

// Get returns the body of the response from the API call.
// It takes in the limit and offset as parameters.
func Get(limit int, offset int) []byte {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.streamelements.com/kappa/v2/points/592b0d5610b3f73ce98ace4c/alltime?limit=%d&offset=%d", limit, offset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: ", resp.StatusCode)
	} else {
		fmt.Print(resp.Status + " - " + "A partition of data has been successfully retrieved from the API!\n")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
