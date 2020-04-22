package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

type testType struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

type problemType struct {
	Limit    int        `json:"limit"`
	Accuracy int        `json:"accuracy"`
	Tests    []testType `json:"tests"`
}

type serverType struct {
	Language string      `json:"language"`
	Code     string      `json:"code"`
	Problem  problemType `json:"problem"`
}

func makeJSON(lang string, codePath string, testPath string, limit int, accuracy int) []byte {
	codeBytes, err := ioutil.ReadFile(codePath)
	if err != nil {
		log.Fatal(err)
	}
	code := string(codeBytes)
	data := serverType{Language: lang, Code: code, Problem: problemType{
		Limit: limit, Accuracy: accuracy,
	}}
	inDir := path.Join(testPath, "in")
	testIn, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalln(err)
	}
	outDir := path.Join(testPath, "out")
	testOut, err := ioutil.ReadDir(outDir)
	if err != nil {
		log.Fatalln(err)
	}
	for i, inFile := range testIn {
		outFile := testOut[i]
		if !inFile.IsDir() && !outFile.IsDir() {
			in, err := ioutil.ReadFile(path.Join(inDir, inFile.Name()))
			if err != nil {
				log.Fatalln(err)
			}
			out, err := ioutil.ReadFile(path.Join(outDir, outFile.Name()))
			if err != nil {
				log.Fatalln(err)
			}
			data.Problem.Tests = append(data.Problem.Tests, testType{In: string(in), Out: string(out)})
		}
	}
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
