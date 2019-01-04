package main

import (
	"net/http"
	"io"
	"time"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"io/ioutil"
)

//var apiKey string = "b6907d289e10d714a6e88b30761fae22"

// Fetch makes network calls using the method (POST/GET..), the URL // to hit, headers to add (if any), and the body of the request.

func fetchData(method string, url string, header map[string]string, body io.Reader) (*http.Response, error) {

	//Dialer is a struct that manages the establishment of the connection. A Dialer contains options for connecting to an address
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
		//Deadline: time.Now(),                                           //takes deadline in terms of actual time
	}

	//A Transport is a struct used by clients to manage the underlying TCP connection
	transport := &http.Transport{
		Dial: dialer.Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		DisableKeepAlives: true, //if true, prevents re-use of TCP connections
	}

	// Create client with required custom parameters.
	client := &http.Client{
		Transport: transport,
		Timeout: time.Duration(10 * time.Second), //Timeout specifies a time limit for requests made by this Client
	}

	// Create request.
	request, err := http.NewRequest(method, url, body)
	if (err != nil) {
		fmt.Println("Erron in making new request :", err.Error())
	}

	// Add any required headers.
	for key, val := range header {
		request.Header.Add(key, val);
	}

	// Perform said network call.
	respone, err2 := client.Do(request)
	if (err2 != nil) {
		fmt.Println("Error :", err2.Error())
		return nil, err2
	}

	// If response from network call is not 200, return error too.
	if (respone.StatusCode != http.StatusOK) {
		return respone, errors.New("Http call failed !!")
	}

	return respone, nil;

}

func main() {
	const sampleUrl string  = "https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22"
	var headers = make(map[string]string)
        var reqBody io.Reader
	res , err := fetchData("GET" , sampleUrl , headers , reqBody)
	if (err != nil){
		fmt.Println("ERROR IN ::" , err.Error())
		return
	}

	respBody , _ := ioutil.ReadAll(res.Body)

	fmt.Println("RESPONSe ::" , string(respBody))

}


