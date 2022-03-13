# Protocol Buffers Basics II

## Defining multiple messages in the same file
* It is possible, in the same .proto file, to define multiple types
* It is then super easy to reference them if we need to
* Create a message Date and add that to our person as a field for birthday

[Code](./my-code/basics-ii/1-person-multiple-messages.proto)

## Nesting messages
* It is possible to define types within types
* The reasons could be:
    * Avoiding naming conflicts
    * Enforcing some level of "locality" for that type
* You can nest types as deeply as you want

[Code](./my-code/basics-ii/2-nested-messages.proto)

## Imports
* You can also have different types in different .proto files
* This is useful if you want to re-use code and import other `.proto` files created by people in your team

[Code](./my-code/basics-ii/3-person-with-imports.proto)

## Packages
* It is very important to defined the packages in which your protocol buffer messages live
    * when your code gets compiled, it will be placed at the package you indicated
    * it also helps to prevent name conflicts between messages `(my.package.Person)`
* Packages will help all the different languages compile correctly from `.proto` files (Java, C#, Python, Go, etc...)

[Code](./my-code/basics-ii/4-person-with-package.proto)

## Exercises II
[Code](./my-code/basics-ii/exercises/README.md)