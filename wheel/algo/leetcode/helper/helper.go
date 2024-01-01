package helper

import (
	"bfw/internal/json"
	"bfw/wheel/algo/leetcode/lc_1"
	"errors"
	"os"
	"strconv"
)

// read opt.json and input.json
// calculate
// write output.json
// diff expect.json and output.json

var (
	optJSONFileError   = errors.New("opt.json file error")
	inputJSONFileError = errors.New("input.json file error")
)

type lc146 struct {
	opt    string
	input  []int
	output string
}

func AutoMachineLC146(dir, suffix, output, opt, input string) {
	output = getFilePath(dir, output, suffix)
	opt = getFilePath(dir, opt, suffix)
	input = getFilePath(dir, input, suffix)
	lcs := ReadOptAndInputJSONFile(opt, input)
	lcs = CalculateLC146(lcs)
	outputsArr := getOutputsFromLCS(lcs)
	WriteOutputJSONFile(output, outputsArr)
}

func CalculateLC146(lcs []lc146) []lc146 {
	var c lc_1.LRUCache
	for i, lc := range lcs {
		switch lc.opt {
		case "LRUCache":
			{
				if len(lc.input) == 1 {
					c = lc_1.Constructor(lc.input[0])
					lcs[i].output = "null"
				} else {
					panic(inputJSONFileError)
				}
			}
		case "put":
			{
				if len(lc.input) == 2 {
					c.Put(lc.input[0], lc.input[1])
					lcs[i].output = "null"
				} else {
					panic(inputJSONFileError)
				}
			}
		case "get":
			{
				if len(lc.input) == 1 {
					lcs[i].output = strconv.Itoa(c.Get(lc.input[0]))
				} else {
					panic(inputJSONFileError)
				}
			}
		default:
			{
				panic(optJSONFileError)
			}
		}
	}
	return lcs
}

func MakeUpLC146(lcs []lc146, opts []string, inputs [][]int, outputs []string) []lc146 {
	optLen, inputLen, outputLen := len(opts), len(inputs), len(outputs)
	maxLen := max(optLen, inputLen, outputLen)
	if len(lcs) == 0 {
		lcs = make([]lc146, maxLen)
	}
	for i, _ := range lcs {
		if optLen > i {
			lcs[i].opt = opts[i]
		}
		if inputLen > i {
			lcs[i].input = inputs[i]
		}
		if outputLen > i {
			lcs[i].output = outputs[i]
		}
	}
	return lcs
}

func ReadOptAndInputJSONFile(optFilePath, inputFilePath string) []lc146 {
	var (
		opts   []string
		inputs [][]int
	)
	GetObjectByJSONFile(optFilePath, &opts)
	GetObjectByJSONFile(inputFilePath, &inputs)
	return MakeUpLC146(nil, opts, inputs, nil)
}

func WriteOutputJSONFile(filePath string, outputs []string) {
	filePtr, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer func(opt *os.File) {
		err := opt.Close()
		if err != nil {
			panic(err)
		}
	}(filePtr)
	outputJSONStr, err := json.MarshalToString(outputs)
	if err != nil {
		panic(err)
	}
	_, err = filePtr.WriteString(outputJSONStr)
	if err != nil {
		panic(err)
	}
}

func GetObjectByJSONFile(filePath string, obj interface{}) {
	filePtr, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func(opt *os.File) {
		err := opt.Close()
		if err != nil {
			panic(err)
		}
	}(filePtr)
	dec := json.NewDecoder(filePtr)
	err = dec.Decode(&obj)
	if err != nil {
		panic(err)
	}
}

func getFilePath(dirName, fileName, suffix string) string {
	return dirName + "/" + fileName + "." + suffix
}

func getOutputsFromLCS(lcs []lc146) []string {
	outputs := make([]string, len(lcs))
	for i, lc := range lcs {
		outputs[i] = lc.output
	}
	return outputs
}
