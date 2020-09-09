package main

import (
	"fmt"
	"github.com/mkideal/log"
	"github.com/mkideal/log/provider"
)

func main() {
	opts2 := fmt.Sprintf(`{
		"dir": "./log",
		"filename": "app.log"
	}`)

	defer log.Uninit(log.InitWithProvider(provider.NewFile(opts2)))
	log.Trace("%s cannot be printed", "TRACE")
	log.Debug("%s cannot be printed", "DEBUG")

	log.Info("%s should be printed", "INFO")
	log.Warn("%s should be printed", "WARN")
	log.Error("%s should be printed", "ERROR")

	log.If(true).Info("%v should be printed", true)

	iq := 250
	log.If(iq < 250).Info("IQ less than 250").
		ElseIf(iq > 250).Info("IQ greater than 250").
		Else().Info("IQ equal to 250")

	log.With("hello").Info("With a string field")
	log.With("hello").Info("")
	log.With(1).Info("With an int field")
	log.With(true).Info("With a bool field")
	log.With(1, "2", false).Info("With 3 fields")
	log.With(log.M{"a": 1}).Info("With a map")
	log.WithJSON(log.M{"a": 1}).Info("With a map and using JSONFormatter")

	// don't print message header
	log.NoHeader()

	log.Info("This message have no header")

	log.Fatal("%s should be printed and exit program with status code 1", "FATAL")

	log.Info("You cannot see me")
}
