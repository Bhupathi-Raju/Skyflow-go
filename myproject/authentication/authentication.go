package authentication

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"bytes"
)

type AccessKey struct{
	Bearer string `json:"accessToken"`
}

const authenticationUrl = "https://api.skyflow.dev/v1/auth/token"

func GetAccessToken(userName, password string) string{
	client := http.Client{
	}

	requestBody, err := json.Marshal(map[string]string{
		"username": userName,	
		"password": password,
	})

	request, err := http.NewRequest("POST", authenticationUrl, bytes.NewBuffer(requestBody))

	if err != nil{
		log.Fatalln(err)
	}

	request.Header.Set("X-SKYFLOW-ORG-ID", "c117973f904011ea95ba2e321592fd49")

	resp, err := client.Do(request)

	if err != nil{
		log.Fatalln(err)
	}
	
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("API access token: \n",  bodyBytes)
	var bearer AccessKey
	fmt.Println("API access token:",  bearer)
	json.Unmarshal(bodyBytes, &bearer)

	fmt.Println("API access token:",  bearer.Bearer)
	return bearer.Bearer
}