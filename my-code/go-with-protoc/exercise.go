// Using this .proto definition:
// https://github.com/google/protobuf/blob/master/examples/addressbook.proto

// Implement a way to read and write Person in the Address Book.

// Use all the learnings from the previous lectures!

// When you are ready, the solution will be here:
// https://developers.google.com/protocol-buffers/docs/gotutorial

// And it will be a good time to read the documentation :)

package main

import (
	"fmt"
	"log"

	"github.com/juancsr/protocol-buffers-udemy/go-with-protoc/src/pb/addressbookpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func exercise() {
	emilia := newPerson("Emilia", "emilia@protobuf.com", 1, []*addressbookpb.Person_PhoneNumber{
		newPhoneNumber("123", addressbookpb.PhoneType_HOME),
		newPhoneNumber("456", addressbookpb.PhoneType_WORK),
		newPhoneNumber("789", addressbookpb.PhoneType_MOBILE),
	})

	juan := newPerson("juan", "juancsr@protobuf.com", 1, []*addressbookpb.Person_PhoneNumber{})

	people := []*addressbookpb.Person{emilia, juan}
	addressBook := newAddressBook(people)

	messageFile := "address_book"

	fmt.Printf("Writing message to %s file\n", messageFile)
	err := writeToFile(messageFile, addressBook)
	if err != nil {
		log.Panicf("Could not write to %s file: %v", messageFile, err)
	}

	var msg addressbookpb.AddressBook

	fmt.Printf("Reading messages from %s file\n", messageFile)
	err = readFromFile(messageFile, &msg)
	if err != nil {
		log.Panicf("Could not read from %s file: %v", messageFile, err)
	}

	fmt.Println("Message info")
	fmt.Println("People: ", msg.People)

	jsonMsg := toJSON(&msg)
	fmt.Println("msg: ", jsonMsg)
}

func newAddressBook(people []*addressbookpb.Person) *addressbookpb.AddressBook {
	return &addressbookpb.AddressBook{
		People: people,
	}
}

func newPerson(name, email string, id int32, phones []*addressbookpb.Person_PhoneNumber) *addressbookpb.Person {
	return &addressbookpb.Person{
		Name:        name,
		Email:       email,
		Id:          id,
		LastUpdated: timestamppb.Now(),
		Phones:      phones,
	}
}

func newPhoneNumber(number string, phoneType addressbookpb.PhoneType) *addressbookpb.Person_PhoneNumber {
	return &addressbookpb.Person_PhoneNumber{
		Number: number,
		Type:   phoneType,
	}
}
