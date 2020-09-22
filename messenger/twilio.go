package messenger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Holds twilio auth information
type TwilioMessenger struct{
	AccountSid string
	AuthToken string
	ToNumber string
	FromNumber string
}

// Expecting auth information to exist as environment variables
func CreateTwilioMessenger() *TwilioMessenger {
	return &TwilioMessenger{
		AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
		ToNumber: os.Getenv("TWILIO_TO_NUMBER"),
		FromNumber: os.Getenv("TWILIO_FROM_NUMBER"),
	}
}

// Uses Twilio API to send a text to my phone
func (m *TwilioMessenger) SendAlert(msg string) error {
	apiUrl := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", m.AccountSid)

	msgData := url.Values{}
	msgData.Set("To",m.ToNumber)
	msgData.Set("From",m.FromNumber)
	msgData.Set("Body",msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, &msgDataReader)
	req.SetBasicAuth(m.AccountSid, m.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return err
		}
	} else {
		log.Println(resp.Status)
		responseData, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf(string(responseData))
	}
	log.Println("Message sent!")
	return nil
}
