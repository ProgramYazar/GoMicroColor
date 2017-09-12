# GoMicroColor
Golang bash color output writer  


# Usage
This library can be use as standart golang io.Writer. Its a simple solution for simple problems on bash environment. Tested on MacOsX but probably works on Linux.

# Example

Basic usage of this library...

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	gcolor "github.com/ProgramYazar/GoMicroColor"
)

func listen(ch chan string) {
	// better error output
	errorOut := gcolor.NewColorWriter(os.Stderr, gcolor.Red, gcolor.BackgroundGreen)

	tw := io.MultiWriter(
		os.Stderr,
		os.Stdout,
		errorOut,
	)
	fmt.Fprintf(tw, "received data: %s\n", <-ch)

}

func main() {
	var ch = make(chan string)
	go listen(ch)
	ch <- "Error line"
	close(ch)

	greenOutput := gcolor.NewColorWriter(os.Stderr, gcolor.Green, gcolor.BackgroundDefault)
	logger := log.New(greenOutput, "", log.LstdFlags)
	logger.Println("logger can log this")

}

```
