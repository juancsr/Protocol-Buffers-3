package com.github.juancsr.protobuf;

import example.simple.Simple.SimpleMessage;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class SimpleMain {
    public static void main(String[] args) {
        System.out.println("Hello world");

        SimpleMessage.Builder builder = SimpleMessage.newBuilder();
        builder.setId(1)
            .setIsSimple(true)
            .setName("My name is simple");

        // repeated field
        builder.addSampleList(1);
        builder.addSampleList(2);
        builder.addSampleList(3);

        builder.addAllSampleList(Arrays.asList(4,5,6));

        // remove all data
        //builder.clearSampleList();

        System.out.println(builder.toString());

        // message
        SimpleMessage message = builder.build();


        // write the protocol buffer binary to a file
        try {
            FileOutputStream outputStream = new FileOutputStream("simple_message.bin");
            message.writeTo(outputStream);
            outputStream.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        // send as byte array (network)
        //byte[] bytes = message.toByteArray();

        try {
            System.out.println("Reading from file");
            FileInputStream inputStream = new FileInputStream("simple_message.bin");
            SimpleMessage messageFromFile = SimpleMessage.parseFrom(inputStream);
            System.out.println(messageFromFile);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

    }
}
