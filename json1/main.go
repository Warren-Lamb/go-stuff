package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Number string `json:"dork"`
}

func main() {
	text := `{"people": [{"crafts": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "crafty": "goofy", "dork":"fart"}`
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)
	fmt.Println(people1, textBytes)
	for i := 1;i <= len(people1.Number);i++ {
		
	fmt.Println(people1.Number)
	}
}
