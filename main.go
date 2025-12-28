package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Option func(*CliConfig) error

type CliConfig struct {
	OutputFile string
	ErrStream io.Writer
	OutStream io.Writer
}

func NewCliConfig(opts ...Option) (CliConfig, error) {
	c := CliConfig{
		OutputFile: "",
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

func app(argsValues []string, outputWriter io.Writer) {
	for _, val := range argsValues {
		if len(val)%2 == 0 {
			fmt.Fprintf(outputWriter, "Word %s is even\n", val)
			return
		} else {
			fmt.Fprintf(outputWriter, "Word %s is odd\n", val)
		}
	}
}

func main() {
	var outputFileName string 
	var outputWriter io.Writer
	
	flag.StringVar(&outputFileName, "f", "", "Output file (default:stdout)")
	flag.Parse()
	
	fmt.Println(outputFileName)	
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Fprintf(os.Stderr, "No words provided")
		os.Exit(1) 
		}	
		cfg, err := NewCliConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		}
		if cfg.OutputFile != "" {
			outputFile, err := os.Create(cfg.OutputFile)
		if err != nil {
			fmt.Fprintf(cfg.ErrStream, "Error Creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outputFile.Close()
		outputWriter = io.MultiWriter(cfg.OutStream, outputFile)
	} else {
		outputWriter = cfg.OutStream
	}
	app(words, outputWriter)
}
