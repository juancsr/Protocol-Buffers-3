package com.github.juancsr.protobuf;

import example.enumerations.EnumExample;
import example.enumerations.EnumExample.DayOfTheWeek;
import example.enumerations.EnumExample.EnumMessage;

public class EnumMain {
    public static void main(String[] args) {
        System.out.println("Example for enum");

        EnumMessage.Builder builder = EnumMessage.newBuilder();

        builder.setId(1)
                .setDayOfTheWeek(DayOfTheWeek.SATURDAY);

        EnumMessage message = builder.build();

        System.out.println(message);
    }
}
