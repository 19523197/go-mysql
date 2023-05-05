package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// list := belajar_golang_mysql.GetAll()
	// fmt.Println(list)

	var jsonString = `{"data" :[
    	{"Name": "john wick", "Age": 27},
    	{"Name": "ethan hunt", "Age": 32}
	]
	}`
	var response = make(map[string]interface{})
	var data []User
	var err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range response {
		// Convert map to json string
		jsonStr, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
		}

		// Convert json string to struct

		if err := json.Unmarshal(jsonStr, &data); err != nil {
			fmt.Println(err)
		}

	}
	// log.Println(response)
	fmt.Println(data)
	// fmt.Println("user 2:", data.Users[1].Name)

}

type User struct {
	Name string
	Age  int
}

type UserResponse struct {
	Users []User
}
