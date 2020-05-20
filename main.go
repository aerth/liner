// Copyright 2020 aerth <aerth@riseup.net>

// Liner command skips n-1 lines and prints a single line to stdout
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var VERSION = "liner v0.0.1"

func printVersion() {
	println(VERSION)
}

func printUsage() {
	println(VERSION + ` (https://github.com/aerth/liner)

skips n-1 lines and prints single line to stdout.


Usage of liner:

liner <n> <filename>

  * if n > number of lines, exit with fatal error.
  * use - for stdin
  * line numbers begin at 1

Examples:

cat myfile.txt | liner 3 -
liner -n -x 3 myfile.txt
liner -n -q 3 myfile.txt
`)
}

func main() {
	log.SetFlags(0)
	var (
		hexOutput    = flag.Bool("x", false, "print hex output")
		quoteOutput  = flag.Bool("q", false, "print quoted and escaped output (slightly more human readable than hex)")
		withNewLines = flag.Bool("n", false, "print new line character (in hex or quoted output)")
		showVersion  = flag.Bool("version", false, "print version and exit")
	)
	flag.BoolVar(showVersion, "v", false, "short for --version")

	flag.Usage = func() {
		printUsage()
		println("Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(111)
	}

	var (
		i       = 1
		target  int
		f       *os.File
		err     error
		scanner *bufio.Scanner
	)

	target, err = strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln("fatal: error parsing digit, ", err)
	}

	if target == 0 {
		log.Fatalln("fatal: line numbers start at 1, got 0")
	}

	if flag.Arg(1) == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(flag.Arg(1))
		if err != nil {
			log.Fatalln("fatal:", err)
		}
	}

	scanner = bufio.NewScanner(f)

	for scanner.Scan() {
		if i != target {
			i++
			continue
		}

		// got match
		defer f.Close() // close file when finished
		var data = scanner.Bytes()
		if *withNewLines && (*hexOutput || *quoteOutput) {
			data = append(data, '\n')
		}
		if *hexOutput {
			fmt.Fprintf(os.Stdout, "% 02X\n", data)
			return
		} else if *quoteOutput {
			fmt.Fprintf(os.Stdout, "%q\n", data)
			return
		}

		// copy bytes to stdout as they are, adding the new line
		os.Stdout.Write(data)
		os.Stdout.Write([]byte{'\n'})
		return // success
	}

	// log.Fatal skips any deferred calls
	f.Close()
	log.Fatalln("fatal: reached end of file")
}
