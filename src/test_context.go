package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FlashPoint struct {
	LabelText string `json:labelText`
	Id        int    `json:id`
}

func main() {

	old := make([]FlashPoint, 0)
	file, err := os.Open("./old.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(content, &old)
	if err != nil {
		fmt.Println(err)
		return
	}

	new := make([]FlashPoint, 0)
	file, err = os.Open("./new.json")
	if err != nil {
		panic(err)
	}

	content, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(content, &new)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(old)
	fmt.Println(new)

}
