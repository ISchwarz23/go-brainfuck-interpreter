# Brainfuck Interpreter

Brainfuck is an esoteric programming language. For the purpose of getting into Go, I implemented an interpreter for the Brainfuck language.


## Build and run
To build the Brainfuck interpreter run:
```sh
    go build bf.go
```

Afterwards the program can be moved anywhere (depending on the build env) and executed using:
```sh
    ./bfi [<args>] <bf-file-or-code>
```


## Usage
Here are some usage examples on how to run the Brainfuck interpreter:
```sh
./bfi ,>,<[>+<-]            # will interpret the passed bf code
./bfi hello-world.bf        # will interpret hello-world.bf
./bfi -d hello-world.bf     # will also print the memory dump before ending
```
