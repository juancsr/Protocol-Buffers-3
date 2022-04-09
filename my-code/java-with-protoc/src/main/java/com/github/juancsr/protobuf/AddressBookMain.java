package com.github.juancsr.protobuf;

import com.example.tutorial.protos.Person;
import com.example.tutorial.protos.Person.PhoneNumber;
import com.example.tutorial.protos.Person.PhoneType;
import com.example.tutorial.protos.AddressBook;
import com.google.protobuf.Timestamp;
import example.simple.Simple;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.sql.Time;
import java.time.Instant;
import java.util.Arrays;

public class AddressBookMain {
    final static private String messagePath = "address_book";

    public static void main(String[] args) {
        // "Implement a way to read and write Person in the Address Book. "

        System.out.println("Address Book Example");

        // juancsr_1
        Iterable<PhoneNumber> phones = Arrays.asList(
                newPhoneNumberMessage("123", PhoneType.MOBILE),
                newPhoneNumberMessage("45", PhoneType.HOME),
                newPhoneNumberMessage("123-345", PhoneType.WORK)
        );
        Person person = newPersonMessage("juancsr_1", 1, "juancsr_1@github.com", phones,
                newTimeStampMessage());

        // juancsr_2
        Iterable<PhoneNumber> phones2 = Arrays.asList(
                newPhoneNumberMessage("654-321", PhoneType.MOBILE)
        );
        Person person2 = newPersonMessage("juancsr_2", 2, "juancsr_2@github.com", phones2,
                newTimeStampMessage());

        Iterable<Person> people = Arrays.asList(
                person, person2
        );

        AddressBook addressBook = newAddressBookMessage(people);

        // Write messagee
        try {
            System.out.println("Sending message to file");
            FileOutputStream outputStream = new FileOutputStream(messagePath);
            addressBook.writeTo(outputStream);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        // Read message
        try {
            AddressBook addressBookMessage = getAddressBookMessage();
            System.out.println(addressBookMessage);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }


    public static AddressBook getAddressBookMessage() throws IOException {

        System.out.println("Reading from file");
        FileInputStream inputStream = new FileInputStream(messagePath);
        return AddressBook.parseFrom(inputStream);
    }

    public static Timestamp newTimeStampMessage() {
        Instant now = Instant.now();
        Timestamp.Builder builder = Timestamp.newBuilder();
        return builder
                .setSeconds(now.getEpochSecond())
                .setNanos(now.getNano())
                .build();
    }

    public static Person newPersonMessage(String name, int id, String email, Iterable<PhoneNumber> phones,
                                   Timestamp lastUpdated) {
        Person.Builder builder = Person.newBuilder();
        return builder.setEmail(email)
                .setId(id)
                .setName(name)
                .addAllPhones(phones)
                .setLastUpdated(lastUpdated)
                .build();
    }

    public static PhoneNumber newPhoneNumberMessage(String number, PhoneType phoneType) {
        PhoneNumber.Builder builder = PhoneNumber.newBuilder();
        return builder.setNumber(number)
                .setType(phoneType)
                .build();
    }

    public static AddressBook newAddressBookMessage(Iterable<Person> people) {
        AddressBook.Builder builder = AddressBook.newBuilder();
        return builder.addAllPeople(people).build();
    }
}
