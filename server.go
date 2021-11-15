package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type student struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
}

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error can't get http")

	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	std := &student{"rahul", "Physics", 71}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(std); err != nil {
		log.Fatalf("error cant encode%s", err)
	}
	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)

	if err != nil {

		log.Fatalf("cant post %s", err)

	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("3. Performing Http Put...")
	std2 := student{"Shubham", "Maths", 98}
	jsonReq, err := json.Marshal(std2)
	req, err := http.NewRequest(http.MethodPut, "https://httpbin.org/put", bytes.NewBuffer(jsonReq))

	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp2.Body.Close()
	io.Copy(os.Stdout, resp2.Body)

	fmt.Println(" Performing Http Delete")

	req2, err := http.NewRequest(http.MethodDelete, "https://httpbin.org/delete", bytes.NewBuffer(jsonReq))
	client1 := &http.Client{}
	resp3, err := client1.Do(req2)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp3.Body.Close()
	io.Copy(os.Stdout, resp3.Body)

}
