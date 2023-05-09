package main

import (
	"fmt"
	"time"

	"github.com/kamildrazkiewicz/go-flow"
)

func main() {
	f1 := func(r map[string]interface{}) (interface{}, error) {
		fmt.Println("f1 input:", r)
		fmt.Println("f1 start")
		time.Sleep(time.Millisecond * 5000)
		fmt.Println("f1 end")
		return 1, nil
	}

	f2 := func(r map[string]interface{}) (interface{}, error) {
		fmt.Println("f2 input:", r)
		fmt.Println("f2 start")
		time.Sleep(time.Millisecond * 2000)
		fmt.Println("f2 end")
		return 2, nil
	}

	f3 := func(r map[string]interface{}) (interface{}, error) {
		fmt.Println("f3 input:", r)
		fmt.Println("f3 start")
		time.Sleep(time.Millisecond * 3000)
		fmt.Println("f3 end")
		return 3, nil
	}

	f4 := func(r map[string]interface{}) (interface{}, error) {
		fmt.Println("f4 input:", r)
		fmt.Println("f4 start")
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("f4 end")
		return nil, nil
	}

	res, err := goflow.New().
		Add("f1", nil, f1).
		Add("f2", []string{"f1"}, f2).
		Add("f3", []string{"f1"}, f3).
		Add("f4", []string{"f2", "f3"}, f4).
		Do()

	fmt.Println(res, err)

}
