package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type FlashPoint struct {
	LabelText string `json:labelText`
	Id        int    `json:id`
}

func main2() {

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

	new := make([]FlashPoint, 0, 2000)
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

	result := make(map[string]int, 0)
	max := -100
	for _, v := range old {
		if _, ok := result[strings.Trim(v.LabelText, "\xef\xb8\x8f")]; ok {
			continue
		}
		result[v.LabelText] = v.Id
		if v.Id > max {
			max = v.Id
		}
	}

	i := max + 1
	for _, v := range new {
		if _, ok := result[strings.Trim(v.LabelText, "\xef\xb8\x8f")]; ok {
			continue
		}
		result[v.LabelText] = i
		i++
	}

	js, _ := json.Marshal(result)

	fmt.Println("-------")

	fmt.Println(string(js))

}

func main() {

	cancelCtx1, cancelFun1 := context.WithCancel(context.Background())
	valueCtx := context.WithValue(cancelCtx1, "key", "val")
	cancelCtx2, _ := context.WithCancel(valueCtx)

	go func() {
		time.Sleep(5 * time.Second)
		cancelFun1()
	}()

	select {
	case <-cancelCtx2.Done():
		fmt.Println("cancelCtx2 is done")
	}

}
