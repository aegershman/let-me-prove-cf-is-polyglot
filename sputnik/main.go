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
	for {
		log.Println("beep")
		time.Sleep(time.Duration(beepRate) * time.Millisecond)
	}
}
