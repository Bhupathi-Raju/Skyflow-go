package networkcall

import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
	"bytes"
)

func MakeGetRequest(requestUrl string, authorizationToken string){
	fmt.Println("Performing Http Post request for")

	client := http.Client{
	}

	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("content-type", "application/json")
	request.Header.Set("X-SKYFLOW-APP-ID", "c0e915a9904011ea95ba2e321592fd49")
	request.Header.Set("X-SKYFLOW-ORG-ID", "c117973f904011ea95ba2e321592fd49")
	request.Header.Set("X-SKYFLOW-APP-SECRET", "MpdB0NTs8KT3oBG7rhSTwB5kJ1c=")
	request.Header.Set("Authorization", authorizationToken)

	if err != nil{
		log.Fatalln(err)
	}

	resp, err := client.Do(request)

	if err != nil{
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		log.Fatalln(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println("API response as String: \n" + bodyString)
}


func MakePostRequest(requestUrl, authorizationToken string, requestBody []byte){
	fmt.Println("Performing Http Post request for")

	client := http.Client{
	}

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("content-type", "application/json")
	request.Header.Set("X-SKYFLOW-APP-ID", "c0e915a9904011ea95ba2e321592fd49")
	request.Header.Set("X-SKYFLOW-ORG-ID", "c117973f904011ea95ba2e321592fd49")
	request.Header.Set("X-SKYFLOW-APP-SECRET", "MpdB0NTs8KT3oBG7rhSTwB5kJ1c=")
	request.Header.Set("Authorization", authorizationToken)

	resp, err := client.Do(request)
	
	if err != nil{
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		log.Fatalln(err)
	}

	log.Println(string(body))
}