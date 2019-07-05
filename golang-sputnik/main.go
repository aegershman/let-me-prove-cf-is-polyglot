package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	beepRate := 1000
	value, exists := os.LookupEnv("BEEP_RATE")
	if exists {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Unable to parse %s as integer\n", value)
		}
		beepRate = v
	}

	log.SetOutput(os.Stdout)
	customMessage, exists := os.LookupEnv("CUSTOM_MESSAGE")
	for {
		if exists {
			log.Println(customMessage, "beep")
		} else {
			log.Println("beep")
		}
		time.Sleep(time.Duration(beepRate) * time.Millisecond)
	}
}
