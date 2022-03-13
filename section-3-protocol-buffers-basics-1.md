# Protocol Buffers Basics I

## First message
Create your first message with Protocol Buffers.

```proto
syntax = "proto3"; // <-- sintax definition

// message definition
message MyMessage {
// filedType filedName fieldTag (e.g. Number)
    int32 id = 1;
    string firstName = 2;
    bool isValidated = 3;
}
```

## Scalar type

### Numbers
* double
* float
* int32
* int64
* uint32
* uint64
* sint32
* sint64
* fixed32
* fixed64
* sfixed32
* sfixed64

int - int32

float - float32 (double more precission)

bolean - true or false

string - UTF-8 encoded or 7-bit ASCII text

bytes - sequence of byte array (used for example for files)

Now let's create a message Person with:
* int32 (age)
* string (first name)
* string (last name)
* bytes (small picture)
* bool (profile verified)
* float (height)

## Tags
* In protocol Buffers, field names are not important (but when programming the fields are important)
* For protobuf the important element is the tag
    * Smallest tag: 1
    * Largest tag: 2**29 - 1 or 536870911
* You **cannot use numbers 19000 through 19999**
* Tags numbered from 1 to 15 use **1 byte** in space, so use them for frequently populated fields.
* Tags numbered from 16 to 2047 use **2 bytes**
* There's a concept of reserved tag (advanced lecture)

## Repeated Fields
* To make a "list" you can use the concept of repeated fields
* The list can take any number of elements you want
* **The opposite of repeated is "singular"**

## Comments
* It is possible to embed comments in your .proto file
* It is actually recommended to use comments as a form of documentation for your schemas
* Comments can be of these two forms:
    * `// this is a line`
    * 
    ```
    /* this is a 
    * multiline comment*/
    ```

## Default values for fields
* **All fields, if not specified or unknown, will take a default value**
* bool: false
* number: 0
* string: empty string
* bytes: emoty bytes
* enum: first value
* repeated: empty list

There is not concept of required field.

## Enums
* Enumerations
* enum must start by the tag 0 (default value)