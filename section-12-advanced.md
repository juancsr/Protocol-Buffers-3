# Protocol Buffers Advanced

## Integer Types Deep Dive

Keep in mind:
    1. The range (32 or 64 bits)
    2. Signed or Unsigned
    3. Varint or not

### The range

Common
* int32: -2.147.483.648 to 2.147.483.647
* int64 ...

*int32 & int64 are not efficient at serializing them

Only positive values
* uint32: 0 to 4.294.967.295
* uint64: 0 to ...

Accepts negative values
* sint32
* sint65

*Less efficient at serializing positive values

Varint accetps float
fixed32, sfixed32: 4 bytes
fixed64, sfixed64: 8 bytes

## Advanced Data Types (oneof, map, Timestamp & Duration)

### oneof

```protobuf
syntax = "proto3";

message Account {
    // ...
    oneof phone_or_email {
        string phone = 3;
        string email = 4;
    }
}
```

It'll take value for only one of them. If both are sending `phone` and `email` bigger tag number will take priority. In this case, only `email` will have value.

**It is useful bit evolving the schema could be more difficult**

### map

```protobuf
syntax = "proto3";

message Address {
    // ...
}

message Account {
    // ...
    map<string, Address> address = 3;
}
```

1. Cannot be `repeated`
2. Cannot use `float`, `double` and `enums/message` as key
3. No ordering

### Timestamp & Duration

```protobuf
syntax = "proto3";

import "google/protobuf/duration.proto";
import "google/protobuf/timestamp.proto";
// ...

message Account {
    // ...
    google.protobuf.Timestamp created_at = 3;
    google.protobuf.Duration validity = 4;
}
```

### More documentation

Extra values defined by google proto are here:
https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#index

## Protocol Buffers Options

https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto

## Naming Conventions

1. protofiles should be in `lower_snake_case`
2. `/* License should be inside comments at the begining of the file */`
3. `syntax = "package3";`
4. `package mypackage;`
5. order of import
```protobuf 
// alphabet order
import "a.proto"; // a
import "b.proto"; // b
import "c.proto"; // c
import "google/protobuf/empty.proto"; // g
```
6. `option java_package = "bla.bla.bla"`
7. enums
```protobuf
enum FooBar {
    // Use UPPER_CAMEL_CASE_FOR_ENUM_VALUES
    FOO_BAR_UNSPECIFIED = 0;
    FOO_BAR_FIRST = 1;
    FOO_BAR_SECOND = 2;
}
```
8. messages
```protobuf
// Name of messages should start with UpperCase
message Message {
    // Use lower_camel_case_for_tags
    FooBar a_simple_foobar = 1;
    // use plural when repeated
    repeated string multiple_names = 2;
}
```
9. services
```protobuf
service FooService {
    // convention for rpc: MyFunctionName
    rpc ListSomething(google.protobuf.Empty) returns (Message);
}
```

## Uber style guiding

https://github.com/uber/prototool/blob/dev/etc/style/uber1/uber1.proto

## Services

Services are designed for communication instead of serialization and desserialization. They are basically a set of endpoints that are defining an API.

```protobuf
// we define contracts

message GetSomethingRequest {
    // ...
}

message GetSomethingResponse {
    // ...
}

message ListSomethingRequest {
    // ...
}

message ListSomethingResponse {
// ...
}


// recives a message and returns a message
service FooService {
    rpc GetSomething(GetSomethingRequest) returns (GetSomethingResponse);
    rpc ListSomething(ListSomethingRequest) returns (ListSomethingResponse);
}
```

## Introduction to gRPC

Nop. check the course.

## Protocol Buffers Internals

Hex: AC 

Binary: 1010 1100

Hex: 02

Binary: 0000 0010

## List of Protocol Buffer Files to Explore

You should explore already existing Protocol Buffers file made by Google:

    Main repository examples: https://github.com/protocolbuffers/protobuf/tree/main/examples

    Some Google Apis types: https://github.com/googleapis/googleapis/tree/master/google/type

    Protocol Buffer itself (may be very complex) https://github.com/protocolbuffers/protobuf/tree/main/src/google/protobuf

Happy learning!