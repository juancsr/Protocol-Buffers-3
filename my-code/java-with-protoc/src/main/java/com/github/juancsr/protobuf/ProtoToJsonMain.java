package com.github.juancsr.protobuf;

import com.google.protobuf.InvalidProtocolBufferException;
import example.simple.Simple;
import com.google.protobuf.util.JsonFormat;
import java.util.Arrays;

public class ProtoToJsonMain {

    public static void main(String[] args) {
        System.out.println("Proto to JSON Example");

        Simple.SimpleMessage.Builder builder = Simple.SimpleMessage.newBuilder();
        builder.setId(1)
                .setIsSimple(true)
                .setName("My name is simple");

        // repeated field
        builder.addSampleList(1);
        builder.addSampleList(2);
        builder.addSampleList(3);

        builder.addAllSampleList(Arrays.asList(4,5,6));

        // Print this as a JSON
        String jsonString;
        try {
            jsonString = JsonFormat.printer().print(builder);

            System.out.println(jsonString);

            Simple.SimpleMessage.Builder builder2 = Simple.SimpleMessage.newBuilder();

            JsonFormat.parser()
                    .ignoringUnknownFields()
                    .merge(jsonString, builder2);

            System.out.println(builder2);
        } catch (InvalidProtocolBufferException e) {
            e.printStackTrace();
        }

    }
}
