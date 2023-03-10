package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Transformer func(state *ParseState, bs []byte) []byte

func getAllBytes(state *ParseState, bs []byte) []byte {
	return append(state.restBytes, bs...)
}

// returns parsed runes and number of parsed bytes
func parseRunes(bs []byte) ([]rune, int) {
	result := make([]rune, 0, utf8.RuneCount(bs))
	resultSize := 0
	for {
		r, size := utf8.DecodeRune(bs)
		if r == utf8.RuneError {
			break
		}
		bs = bs[size:]
		resultSize += size
		result = append(result, r)
	}
	return result, resultSize
}

func toCase(state *ParseState, bs []byte, f func(rune) rune) []byte {
	allBytes := getAllBytes(state, bs)
	runes, bytesNum := parseRunes(allBytes)
	state.restBytes = allBytes[bytesNum:]
	for i, r := range runes {
		runes[i] = f(r)
	}
	return []byte(string(runes))
}

func lowerCaseTransformer(state *ParseState, bs []byte) []byte {
	return toCase(state, bs, unicode.ToLower)
}

func upperCaseTransformer(state *ParseState, bs []byte) []byte {
	return toCase(state, bs, unicode.ToUpper)
}

func trimSpacesTransformer(state *ParseState, bs []byte) []byte {
	allBytes := getAllBytes(state, bs)
	runes, bytesNum := parseRunes(allBytes)
	state.restBytes = allBytes[bytesNum:]
	text := make([]rune, 0)
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			text = append(text, state.spaces...)
			state.spaces = make([]rune, 0)
			state.isBeginning = false
			text = append(text, r)
		} else if !state.isBeginning {
			state.spaces = append(state.spaces, r)
		}
	}
	return []byte(string(text))
}

var conversions = map[string]Transformer{
	"lower_case":  lowerCaseTransformer,
	"upper_case":  upperCaseTransformer,
	"trim_spaces": trimSpacesTransformer,
}

type ParseState struct {
	transformers []Transformer
	restBytes    []byte
	spaces       []rune
	isBeginning  bool
}

type Options struct {
	From      string
	To        string
	Offset    int
	Limit     int
	BlockSize int
	Conv      string
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.From, "from", "", "file to read. by default - stdin")
	flag.StringVar(&opts.To, "to", "", "file to write. by default - stdout")
	flag.IntVar(&opts.Offset, "offset", 0, "number of bytes to skip. by default - 0")
	flag.IntVar(&opts.Limit, "limit", math.MaxInt, "maximum number of bytes to read. by default - max integer value")
	flag.IntVar(&opts.BlockSize, "block-size", 1, "file to write. by default - 1024")
	flag.StringVar(&opts.Conv, "conv", "", "conversions to apply on input data. by default - nothing")

	flag.Parse()

	if opts.Offset < 0 {
		return &opts, errors.New("offset should not be less than 0")
	}
	if opts.Limit < 0 {
		return &opts, errors.New("limit should not be less than 0")
	}
	if opts.BlockSize < 0 {
		return &opts, errors.New("block-size should not be less than 0")
	}

	return &opts, nil
}

func applyTransformers(state *ParseState, bs []byte) []byte {
	result := make([]byte, len(bs))
	copy(result, bs)
	for _, transformer := range state.transformers {
		result = transformer(state, result)
	}
	return result
}

func run(opts Options) error {
	conversionsSet := make(map[string]struct{})
	state := &ParseState{
		transformers: make([]Transformer, 0),
		restBytes:    make([]byte, 0),
		spaces:       make([]rune, 0),
		isBeginning:  true,
	}

	if opts.Conv != "" {
		for _, conversion := range strings.Split(opts.Conv, ",") {
			transformer, ok := conversions[conversion]
			if !ok {
				return errors.New("no such convertor:" + conversion)
			}
			conversionsSet[conversion] = struct{}{}
			state.transformers = append(state.transformers, transformer)
		}
	}

	_, ok1 := conversionsSet["lower_case"]
	_, ok2 := conversionsSet["upper_case"]
	if ok1 && ok2 {
		return errors.New("lower_case conversion can't be used with upper_case conversion")
	}

	reader := os.Stdin
	if opts.From != "" {
		file, err := os.Open(opts.From)
		if err != nil {
			return errors.New("can't open file" + err.Error())
		}
		reader = file
	}
	defer reader.Close()

	writer := os.Stdout
	if opts.To != "" {
		_, err := os.Stat(opts.To)
		if !os.IsNotExist(err) {
			return errors.New("output file is already exist")
		}
		file, err := os.Create(opts.To)
		if err != nil {
			return errors.New("can't open file" + err.Error())
		}
		writer = file
	}
	defer writer.Close()

	for i := 0; i < opts.Offset; {
		block := make([]byte, min(opts.BlockSize, opts.Offset-i))
		bytesNum, err := reader.Read(block)
		i += bytesNum
		if err == io.EOF {
			return errors.New("offset is more than the file size")
		}
		if err != nil {
			return fmt.Errorf("can't read from file:%w", err)
		}
	}

	for i := 0; i < opts.Limit; {
		block := make([]byte, min(opts.BlockSize, opts.Limit-i))
		bytesNum, err := reader.Read(block)
		i += bytesNum
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("can't read from file:%w", err)
		}
		_, err = writer.Write(applyTransformers(state, block[:min(bytesNum, len(block))]))
		if err != nil {
			return fmt.Errorf("can't write in file:%w", err)
		}
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	opts, err := ParseFlags()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)
		os.Exit(1)
	}

	err = run(*opts)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "something went wrong:", err)
		os.Exit(1)
	}
}
