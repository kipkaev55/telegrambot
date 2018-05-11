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
	bot := telegrambot.NewBot(os.Getenv("TELEGRAM_TOKEN"), "DBMS: ")
	_, err := bot.SendMessage(os.Getenv("CHAT_ID_NOTIFY"), "Start application successfully")
	if err != nil {
		log.Printf("%s: %s\n", time.Now().Format(os.Getenv("LOG_TIME_FORMAT")), err.Error())
	}
}
```
