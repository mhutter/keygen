package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const usage = `usage: keygen LENGTH

  LENGTH
    length of the key in bytes
`

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) != 1 {
		return errors.New(usage)
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New(usage)
	}

	key := make([]byte, n)

	if _, err := rand.Read(key); err != nil {
		return fmt.Errorf("could not generate random data: %w", err)
	}

	out := base64.StdEncoding.EncodeToString(key)
	fmt.Fprintln(stdout, out)

	return nil
}
