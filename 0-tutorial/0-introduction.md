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
    