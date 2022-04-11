package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/juancsr/protocol-buffers-udemy/go-with-protoc/src/pb/simplepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Recover from error:", err)
		}
	}()

	sm := doSimple()
	messageFile := "simple_message.bin"
	if err := writeToFile(messageFile, sm); err != nil {
		log.Fatalf("could not write to file: %v\n", err)
	}
	fmt.Printf("Data has been saved into the file: %v\n", messageFile)

	// sm2 := &simplepb.SimpleMessage{}
	sm2 := new(simplepb.SimpleMessage)

	if err := readFromFile(messageFile, sm2); err != nil {
		log.Fatalf("could not read message from file: %v\n", err)
	}
	fmt.Println("Message: ", sm2)

	sm2JSON := toJSON(sm2)
	fmt.Println(sm2JSON)

	var sm3 simplepb.SimpleMessage
	fromJSON(sm2JSON, &sm3)
	fmt.Println("Message from JSON: ", sm3)
}

func fromJSON(json string, pb protoiface.MessageV1) {
	err := jsonpb.UnmarshalString(json, pb)
	if err != nil {
		log.Panicln("could not unmarshal json string to message:", err)
	}
}

func toJSON(pb protoiface.MessageV1) string {
	marshaler := jsonpb.Marshaler{}
	msg, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("could not marshal message to json")
	}
	return msg
}

func doSimple() *simplepb.SimpleMessage {
	fmt.Println("Simple message")
	sm := simplepb.SimpleMessage{
		Id:         1,
		IsSimple:   true,
		Name:       "This is a simple message",
		SampleList: []int32{1, 2, 3},
	}

	sm.Name = "I renamed you"

	fmt.Printf("The id is: %v\n", sm.GetId())
	return &sm
}

func writeToFile(path string, pb proto.Message) error {
	bytesIn, err := proto.Marshal(pb)
	if err != nil {
		return fmt.Errorf("cannot serialize to path: %v", err)
	}

	return ioutil.WriteFile(path, bytesIn, 0644)
}

func readFromFile(path string, pm proto.Message) error {
	bytesOut, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read from file: %v", err)
	}

	return proto.Unmarshal(bytesOut, pm)
}
