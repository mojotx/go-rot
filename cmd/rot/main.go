package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mojotx/go-rot/pkg/rot"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		c, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		fmt.Printf("%c", rot.Rot(rune(c)))
	}

}
