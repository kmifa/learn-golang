package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/kmifa/gocloc-copy"
)

// Version is version string for gocloc command
var Version string

// GitCommit is git commit hash string for gocloc command
var GitCommit string

// OutputTypeDefault is cloc's text output format for --output-type option
const OutputTypeDefault string = "default"

// OutputTypeClocXML is Cloc's XML output format for --output-type option
const OutputTypeClocXML string = "cloc-xml"

// OutputTypeSloccount is Sloccount output format for --output-type option
const OutputTypeSloccount string = "sloccount"

// OutputTypeJSON is JSON output format for --output-type option
const OutputTypeJSON string = "json"

const fileHeader string = "File"
const languageHeader string = "Language"
const commonHeader string = "files          blank        comment           code"
const defaultOutputSeparator string = "-------------------------------------------------------------------------" +
	"-------------------------------------------------------------------------" +
	"-------------------------------------------------------------------------"

var rowLen = 79

// CmdOptions is gocloc command options.
// It is necessary to use notation that follows go-flags.
// あとで足していく
type CmdOptions struct {
	ByFile bool `long:"by-file" description:"report results for every encountered source file"`
}

type outputBuilder struct {
	opts   *CmdOptions
	result *gocloc.Result
}

func main() {
	var opts CmdOptions
	clocOpts := gocloc.NewClocOptions()
	// parse command line options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "gocloc"
	parser.Usage = "[OPTIONS] <path>"

	paths, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	// value for language result
	languages := gocloc.NewDefinedLanguages()

	if opts.ByFile {
		fmt.Println("ByFile")
	}

	processor := gocloc.NewProcessor(languages, clocOpts)
	result, err := processor.Analyze(paths)
	if err != nil {
		fmt.Printf("fail gocloc analyze. error: %v\n", err)
		return
	}

}
