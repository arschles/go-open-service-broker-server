package osbserver

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "osbserver", log.Lshortfile|log.Ldate|log.Ltime)
