# Golang Promaming with Protocol Buffers

## Setup & Code Download in Golang

- Golang installed: https://golang.org/doc/install

- VSCode Installed

- The VSCode Golang extension:https://code.visualstudio.com/docs/languages/go

- The source code for this project (download / star the project): https://github.com/simplesteph/protobuf-example-go

## Code generation in Golang

More information:
https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

Command
```bash
protoc -I src/ --go_out=src/ src/simple
```

Proto file
```proto
syntax = "proto3";

package example.simple;

option go_package = "./simple";

message SimpleMessage {
  int32 id = 1;
  bool is_simple = 2;
  string name = 3;
  repeated int32 sample_list = 4;
}
```

## Simple Proto Struct in Golang

It is recommended to use the handler methods `Get` to avoid errors when message is nill.

```go
fmt.Println("Simple message")
sm := simplepb.SimpleMessage{
    Id:         1,
    IsSimple:   true,
    Name:       "This is a simple message",
    SampleList: []int32{1, 2, 3},
}

sm.Name = "I renamed you"

fmt.Printf("The id is: %v\n", sm.GetId())
```

## go_package option

To specify the packages directory we can use `go_package` option in the proto file.

```prot
...
option go_package = "pb/simplepb";
...
```

## Reading and Writing to Disk

```go

func main() {
	// sm is a Simple Message
	messageFile := "simple_message.bin"
	if err := writeToFile(messageFile, sm); err != nil {
		log.Fatalf("could not write to file: %v\n", err)
	}
	fmt.Printf("Data has been saved into the file: %v\n", messageFile)

	sm2 := &simplepb.SimpleMessage{}

	if err := readFromFile(messageFile, sm2); err != nil {
		log.Fatalf("could not read message from file: %v\n", err)
	}
	fmt.Println("Message: ", sm2)
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
```

## Reading and Writing from JSON

**Mostly used for debug pourposes, not for PROD**

```go
func main() {
    ...
    sm2JSON := toJSON(sm2)
	fmt.Println(sm2JSON)

	var sm3 simplepb.SimpleMessage
	fromJSON(sm2JSON, &sm3)
	fmt.Println("Message from JSON: ", sm3)
    ...
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
```

## Enum Proto Struct in Golang

```go
func main() {
	// ...
	// enum
	fmt.Println("doEnum:", doEnum())
	// ...
}

func doEnum() *simplepb.Enumeration {
	return &simplepb.Enumeration{
		EyeColor: simplepb.EyeColor_EYE_COLOR_BROWN,
	}
}

```

## Complex Proto Struct in Golang

```go
func main() {
	// ...
	// complex
	fmt.Println("doComplex:", doComplex())
	// ...
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
```

## Handling Maps in Golang

```go
func() {
	// ...
	// doMap
	fmt.Println("doMap:", doMap())
	// ...
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
```

## Handling OneOfs in Golang
More info here:
https://developers.google.com/protocol-buffers/docs/proto3#oneof




```go
func main() {
	// ...
	// oneof
	doOneOf(&oneofspb.Result_Id{Id: 1})
	doOneOf(&oneofspb.Result_Message{Message: "juancsr"})
	// ...
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
```


## Practice Exercise Goalang

Code is [here](./my-code/go-with-protoc/exercise.go)