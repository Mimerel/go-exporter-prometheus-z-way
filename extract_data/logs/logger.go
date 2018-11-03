package logs

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func toElastic(url string, host string, level string, message string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	postingUrl := url + "/" + host + "/logs"

	n := int64(time.Now().Unix())
	moment := strconv.FormatInt(n, 10)

	json := "{ \"Level\": \"" + level + "\", \"message\" : \"" + message + "\", \"timestamp\":" + moment + "}"

	_, err := client.Post(postingUrl, "application/json" ,
		bytes.NewBuffer([]byte(json)))
	if err != nil {
		fmt.Printf("Failed to log error message. Post failed")
	}
}

func Info(url string, host string, message string) {
	if url != "" {
		toElastic(url, host, "Info", message)
	} else {
		fmt.Printf("Info : %s \n",  message)
	}
}

func Error(url string, host string, message string) {
	if url != "" {
		toElastic(url, host, "Error", message)
	} else {
		fmt.Printf("Error : %s \n",  message)
	}
}