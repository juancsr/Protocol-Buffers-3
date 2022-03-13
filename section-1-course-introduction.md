# Protocol Bufferws

# Course Introduction

## The need for Protocol Buffers

### Comma Separated Values (CSV)
* Advantages:
    * Easy to parse
    * Easy to read
    * Easy to make sense of

* Disadvantages:
    * The data types of elements has to be inferred and is not a guarantee
    * Parsing becomes tricky when data contains comma

### Relational tables definitions
* Advantages:
    * Data is fully typed
    * Data fits in a table

* Disadvantages:
    * Data has to be flat
    * DAta is stored in database, and data definition will be different for each database

### JavaScript Object Notation (JSON)
* Advantanges
    * Take any form
    * Widely accepted format on the web
    * read by pretty much any language
    * easily shared over a network

* Disadvantages:
    * No schema enforcing
    * JSON Objects can be quite big in size because of repeated keys

### Protocol Buffers
* Protocol Buffers is defined by a .proto text file
* You can easily read it and understand it as a human
* Is used at Google for almost all their internal applications
* They have over 48000 Protobuf messages types in 12000 .proto files

* Advantages:
    * Data is fully typed
    * Data is compressed automatically (less CPU)
    * Schema (defined a .proto file) is neded to generate code and read the data
    * Documentation can be embedded in the schema
    * Data can be read across any languagge (C#, Java, Go, Python, Javascript, etc...)
    * Schema can evolve over time, in a safe manner
    * 3-10x smaller, 20-100x faster than XML
    * Code dis generated for you automatically

* Disadvantages:
    * Support for some languages (main ones)
    * Can't see the serialized data with a text editor (because it's compressed and serialized)

## How are Protocol Buffers used?

**To share data across languages**

1. `.proto` file  Human-readable
2. Generate code
    * Java
    * Python
    * Go
    * etc...
3. Create objects
4. Serialized the data (Encode/decode)

* Some databases may have support for Protocol Buffers data format
* Lots RPC frameworks, including gRPC, use Protocol Buffers for exchange data
* Google uses it for all their internal API
* Some big projects like `etcd` use Protocol Buffers for trasporting data

### Proto2 vs Proto3
* Mid 2016, Google release the 3rd iteration of Protocol Buffers, named proto3
* Proto3 it's the easiest to learn

## Course Structure
## Course objectives
1. Write simple and complex `.proto` files
2. Practice Exercises to Confirm the learnings
3. Leverage imports and packages appropriately
4. Generate code using `protoc` in any language
5. Code in Java/Python with Protocol Buffers
6. Understand how Data Evolution works for Protobuf
7. Learn about advanced Protocol Buffers concepts