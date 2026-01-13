# Notes:

# 13-01-2026
- Everything is organised in packages.
- All the code must belong tp a package.
- Packages : https://pkg.go.dev/std

    - # Variables : Camel-case for naming.
        - syntax : var variable_name data_type = value
        - variables can be declared as constants using 'const' keyword. 
            - Constants : Dont change their value once assigned.
        - V are used to store values like containers.
        - update the value only once!
        - Int Bit sizes :
            - int8 : -128 to 127
            - int16 : -32,768 to 32,767
            - int32 : -2,147,483,648 to 2,147,483,647
            - int64 : -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
    - # Data Types :
        - int
        - float32
        - float64
        - string
        - bool
    
# Pointer: 
- A pointer is a variable that stores the memory address of another variable.
- In Go, pointers are represented using the asterisk (*) symbol.

# Arrays:
- An array is a collection of elements of the same type, stored in contiguous memory locations.
- Arrays have a fixed size, which is defined at the time of declaration.
- Syntax : var array_name [size]data_type

# Slices:
- A slice is a dynamically-sized, flexible view into the elements of an array.      
- Slices are more commonly used than arrays in Go, as they provide more flexibility.
- Syntax : var slice_name []data_type

# Maps:
- A map is a built-in data type in Go that represents a collection of key-value pairs
- Maps are also known as associative arrays or dictionaries in other programming languages.
- Syntax : var map_name map[key_data_type]value_data_type

# Functions:
- A function is a block of code that performs a specific task.  
- Functions are defined using the 'func' keyword, followed by the function name, parameters (if any), and return type (if any).
- Syntax : func function_name(parameter1 data_type, parameter2 data_type) return_data_type
- Functions can return multiple values.

# Structs:
- A struct (short for "structure") is a composite data type that groups together related fields.
- Structs are used to create custom data types that can hold multiple values of different types.
- Syntax : type struct_name struct {
                field1 data_type
                field2 data_type
            }

# Methods:
- A method is a function that is associated with a specific type (usually a struct).
- Methods are defined using the 'func' keyword, followed by a receiver (the type the method is associated with), the method name, parameters (if any), and return type (if any).
- Syntax : func (receiver receiver_type) method_name(parameter1 data_type) return_data_type

# Interfaces:
- An interface is a type that defines a set of method signatures.
- Interfaces allow you to define behavior that can be implemented by different types.
- A type implements an interface by implementing all the methods defined in the interface.
- Syntax : type interface_name interface {
                method1(parameter1 data_type) return_data_type
                method2(parameter1 data_type) return_data_type
            }
    -e.g. :
        type Shape interface {
            Area() float64
            Perimeter() float64
        }
        
# Packages:
- A package is a collection of related Go source files that are organized together.
- Packages help in organizing code and promoting code reuse.
- Each Go source file starts with a 'package' declaration that specifies the package name.
- To use functions, types, or variables from another package, you need to import that package using the 'import' keyword.
- Syntax : import "package_name"

# Error Handling:
- In Go, error handling is done using the built-in 'error' type.
- Functions that can encounter errors typically return an 'error' value as the last return value.
- You can check if an error occurred by comparing the returned error value to 'nil'.
- Syntax : 
    result, err := function_name(parameters)
    if err != nil {
        // Handle the error
    }

# Concurrency:
- Go has built-in support for concurrency through goroutines and channels.
- A goroutine is a lightweight thread of execution that allows you to run functions concurrently.
- You can create a goroutine by using the 'go' keyword before a function call.
- A channel is a communication mechanism that allows goroutines to send and receive values.
- Channels are created using the 'make' function.
- Syntax : 
    go function_name(parameters)  // Creating a goroutine
    channel_name := make(chan data_type)  // Creating a channel

