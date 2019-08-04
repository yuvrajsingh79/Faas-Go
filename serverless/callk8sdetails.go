package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	api_key = "YOUR_API_KEY" //I have used IBM Bluemix trigger api key
)

func main() {
	// Use custom certificate for testing.  Not exactly required by task.
	b, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM(b); !ok {
		log.Fatal("Failed to append cert")
	}
	tc := &tls.Config{RootCAs: pool}
	tr := &http.Transport{TLSClientConfig: tc}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", "https://eu-gb.functions.cloud.ibm.com/api/v1/namespaces/singhyuvraj79%40gmail.com_dev/triggers/trigger1", nil)
	if err != nil {
		log.Fatal(err)
	}

	// This one line implements the authentication required for the task.
	req.SetBasicAuth(api_key)

	// Make request and show output.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
