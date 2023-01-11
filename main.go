package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	emp := &modeldisini{}
	e, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}
