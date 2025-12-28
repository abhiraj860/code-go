package main

import (
	"fmt"
	"io"
	"os"
)

type Option func(*CliConfig) error

type CliConfig struct {
	ErrStream io.Writer
	OutStream io.Writer
}

func NewCliConfig(opts ...Option) (CliConfig, error) {
	c := CliConfig{
		ErrStream: os.Stderr,	
		OutStream: os.Stdout,
	}

	for _, op := range opts {
		if err := op(&c); err != nil {
			return CliConfig{}, err
		}
	}

	return c, nil
}

func UpdateErrStream(inpErrStream io.Writer) Option {
	return func(c * CliConfig) error {
		c.ErrStream = inpErrStream
		return nil
	}
}

func UpdateOutStream(outStream io.Writer) Option {
	return func(c * CliConfig) error {
		c.OutStream = outStream
		return nil
	}
}

func app(argsValues []string, cliconfig CliConfig) {
	for _, val := range argsValues {
		if len(val)%2 == 0 {
			fmt.Fprintf(cliconfig.OutStream, "Word %s is even\n", val)
			return
		} else {
			fmt.Fprintf(cliconfig.ErrStream, "%Word %s is odd\n", val)
		}
	}
}

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Fprintf(os.Stderr, "No words provided")
		os.Exit(1) 
	}	
	cfg, err := NewCliConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
	}
	app(words, cfg)
}
