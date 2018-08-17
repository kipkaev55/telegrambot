# Simple Golang telegram bot
### How to use
```js
package main

import (
	"log"
	"os"
	"time"
	telegrambot "github.com/kipkaev55/telegrambot"
)

func main() {
	bot := telegrambot.NewBot(os.Getenv("TELEGRAM_TOKEN"), "Prefix: ")
	// or
	// var sendingPeriod float64 // frequency of sending messages, default 3 second 
	// sendingPeriod = 5.0 // in seconds
	// bot := telegrambot.NewBot(os.Getenv("TELEGRAM_TOKEN"), "Prefix: ", sendingPeriod)
	_, err := bot.SendMessage(os.Getenv("CHAT_ID_NOTIFY"), "Start application successfully")
	// or send with marker
	// _, err := bot.SendMessage(os.Getenv("CHAT_ID_NOTIFY"), "marker1", "Start application successfully")
	if err != nil {
		log.Printf("%s: %s\n", time.Now().Format(os.Getenv("LOG_TIME_FORMAT")), err.Error())
	}
}
```
