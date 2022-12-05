package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestReadFile(m *testing.T) {
	resp, err := ReadExcel()
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonStr, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	} else {
		fmt.Println(string(jsonStr))
		_ = ioutil.WriteFile("response.json", jsonStr, 0644)

	}
}
