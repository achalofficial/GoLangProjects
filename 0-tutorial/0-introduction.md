# GoLang Tutorial

## Six Main Points about Go
### Statically Typed Language 
    Every variable is defined with either its type or value must be defined so that the type can be 
    infered from decleared variable.
    >   var myVariable string
    >   var myVariable = "Hello" 
### Strongly Type Language
    One type of variable cant be operated with other type of variable. 
### GO is complied
    GO files are compiled into binary file. 
### Fast compile Time
### Build-in Concurrency
    Features are parallelism are in-built through 'GoRoutines' 
### Simplicity
    Syntax is concise, auto-garbage collector
    
## Initiate a new project 
    >  To initiate a new project we write, go mod init <dir-name>

## Contents of go.mod file
    The go.mod file contains the following thing 
        1. Name of the our module 
        2. 'Go' version
        3. Additional/External 'Go' modules which we will install 

## Project structuring 
    1.  Every file in 'Go' project starts with the name of the package. This package name has to be the same for the file in which the package is present. 
    2.  'package main' is the special package name, to tell the compiler to look for the entry point of the package. 

## Running a 'Go' Project
    >   go build <dir-name>
    >   run the <dir-name>.exe file 
    >   go run <dir-name>

## Data Type in GoLang
    1.  const       : 
    2.  variable    :   use var keyword for declaraing an variable. Every variable we need to define the type
    3.  bool        :   
    4.  float32     :
    5.  float64     :
    6.  int         :
    7.  int16       :
    8.  int32       :
    9.  int64       :
    10. rune        :
    11. string      :
    12. uint        :
    13. uint8       :
    14. uint16      :
    15. uint32      :
    16. uint64      :

    >   We can't perform opertions between two different type of variable like int and float type variable can't be added.
    >   We need to cast the variable in same type before performing the Ops.

### String 
    >   "" for single line strings 
    >   `` for multiple line string

#### Operations on string 
##### Length of the string 
    >   var myVar string = "test"
    >   fmt.Println(len(myVar)); // Output is 4
        >   This is not the length but the number of bytes taken by memory
    >   For finding the length of the string 
        >   import "unicode/utf8" // in-built package 
        >   fmt.Println(utf8.RuneCountInString('text-literal')) --> This is return the number of characters in the string
##### Rune in short 
    >   var myRune rune = 'a' 
    >   fmt.Println(myRune) >>> 97

##### Bool in GoLang
    >   the bool are either true or false
    >   the default for bool is 'false'

### New Way to declear the variables in GoLang
    >   instead of 
            var myVar string = 'test'
        we can write 
            var myVar = 'text' >>> Here the type of the variable is inferred.
        we can also do as 
            myVar := 'test'
##### Declearing multiple variables
    >   var var1, var2 int = 1, 2
            OR
    >   var var1, var2 = 1, 2
            OR
    >   var1, var2 := 1, 2

## Constants in GoLang
    >   Can't change the value
    >   Can't declare without initialisation

## Functions and Control Stucture 
### Declaration of function
        >   to declare a function, 'func' keyword is used.
### Passing a parameter in function
        >   func function1 ( printValue string){
            fmt.Println(printValue)
        }
### return type in function
        >   func function2 ( ...<parameter><type>, <parameter><type> ) int {
            <body>
            return <from-body>
        }
        >   func function2 ( ...<parameter><type>, <parameter><type> ) (int, int ){
            <body>
            return <from-body>, <from-body>
        }
## Error type 
    >   initial type of "error" type is 'nil'
    >   import "errors"
    >   var err error
    >   error.New("<error-msg>")
    >   func function3 (...<parameter><type>)(...<return-type>, err)
    >   return ...<return-variable/values>, err

## Conditional Statements
    >   if <condition> {}
    >   if <condition-1><conditional-binary-operator><condition-2>{}
    >   for 'switch-case', we dont need to add the break after each case in GoLang

##  Arrays, Slices and Maps
### Arrays
    >   Fixed length, Same type, INdexable and Contiguous in Memory
    >   var intArr[3]int32  >> [0,0,0]
    >   fmt.Println(intArr[1:3])
    >   Print location of element of Arr    >>> fmt.Println(&intArr[0])     >>> using '&' sign
    >   Immediately initialise 
        >   var intArr[3] int32 = [3]int32{1,2,3}
        >   fmt.Println(intArr)
    >   Other Alternative
        >   intArr := [3]int32{1,2,3}
        >   fmt.Println(intArr)
    >   Another Shorthand 
        >   intArr := [...]int32{1,2,3} >>> since we decleared 3 elements, its an array if size 3
        >   These [...] means that the complier will infer the size of the array by itself 

### Slice 
    >   Slices are just wrappers around arrays.
    >   Slices are just arrays with additional functionlity

    >   by omitting the length value we have slice
        >   var intSlice []int32 = []int32{4,5,6}
    
    >   unlike arrays, we can add values to slices
        >   intSlice = append(intSlice,7)
            >   initially an array is allocated that can hold 3 values.
            >   a check is perform when we try to append, whether the array have room to store value
            >   if not, new array is made with enough capacity to store the value.
            >   we get totally new array with one stored in diffrent memory location.
        >   to know the capacity of slice, we use "cap()" function and 
        >   to get the length we use "len()" function
            >   {1,2,3,5,*,*} >>>>> here cap is 6 and len is 4
        >   adding multiple values to slice, use spread operator
            >   append(intSlice, intSlice2...)
#### MAKE() function
    >   var intSlice3 []int32 = make(<type>, <len>, <capacity(optional, else cap is length of slice)>)

### Map
    >   Map is the set of {'Key':'Value'} pair.
    >   var myMap map[string]uint8 = make(map[string]uint8) >>> declaration
    >   car myMap2 map[string]uint8 = map[string]uint8{'Adam':23, 'Sarah':45}
    >   fmt.Println(myMap['Adam'])  >>> 23
        >   if we try to print a value to key which doesnot exist, then the prints its default value. 
    
    >   var age, ok = myMap["Jason"]
        >   map returns two value, first the value to the key and other a boolean, for existing

    >   to delete a value from map
        >   delete(<map-varName>,<key-needs-to-be-deleted>)

## Loops in GoLang
    >   for name:= range myMap2{
        ftm.Println("Name: %v\n", name)
    }
        >   in map, order of looping element is not preserved.
    
    >   for i,  v := range intArr {
        fmt,Println("Index: %v and value: %v", i, v)
    }
        >   GoLang doesnot have while-Loop 
            >   var i int  = 0 
            >   for <condition> {
                fmt.Println(i)
                <iteration>
            }
        >   Another way 
            >   for {
                if <condition> {
                    break
                }
                fmt.Println(i);
                <iteration>
            }
        >   One More 
            >   for i:= 0 ; i < 10; ++i {
                fmt.Println(i)
            }