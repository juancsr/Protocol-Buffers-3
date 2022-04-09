# Java Programming with Protocol Buffers 

## Setup & Code Download in Java

- Java 8 JDK:
http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

- IntelliJ IDEA Community Edition:
https://www.jetbrains.com/idea/download/

- The source code for this project (download / star the project):
https://github.com/simplesteph/protobuf-example-java

# Gradle project setup & code generation

## Protobuf Gradle Plugin
https://github.com/google/protobuf-gradle-plugin

```
apply plugin: 'com.google.protobuf'
```


```
buildscript {
  repositories {
    mavenCentral()
  }
  dependencies {
    classpath 'com.google.protobuf:protobuf-gradle-plugin:0.8.18'
  }
}
```

protobuf compiler maven versions:
https://mvnrepository.com/artifact/com.google.protobuf/protoc

```
protobuf {
    // Configure the protoc executable
    protoc {
        // Download from repositories
        artifact = 'com.google.protobuf:protoc:3.20.0'
    }
}
```

[Tips for IntelliJ](https://github.com/google/protobuf-gradle-plugin#tips-for-ides)

Protobuf core:
https://mvnrepository.com/artifact/com.google.protobuf/protobuf-java

```
dependencies {
    ...
    compile 'com.google.protobuf:protobuf-java:3.20.0'
    ...
}
```

## Simple Message Creation in Java
```java
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
```

## Enums in Java

```java
public static void main(String[] args) {
        System.out.println("Example for enum");

        EnumMessage.Builder builder = EnumMessage.newBuilder();

        builder.setId(1)
                .setDayOfTheWeek(DayOfTheWeek.SATURDAY);

        EnumMessage message = builder.build();

        System.out.println(message);
    }
```

## Complex message

```java
public static void main(String[] args) {

        System.out.println("Complex example");

        DummyMessage dummyMessage = newDummyMessage(55, "New dump message");

        System.out.println(dummyMessage.toString());

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

        System.out.println(message.toString());

        // Get
        message.getMultipleDummyList();
    }

    public static DummyMessage newDummyMessage(int id, String name) {
        DummyMessage.Builder builder = DummyMessage.newBuilder();
        DummyMessage dummymessage = builder.setName(name)
                .setId(id)
                .build();
        return dummymessage;
    }
```

## Java Options

More info here: https://developers.google.com/protocol-buffers/docs/proto3#options

```protobuf
syntax = "proto3";

package example.options;

// Setting up java packages when generated
option java_package = "com.example.options";

// using one file for each message
option java_multiple_files = true;

message OptionMessage {
  int32 id = 1;
  InnerOption inner_option = 2;
}

message InnerOption {
  string value = 1;
}
```

## Conversion to JSON in Java

NOT RECOMENDED FOR PRODUCTION!

```gradle
dependecies {
  ...
  implementation 'com.google.protobuf:protobuf-java-util:3.20.0'
  ...
}
```

```java
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
```

## Practice Exercise Java

"*Implement a way to read and write Person in the Address Book.*"

[AddressBookMain.java](./my-code/java-with-protoc/src/main/java/com/github/juancsr/protobuf/AddressBookMain.java)