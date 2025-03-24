# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
- Commence      : 09:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 03:00 PM (20 mins)
- Wind up       : 05:00 PM

## Methodology
- No powerpoint
- Code & Discuss

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code OR any editor
    - Go extension (https://marketplace.visualstudio.com/items?itemName=golang.Go)

### Verification
```shell
go version
```

## Repository
- https://github.com/tkmagesh/cisco-gofoundation-mar-2025

## Why Go?
- Simplicity
    - Only one way doing things
        - var, :=
        - if else, switch case
        - for
        - function
        - type
        - package
    - ONLY 25 keywords
    - No access modifiers
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No classes (Only structs)
    - No inheritance (Only composition)
    - No exceptions (Only errors & errors are just values)
    - No try..catch..finally
    - No implicit type conversion
- Performance
    - Comparable with C++
    - Close to hardware
        - Build for specific platform
        - Compiler has native support for cross compilation
- Managed Concurrency
    - Builtin Scheduler to manage concurrent operations
    - "Goroutine" for concurrent operations
        - Cheap (~2KB) vs OS Thread (~4MB)
    - Concurrency support is built in the language
        - "go" keyword
        - "chan" datatype
        - "<-" operator
        - "range" construct
        - "select-case" construct
    - Standard library support
        - "sync" package
        - "sync/atomic" package

## Compilation
```shell
go build [filename.go]
```

## Compile & Execute
```shell
go run [filename.go]
```

## List the go environment variables
```shell
go env
```

## List any specific environment variables
```shell
go env [var_1] [var_2] ....
```

## Change the env variables
```shell
go env -w [var_1]=[value_1] [var_2]=[value_2] ....
```
## Cross compilation
### Env variables for cross-compilation
- GOOS
- GOARCH

### Get the list of supported platforms for cross compilation
```shell
go tool dist list
```

### How to cross compile?
```shell
go env -w GOOS=[target_os] GOARCH=[target_arch]
go build [filename.go]
```
OR
```shell
GOOS=[target_os] GOARCH=[target_arch] go build [filename.go]
```