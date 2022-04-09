# Setting up Protoc Compiler

## Setup Protoc Compiler

`prtoc` should be installed in your system in order to use the code generator.

### MacOS
`brew install protobuf` 

### Ubuntu

Find the correct protocol buffers version based on your Linux Distro: https://github.com/google/protobuf/releases

## Windows

Download the windows archive: https://github.com/google/protobuf/releases

Example: https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-win32.zip

Extract all to C:\proto3  

Your directory structure should now be

C:\proto3\bin 

C:\proto3\include 

Finally, add C:\proto3\bin to your PATH:

    From the desktop, right-click My Computer and click Properties.
    In the System Properties window, click on the Advanced tab
    In the Advanced section, click the Environment Variables button.

    Finally, in the Environment Variables window (as shown below), highlight the Path variable in the Systems Variable section and click the Edit button. Add or modify the path lines with the paths you wish the computer to access. Each different directory is separated with a semicolon as shown below.

        C:\Program Files; C:\Winnt; ...... ; C:\proto3\bin

(you need to add ; C:\proto3\bin  at the end)

## Use `protoc` to generate code in any language

## Python
```bash
protoc \
-I=./my-code/using-protoc/proto \
--python_out=./my-code/using-protoc/python \
./my-code/using-protoc/proto/*.proto
```

## Java
```bash
protoc \
-I=./my-code/using-protoc/proto \ # Route source of the proto files
--java_out=./my-code/using-protoc/java \
./my-code/using-protoc/proto/*.proto # which files will be taken from Input (-I)
```

## Practice using `protoc`

## Javascript

```bash
protoc \
-I=./my-code/using-protoc/proto \
--js_out=./my-code/using-protoc/javascript \
./my-code/using-protoc/proto/*.proto
```