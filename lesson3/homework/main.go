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

type Transformer interface {
	Transform(bs []byte) []byte
}

type TransformersComposer struct {
	transformers []Transformer
}

func (transformer *TransformersComposer) AddTransformer(t Transformer) {
	transformer.transformers = append(transformer.transformers, t)
}

func (transformer *TransformersComposer) Transform(bs []byte) []byte {
	for _, transformer := range transformer.transformers {
		bs = transformer.Transform(bs)
	}
	return bs
}

func NewTransformersComposer() *TransformersComposer {
	return &TransformersComposer{transformers: make([]Transformer, 0)}
}

type CaseTransformer struct {
	state    *ParseState
	function func(r rune) rune
}

func (transformer *CaseTransformer) Transform(bs []byte) []byte {
	allBytes := transformer.state.GetAllBytes(bs)
	runes, bytesNum := parseRunes(allBytes)
	transformer.state.restBytes = allBytes[bytesNum:]
	for i, r := range runes {
		runes[i] = transformer.function(r)
	}
	return []byte(string(runes))
}

type LowerCaseTransformer struct {
	CaseTransformer
}

func NewLowerCaseTransformer(state *ParseState) *LowerCaseTransformer {
	return &LowerCaseTransformer{
		CaseTransformer{state: state, function: unicode.ToLower},
	}
}

type UpperCaseTransformer struct {
	CaseTransformer
}

func NewUpperCaseTransformer(state *ParseState) *UpperCaseTransformer {
	return &UpperCaseTransformer{
		CaseTransformer{state: state, function: unicode.ToUpper},
	}
}

type TrimSpacesTransformer struct {
	state       *ParseState
	spaces      []rune
	isBeginning bool
}

func (transformer *TrimSpacesTransformer) Transform(bs []byte) []byte {
	allBytes := transformer.state.GetAllBytes(bs)
	runes, bytesNum := parseRunes(allBytes)
	transformer.state.restBytes = allBytes[bytesNum:]
	text := make([]rune, 0)
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			text = append(text, transformer.spaces...)
			transformer.spaces = make([]rune, 0)
			transformer.isBeginning = false
			text = append(text, r)
		} else if !transformer.isBeginning {
			transformer.spaces = append(transformer.spaces, r)
		}
	}
	return []byte(string(text))
}

func NewTrimSpacesTransformer(state *ParseState) *TrimSpacesTransformer {
	return &TrimSpacesTransformer{
		state:       state,
		spaces:      make([]rune, 0),
		isBeginning: true,
	}
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

type ParseState struct {
	restBytes []byte
}

func (state *ParseState) GetAllBytes(bs []byte) []byte {
	return append(state.restBytes, bs...)
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

func run(opts Options) error {
	conversionsSet := make(map[string]struct{})
	state := &ParseState{
		restBytes: make([]byte, 0),
	}
	conversions := map[string]Transformer{
		"lower_case":  NewLowerCaseTransformer(state),
		"upper_case":  NewUpperCaseTransformer(state),
		"trim_spaces": NewTrimSpacesTransformer(state),
	}
	transformer := NewTransformersComposer()

	if opts.Conv != "" {
		for _, conversion := range strings.Split(opts.Conv, ",") {
			t, ok := conversions[conversion]
			if !ok {
				return errors.New("no such convertor:" + conversion)
			}
			conversionsSet[conversion] = struct{}{}
			transformer.AddTransformer(t)
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
		_, err = writer.Write(transformer.Transform(block[:min(bytesNum, len(block))]))
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
