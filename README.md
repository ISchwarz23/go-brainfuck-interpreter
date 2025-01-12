# Go Brainfuck Interpreter

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
./bfi -h                    # will print help
./bfi hello-world.bf        # will interpret hello-world.bf
./bfi ,>,<[>+<-]            # will interpret the passed bf code
./bfi -d ,>,<[>+<-]         # will also print memory dump before exiting
```

## Build and Run for wasi
To build the Brainfuck interpreter for wasi just run the following command.
```sh
GOOS=wasip1 GOARCH=wasm go build -o bfi.wasm bfi.go
```

If you have no wasm runtime installed do so. To install `wasmtime` run the follwing command:
```sh
curl https://wasmtime.dev/install.sh -sSf | bash
```

The run the created wasm file using the following command:
```sh
wasmtime bfi.wasm -h
```

If you want to pass a file, you also need to give file access:
```sh
wasmtime --dir=$(pwd) bfi.wasm $(pwd)/hello-world.bf
```