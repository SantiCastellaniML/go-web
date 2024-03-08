package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println(sendRequest("http://localhost:8080/greetings", "POST", `{"nombre": "Juan", "apellido": "Perez"}`))
	fmt.Println(NewSendRequest("http://localhost:8080/greetings", "POST", `{"nombre": "Juan", "apellido": "Perez"}`))
}

func NewSendRequest(url string, method string, body string) (response string, err error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response = string(bodyBytes)

	defer resp.Body.Close()

	return response, nil
}

func sendRequest(url string, method string, body string) (response string, err error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)

	response = string(responseBytes)

	return response, nil
}
