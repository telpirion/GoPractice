package prodneutral

import (
	"log"

	"google.golang.org/protobuf/types/known/structpb"
)

func creatingProtobufValue() error {

	m, err := structpb.NewValue(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Smith",
		"isAlive":   true,
		"age":       27,
		"address": map[string]interface{}{
			"streetAddress": "21 2nd Street",
			"city":          "New York",
			"state":         "NY",
			"postalCode":    "10021-3100",
		},
		"phoneNumbers": []interface{}{
			map[string]interface{}{
				"type":   "home",
				"number": "212 555-1234",
			},
			map[string]interface{}{
				"type":   "office",
				"number": "646 555-4567",
			},
		},
		"children": []interface{}{},
		"spouse":   nil,
	})
	if err != nil {
		return err
	}

	log.Println(m)

	return nil
}
