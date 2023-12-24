package dependencie

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, s string) {
	fmt.Fprintf(writer, "Hello, %s", s)
}

func main() {
	Greet(os.Stdout, "l")
}
