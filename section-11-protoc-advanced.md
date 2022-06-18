# Section 11: Protoc Advanced

## --decode_raw option

Decodes a bin file.

Bin file
```
*My name"
```
command:
```bash
$ cat my-code/protoc/simple.bin | protoc --decode_raw

# output
1: 42
2: 1
3: "My name"
4: "\001\002\003\004\005\006"
```

## --decode

Decodes a bin file with a protofile.

Bin file
```
*My name"
```

Proto file

```protobuf
syntax = "proto3";

message SimpleMessage {
  int32 id = 1;
  bool is_simple = 2;
  string name = 3;
  repeated int32 sample_list = 4;
}
```

command:
```bash
$ cat my-code/protoc/simple.bin | protoc --decode=SimpleMessage my-code/protoc/simple.proto

# output
id: 42
is_simple: true
name: "My name"
sample_list: 1
sample_list: 2
sample_list: 3
sample_list: 4
sample_list: 5
sample_list: 6
```

In case the message is inside a package

```bash
$ cat my-code/protoc/simple.bin | \
    protoc \
    --decode=simple.SimpleMessage \
    my-code/protoc/simple.proto
```

`--decode=<packageName>.<MessageName>`

## --encode

Encodes a text into a message

message.txt
```
id: 42
is_simple: true
name: "My name"
sample_list: 1
sample_list: 2
sample_list: 3
sample_list: 4
sample_list: 5
sample_list: 6
```

command:
```bash
$ cat my-code/protoc/simple.txt | protoc --encode
=SimpleMessage my-code/protoc/simple.proto

# output
*My name"% 
```