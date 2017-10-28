package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"charisma-beat/integration/events"
)

const EVENTS_CONFIG_LOCATION = "./config"

func init() {
	readDynamicEventDefinitions()
}


func readDynamicEventDefinitions() {
	eventsConfiguration = make([]events.Config, 0)
	files, err := ioutil.ReadDir(EVENTS_CONFIG_LOCATION)
	if err != nil {
		log.Fatalln("reading directory config error", err)
	}
	for _, file := range files {
		fileContent, err := ioutil.ReadFile(EVENTS_CONFIG_LOCATION + "/" + file.Name())
		if err != nil {
			log.Println("ReadFile error:", err)
			continue
		}

		config := events.Config{}
		err = json.Unmarshal(fileContent, &config)
		if err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		eventsConfiguration = append(eventsConfiguration, config)
	}
}
