package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"code.byted.org/gopkg/logs"
)

func sendRequest(bdDid string) {
	url := "https://api16-access-ttp.tiktokpangle.us/monitor/collect/c/session?version_code=6400&device_platform=android&aid=10000001"
	payload := fmt.Sprintf(`{
    "apm_id": "20000001",
    "header":
    {
        "custom":
        {
            "sdk_version": "6.4.0.0",
            "host_app_id": "8025677"
        },
        "os": "Android",
        "os_version": "12",
        "device_model": "SM-A315G",
        "device_brand": "samsung",
        "sdk_version_name": "0.0.5",
        "aid": "10000000",
        "update_version_code": 6400,
        "bd_did": "%s"
    },
    "local_time": 1734000202200,
    "launch":
    [
        {
            "local_time_ms": 1734000202200
        }
    ]
}`, bdDid)

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 10; SM-G8870 Build/QP1A.190711.020)")
	req.Header.Set("Host", "api16-access-ttp.tiktokpangle.us")
	req.Header.Set("transfer-param", "{\"gaid\":\"925d7fe4-5781-46f2-a0d0-e1584eaff1ff\"}")
	req.Header.Set("cypher", "0")
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	logs.Info("status_code=%d, log_id= %s , body=%s", resp.StatusCode, resp.Header.Get("X-Tt-Logid"), body)
}

func sendOnce() {
	sendRequest("8744715169892138")
	time.Sleep(3 * time.Second)
}

func sendRepeatedly() {

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	go doSendRepeatedly(ticker)
	time.Sleep(10 * time.Second)
	// time.Sleep(60 * time.Minute)
}

func doSendRepeatedly(ticker *time.Ticker) {
	n := 0
	for range ticker.C {
		var bdDid string
		for i := 0; i < 4; i++ {
			logs.Info("sending request [%d]", n)
			bdDid = strconv.Itoa(rand.Int())
			sendRequest(bdDid)
			n++
		}
	}
	logs.Info("doSendRepeatedly is Over.")
}

func main() {
	// sendOnce()
	sendRepeatedly()
	// time.Sleep(30 * time.Second)
}
