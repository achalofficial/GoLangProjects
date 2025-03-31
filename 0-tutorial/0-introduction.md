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
