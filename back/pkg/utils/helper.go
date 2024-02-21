package utils

import (
	"log"
	"runtime"
)

func PrintFunctionName() {
	pc, _, _, ok := runtime.Caller(1) // Skip one frame to get the caller's name
	if !ok {
		log.Println("Failed to get function name")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()

	log.Println("**FUNCTION**:", funcName)
}

func GetContestName() string {
	// Make the HTTP request

	return "The " + apiWordAPI("adjective") + " " + apiWordAPI("noun")
}
