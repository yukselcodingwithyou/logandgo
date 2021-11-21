package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func toJson(lf LogFields) string {
	data, err := json.Marshal(lf)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func toText(_ string, fields LogFields) string {
	var builder string
	for mKey, mVal := range fields {
		switch typedVal := mVal.(type) {
		case string:
			builder += fmt.Sprintf("%s=%s ", mKey, typedVal)
		case float64:
			builder += fmt.Sprintf("%s.%-1.0f ", mKey, typedVal)
		case map[string]interface{}:
			builder += toText(mKey, typedVal)
		}
	}

	return builder
}

func output(logger *log.Logger, logLevel LogLevel, level LogLevel, log string) {
	if logLevel <= level {
		err := logger.Output(3, fmt.Sprintln(log))
		if err != nil {
			fmt.Println("log action fails")
		}
	}
}
