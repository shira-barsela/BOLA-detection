package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LogData struct {
	Req struct {
		URL        string            `json:"url"`
		QSParams   string            `json:"qs_params"`
		Headers    map[string]string `json:"headers"`
		ReqBodyLen int               `json:"req_body_len"`
	} `json:"req"`
	Rsp struct {
		StatusClass string `json:"status_class"`
		RspBodyLen  int    `json:"rsp_body_len"`
	} `json:"rsp"`
}

func readFile(fileName string) LogData {
	// read the file
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshal the data into a LogData structure
	var logData LogData
	err = json.Unmarshal(content, &logData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return logData
}

func potentialAttack(logData LogData) bool {
	message := ""

	// check the StatucClass
	if logData.Rsp.StatusClass == "4xx" {
		message += "The status class indicate of a client error, may cause by a BOLA attack\n"
	}

	// more checks?

	if message != "" {
		fmt.Print(message)
		return true
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <log_file_path>")
	}

	filePath := os.Args[1]

	logData := readFile(filePath)
	potentialAttack(logData)
}
