package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/juancsr/protocol-buffers-udemy/go-with-protoc/src/pb/complexpb"
	"github.com/juancsr/protocol-buffers-udemy/go-with-protoc/src/pb/mapspb"
	"github.com/juancsr/protocol-buffers-udemy/go-with-protoc/src/pb/oneofspb"
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

	// course()
	exercise()
}

func course() {
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

	// complex
	fmt.Println("doComplex:", doComplex())

	// enum
	fmt.Println("doEnum:", doEnum())

	// doMap
	fmt.Println("doMap:", doMap())

	// oneof
	doOneOf(&oneofspb.Result_Id{Id: 1})
	doOneOf(&oneofspb.Result_Message{Message: "juancsr"})
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *oneofspb.Result_Id:
		fmt.Println(message.(*oneofspb.Result_Id).Id)
	case *oneofspb.Result_Message:
		fmt.Println(message.(*oneofspb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpeced type: %v", x)
	}
}

func doMap() *mapspb.MaxExample {
	return &mapspb.MaxExample{
		Ids: map[string]*mapspb.IdWrapper{
			"my_id_43": {Id: 43},
			"my_id_44": {Id: 44},
			"my_id_45": {Id: 45},
		},
	}
}

func doComplex() *complexpb.Complex {
	return &complexpb.Complex{
		OneDummy: &complexpb.Dummy{Id: 32, Name: "juancsr"},
		MultiDummy: []*complexpb.Dummy{
			{Id: 1, Name: "juancsr_1"},
			{Id: 2, Name: "juancsr_2"},
			{Id: 3, Name: "juancsr_3"},
		},
	}
}

func doEnum() *simplepb.Enumeration {
	return &simplepb.Enumeration{
		EyeColor: simplepb.EyeColor_EYE_COLOR_BROWN,
	}
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
