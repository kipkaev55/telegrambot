package telegram

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// APISendMessageURL api url to send telegram message
	APISendMessageURL = "https://api.telegram.org/bot%s/sendMessage"
	// bodySendMessageFormat is format of send message
	bodySendMessageFormat = `{"chat_id":"%s","text":"%s","disable_notification":"%b"}`
)

type Bot struct {
	Token  string
	Prefix string
}

func NewBot(token, prefix string) *Bot {
	return &Bot{Token: token, Prefix: prefix}
}

// SendMessage sending message to client or chat
func (b Bot) SendMessage(chatID, text string) (bool, error) {
	var err error
	var success bool
	defer func() {
		if x := recover(); x != nil {
			log.Printf("%s run time panic: %v\n", time.Now().Format("2006-01-02 15:04:05"), x)
			success = false
			err = errors.New(fmt.Sprintf("%v", x))
		}
	}()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	URL := fmt.Sprintf(APISendMessageURL, b.Token)
	body := fmt.Sprintf(bodySendMessageFormat, chatID, b.Prefix+text, 0)
	var jsonStr = []byte(body)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(time.Now().Format("2006-01-02 15:04:05"), err.Error())
		log.Println(time.Now().Format("2006-01-02 15:04:05"), "response Status:", resp.Status)
		responseBody, _ := ioutil.ReadAll(resp.Body)
		log.Println(time.Now().Format("2006-01-02 15:04:05"), "response Body:", string(responseBody))
		success = false
	} else {
		success = true
	}
	defer resp.Body.Close()
	return success, err
}
