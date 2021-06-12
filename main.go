package main

import (
	"fmt"
	"log"
	"math"
	"os"
	
	"github.com/BurntSushi/toml"
	"github.com/distatus/battery"
	"github.com/gregdel/pushover"
)

type tomlConfig struct {
	Id string
	ApiKey string
	RecipientKey string
}

func main() {
	var config tomlConfig

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	mybattery := 0.0
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, battery := range batteries {
		fmt.Printf("Bat%d: ", i)
		mybattery = battery.Current / battery.Full * 100
	}
	app := pushover.New(config.ApiKey)

	recipient := pushover.NewRecipient(config.RecipientKey)
	if mybattery < 30 {
		message := pushover.NewMessage(fmt.Sprintf("%s Charge Low!: %d", config.Id, int64(math.Floor(mybattery))))

		response, err := app.SendMessage(message, recipient)
		if err != nil {
			log.Panic(err)
		}

		log.Println(response)
	}

}
