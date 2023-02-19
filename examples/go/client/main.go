package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	cep := "08226021"
	url := fmt.Sprintf("http://localhost:8080/api/v1/%s", url.QueryEscape(cep))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http.NewRequest:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do:", err)
		return
	}

	defer resp.Body.Close()
	var code int = resp.StatusCode

	if code == 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("io.ReadAll:", err)
			return
		}
		fmt.Println(string(b))
	}
	fmt.Println("code:", code)

}
