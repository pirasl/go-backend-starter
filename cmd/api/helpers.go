package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func readJSON(w http.ResponseWriter, r *http.Request) {

}

func writeJSON(w http.ResponseWriter, r *http.Request, status int, message string) {}

func getEnvStrToI(envName string, fallback int, throwError ...bool) int {
	mandatory := false
	if len(throwError) > 0 {
		mandatory = throwError[0]
	}

	if env := os.Getenv(envName); env != "" {
		parsedValue, err := strconv.Atoi(env)
		if err != nil {
			log.Fatalf("Error while parsing env var: %v", envName)
		}
		return parsedValue
	}
	if mandatory {
		log.Fatalf("the following env var is missing: %v.\nIt should be defined  in the .env file at root level. See .env.example file for more info", envName)
	}
	return fallback
}

func getEnvStr(envName string, fallback string, throwError ...bool) string {
	mandatory := false
	if len(throwError) > 0 {
		mandatory = throwError[0]
	}

	if env := os.Getenv(envName); env != "" {
		return env
	}
	if mandatory {
		log.Fatalf("the following env var is missing: %v\n. It should be defined  in the .env file at root level\n. See .env.example file for more info", envName)
	}
	return fallback
}

func getEnvDuration(env string, fallback time.Duration, throwError ...bool) time.Duration {
	mandatory := false
	if len(throwError) > 0 {
		mandatory = throwError[0]
	}

	if env := os.Getenv(env); env != "" {
		duration, err := time.ParseDuration(env)
		if err != nil {
			fmt.Printf("Error while parsing env: %v. defaulting to fallback: %v", env, fallback)
		}
		return duration
	}
	if mandatory {
		log.Fatalf("the following env var is missing: %v\n. It should be defined in the .env file at root level\n. See .env.example file for more info", env)
	}
	return fallback
}

func getEnvBool(env string, fallback bool, throwError ...bool) bool {
	mandatory := false
	if len(throwError) > 0 {
		mandatory = throwError[0]
	}

	if env := os.Getenv(env); env != "" {
		boolean, err := strconv.ParseBool(env)
		if err != nil {
			fmt.Printf("Error while parsing env: %v. defaulting to fallback: %v", env, fallback)
		}
		return boolean
	}

	if mandatory {
		log.Fatalf("the following env var is missing: %v\n. It should be defined in the .env file at root level\n. See .env.example file for more info", env)
	}
	return fallback
}
