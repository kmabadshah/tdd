package main

import (
	"fmt"
	"io"
)

func InjGreet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)

	if err != nil {
		return err
	}
	return nil
}
