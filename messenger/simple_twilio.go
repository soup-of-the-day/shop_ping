package messenger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Holds twilio auth information
type twilioAuth struct{
	AccountSid string
	AuthToken string
	ToNumber string
	FromNumber string
}

func createTwilioAuth() *twilioAuth {
	return &twilioAuth{
		AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
		ToNumber: os.Getenv("TWILIO_TO_NUMBER"),
		FromNumber: os.Getenv("TWILIO_FROM_NUMBER"),
	}
}

// Uses Twilio API to send a text to my phone
func SendAlert() {
	auth := createTwilioAuth()
	apiUrl := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", auth.AccountSid)

	msgData := url.Values{}
	msgData.Set("To",auth.ToNumber)
	msgData.Set("From",auth.FromNumber)
	msgData.Set("Body","Yo, your item is suddenly available!")
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, &msgDataReader)
	req.SetBasicAuth(auth.AccountSid, auth.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if (err != nil) {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(resp.Status)
		responseData, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(responseData))
	}
}
