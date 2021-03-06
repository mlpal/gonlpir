/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file cmdline.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Nov 23 20:59:30 2014
 *
 **/

package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

var (
	Help           bool
	Verbose        bool
	DataPath       string
	InputEncoding  string
	OutputEncoding string
	LineDelimiter  string
	FieldDelimiter string
	ShowPOS        bool
)

func InitFlags() {
	flag.BoolVar(&Help, "help", false, "Print help message")
	flag.BoolVar(&Help, "h", false, "Print help message")
	flag.BoolVar(&Verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&Verbose, "v", false, "Use verbose output")
	flag.StringVar(&DataPath, "data", "", "ICTCLAS data path")
	flag.StringVar(&DataPath, "d", "", "ICTCLAS data path")
	flag.StringVar(&InputEncoding, "ie", "UTF8", "Encoding of input stream")
	flag.StringVar(&InputEncoding, "input-encoding", "UTF8", "Encoding of input stream")
	flag.StringVar(&OutputEncoding, "oe", "UTF8", "Encoding of output stream")
	flag.StringVar(&OutputEncoding, "output-encoding", "UTF8", "Encoding of output stream")
	flag.StringVar(&LineDelimiter, "ld", "\n", "Line delimiter")
	flag.StringVar(&LineDelimiter, "line-delimiter", "\n", "Line delimiter")
	flag.StringVar(&FieldDelimiter, "fd", "\t", "Field delimiter")
	flag.StringVar(&FieldDelimiter, "field-delimiter", "\t", "Field delimiter")
	flag.BoolVar(&ShowPOS, "pos", false, "Show POS (part of speech) info")
}

func Parse() {
	flag.Parse()
	if Help {
		showHelp(0)
	}
	if len(DataPath) == 0 {
		fmt.Println("invalid arguments: no data")
		showHelp(1)
	}
	OutputEncoding = strings.ToLower(OutputEncoding)
}

//===================================================================
// Private
//===================================================================

const (
	logoString = `                        __
.--.--.--.-----.----.--|  |.-----.-----.-----.
|  |  |  |  _  |   _|  _  ||__ --|  -__|  _  |
|________|_____|__| |_____||_____|_____|___  |
                                       |_____|
`
	helpString = `wordseg
Usage:
    wordseg [options]
Options:
    -h, --help               Print this message
    -v, --verbose            Use verbose output

    -d, --data               Data path

    --ie, --input-encoding   Specify encoding of input stream
                             available: gbk, utf8, big5, gbk_fanti, utf8_fanti
    --oe, --output-encoding  Specify encoding of output file
                             available: gbk, utf8, big5 and other codenames

    --ld, --line-delimiter   Specify line delimiter, default: \n
    --fd, --field-delimiter  Specify field delimiter, default: \t

    --pos                    Show POS (Part of speech) info
`
)

func showHelp(ret int) {
	fmt.Print(helpString)
	os.Exit(ret)
}

func LogoString() string {
	return logoString
}
