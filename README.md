# logger
Logger written in Go

## Installation
Add package to your Go project
```bash
go get github.com/Pilladian/logger
```

## Example
```go
package main

import "github.com/Pilladian/logger"

func main() {
    // Possible log level between 0-2
    // [0] Errors are getting logged,
    // [1] Errors and Warnings are getting logged,
    // [2] Errors, Warnings and Infos are getting logged
    logger.SetLogLevel(2)

    // Set filename for logging output 
    // default: "" (logs will be written to std.out)
    logger.SetLogFilename("./output.log")

    // Log Info
    logger.Info("Some Info to log")

    // Log Warning
    logger.Warning("Some Warning to log")

    // Log Error
    logger.Error("Some Error to log")
    
    // Log Error and exit with status code 1
    logger.Fatal("Some Error to log")
}
```
