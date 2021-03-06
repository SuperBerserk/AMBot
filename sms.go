package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func sendSMS(user string, msg string) {
	accountSid := "AC6861dbd95a4542d1be1c0a219d7cecec"
	authToken := ""
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// build our data for our message
	v := url.Values{}
	fmt.Printf("Lookup %s\n", directory[user])
	v.Set("To", directory[user])

	v.Set("From", "+17873392841")
	v.Set("Body", msg+"\n- AMBot")

	rb := *strings.NewReader(v.Encode())

	// create and set parameters for the http client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make the request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

}
