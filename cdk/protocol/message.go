package protocol

import (
	"encoding/json"
	"errors"
)

const (
	RECORD            = "RECORD"
	CATALOG           = "CATALOG"
	LOG               = "LOG"
	CONNECTION_STATUS = "CONNECTION_STATUS"
	SPEC              = "SPEC"
	STATE             = "STATE"
)

type AirbyteMessage struct {
	messageType string
}

type AirbyteRecordMessage struct {
	messageType string
	stream      string
	namespace   string
	emittedAt   int `json:emitted_at`
	data        map[string]interface{}
}

// ParseMessage converts string into a recorc
func ParseMessage(s string) (AirbyteRecordMessage, error) {
	var err error
	message := AirbyteMessage{}
	err = json.Unmarshal([]byte(s), &message)

	// Unable to parse the string into json
	if err != nil {
		return AirbyteRecordMessage{}, errors.New("Cannot parse the message")
	}

	switch messageType := message.messageType; messageType {
	case RECORD:
		recordMessage := AirbyteRecordMessage{}
		err = json.Unmarshal([]byte(s), &recordMessage)

		// Unable to parse the string into json
		if err != nil {
			return AirbyteRecordMessage{}, errors.New("Cannot parse the message into record format")
		}

		return recordMessage, nil
	case CATALOG:
	case LOG:
	case STATE:
	case CONNECTION_STATUS:
	case SPEC:
		return AirbyteRecordMessage{
			messageType: message.messageType,
		}, nil
	default:
		return AirbyteRecordMessage{}, errors.New("Invalid Airbyte message")
	}

	return AirbyteRecordMessage{
		messageType: message.messageType,
	}, nil
}

// Record converts the details into record message
func AirbyteRecord(message AirbyteMessage, stream string, namespace string, emittedAt int, data map[string]interface{}) AirbyteRecordMessage {
	return AirbyteRecordMessage{
		messageType: message.messageType,
		stream:      stream,
		namespace:   namespace,
		emittedAt:   emittedAt,
		data:        data,
	}
}
