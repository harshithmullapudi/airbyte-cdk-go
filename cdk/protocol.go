package main

// AirbyteLogLevel describes the types of logs you can perform in connector
var AirbyteLogLevel = map[string]string{
	"FATAL": "FATAL",
	"ERROR": "ERROR",
	"WARN":  "WARN",
	"INFO":  "INFO",
	"DEBUG": "DEBUG",
	"TRACE": "TRACE",
}

// AirbyteConnectionStatus describes the status of connection
var AirbyteConnectionStatus = map[string]string{
	"SUCCEEDED": "SUCCEEDED",
	"FAILED":    "FAILED",
}

// AirbyteConfig is a generic configuration object
var AirbyteConfig map[string]interface{}

type AirbyteMessage struct {
	messageType string
}
