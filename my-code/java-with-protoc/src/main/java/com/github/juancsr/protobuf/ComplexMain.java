package com.github.juancsr.protobuf;

import example.complex.Complex.DummyMessage;
import example.complex.Complex.ComplexMessage;

import java.util.Arrays;

public class ComplexMain {
    public static void main(String[] args) {

        System.out.println("Complex example");

        DummyMessage dummyMessage = newDummyMessage(55, "New dump message");

        System.out.println(dummyMessage);

        ComplexMessage.Builder builder = ComplexMessage.newBuilder();

        // singular message
        builder.setOneDummy(dummyMessage);

        // repeated message
        builder.addMultipleDummy(newDummyMessage(66, "New message 66"));
        builder.addMultipleDummy(newDummyMessage(67, "New message 67"));
        builder.addMultipleDummy(newDummyMessage(68, "New message 68"));

        builder.addAllMultipleDummy(Arrays.asList(
            newDummyMessage(20, "Hi"), newDummyMessage(21, "There")
        ));

        ComplexMessage message = builder.build();

        System.out.println(message);

        // Get
        // message.getMultipleDummyList();
    }

    public static DummyMessage newDummyMessage(int id, String name) {
        DummyMessage.Builder builder = DummyMessage.newBuilder();
        return builder.setName(name)
                .setId(id)
                .build();
    }
}
