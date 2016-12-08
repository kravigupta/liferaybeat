package main

import (
	"os"
    "fmt"
	"github.com/elastic/beats/libbeat/beat"

	"github.com/packt/liferaybeat/beater"
)

func main() {
    fmt.Println("Starting Liferay beat")
	err := beat.Run("liferaybeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
