package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// var jsonString = `{"Name": "Moh. Irsad", "Age" : 22}`
	// var jsonData = []byte(jsonString)

	// Decode JSON ke variable objek struct

	// var data user

	// var err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("User : ", data.FullName)
	// fmt.Println("Age : ", data.Age)

	// Decode JSON map[string]interface{} dan interface{}

	// var data1 map[string]interface{}
	// json.Unmarshal(jsonData, &data1)

	// fmt.Println("User : ", data1["Name"])
	// fmt.Println("Age : ", data1["Age"])

	// var data2 interface{}
	// json.Unmarshal(jsonData, &data2)

	// var decodeData = data2.(map[string]interface{})
	// fmt.Println("User : ", decodeData["Name"])
	// fmt.Println("Age : ", decodeData["Age"])

	// Decode Array JSON ke array objek

	// var jsonString = `[

	// {"Name":"Moh Irsad","Age":22},
	// {"Name":"Riyan Maulana", "Age":35}
	// ]`

	// var data []user

	// var err = json.Unmarshal([]byte(jsonString), &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("User1 :", data[0].Age)
	// fmt.Println("User2 :", data[1].Age)

	// Encode Objek ke json string

	var object = []user{{"Moh. Irsad", 22}, {"Riyan Maulana", 35}}
	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)

}
