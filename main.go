package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"toolforthi/utils"
)

func main() {
	resp, err := utils.ReadExcel()
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonStr, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
		return
	} else {
		fmt.Println(string(jsonStr))
		_ = ioutil.WriteFile("response.json", jsonStr, 0644)
	}
}
