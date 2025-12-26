// Declare the main package
package main

// Import the fmt, io, and os package
import (
	"fmt"
	"io"
	"os"
)

// Declares the CliConfig struct having the ErrStream and Outstream
type CliConfig struct {
	ErrStream io.Writer
	OutStream io.Writer
}

type Option func(*CliConfig) error

func NewCliConfig(opt ...Option) (CliConfig, error) {
	c := CliConfig{
		ErrStream: os.Stderr,
		OutStream: os.Stdout,
	}

	for _, opt := range opt {
		if err := opt(&c); err != nil {
			return CliConfig{}, err
		}
	}

	return c, nil
}

func WithErrStream(errStream io.Writer) Option {
	return func(c *CliConfig) error {
		c.ErrStream = errStream
		return nil
	}
}

func WithOutStream(outStream io.Writer) Option {
	return func(c *CliConfig) error {
		c.OutStream = outStream
		return nil
	}
}



func app(words []string, cfg CliConfig) {
	for _, w := range words {
		if len(w)%2 == 0 {
			fmt.Fprintf(cfg.OutStream, "word %s is even \n", w)
		} else {
			fmt.Fprintf(cfg.ErrStream, "word %s is odd \n", w)
		}
	}
}

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "No words provided.")
		os.Exit(1)
	}
	cfg, err := NewCliConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		os.Exit(1)
	}
	app(words, cfg)
}
