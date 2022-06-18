# Section 10: Data evolution with Protobuf

## The need to updating the protocol

We may need to change our message structures because new requirements of the business.
The real problem here is **we don't want breaking changes**

### Backward compatible

```proto
// Wirte
syntax = "proto3";

message Account {
    string first_name = 1;
    string last_name = 2;
}

// Read
syntax = "proto3";
message Account {
    string first_name = 1;
    string last_name = 2;
    string phone = 3;
}
```

In the above case, the server sends `first_name` & `last_name`. The client reads (deserialize) those two values correctly, bur `phone` will take its default value ("" emtpy string).

### Forward compatible

```proto
// Wirte
syntax = "proto3";

message Account {
    string first_name = 1;
    string last_name = 2;
    string phone = 3;
}

// Read
syntax = "proto3";
message Account {
    string first_name = 1;
    string last_name = 2;
}
```
Above case, the server sends the three values, but client only reads `first_name` & `last_name`. `phone` it'll be ignored.

## Rules for data Evolution

1. Do not change tags! `(string first_name = <tag>)`
2. Add new fields
3. Use reserved tags
4. Before changing type:
    Check the doc: https://developers.google.com/protocol-buffers/docs/proto3?hl=en#updating
5. Sometimes it is better to add a new field, instead of chaging the type.

## Renaming Fields

```proto
// Server
syntax = "proto3";

message Account {
    uint32 id = 1;
    string first_name = 2;
}

// Client
syntax = "proto3";

message Account {
    uint32 id = 1;
    string alias = 2;
}
```

In the above code deserialization will be correct due `first_name` and `alias` have the same data type. 

**Just keep in mind that `first_name` will be deserilized in `alias` and vicesersa.**

## Removing Fields

```proto
// Server
syntax = "proto3";

message Account {
    uint32 id = 1;
    string first_name = 2;
}

// Client
syntax = "proto3";

message Account {
    reserved 2;
    reserved "first_name"; // optional
    uint32 id = 1;
}
```

we can use `reserved` in order to reserve a tag or the name of a property.

## Reserved Keyword

```proto
syntax = "proto3";

message Account {
    reserved 2, 15, 9 to 11;
    reserved "first_name", "last_name";
    uint32 id = 1;
}
```

## Beware of Defaults

### The good

Allows us:
* Forward compatibility
* Backward compatibility
* Prevents of having non null values

### The dark side

* Cannot differentiate missing or unset

### What can i do

* Think if the value has no meaning in your business when is default (empty string for example)
* `if` or `switch` statements in your code