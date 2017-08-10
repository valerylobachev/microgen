package templates_test

import (
	goparser "go/parser"
	"go/token"
	"testing"

	"bytes"

	"io/ioutil"

	"github.com/devimteam/microgen/generator"
	"github.com/devimteam/microgen/generator/template"
	parser "github.com/devimteam/microgen/parser"
)

func TestEndpointsForCountSvc(t *testing.T) {
	src, err := ioutil.ReadFile("service.txt")
	if err != nil {
		t.Fatalf("read source file error: %v", err)
	}
	out, err := ioutil.ReadFile("endpoints.txt")
	if err != nil {
		t.Fatalf("read out file error: %v", err)
	}

	f, err := goparser.ParseFile(token.NewFileSet(), "", src, 0)
	if err != nil {
		t.Errorf("unable to parse file: %v", err)
	}
	fs, err := parser.ParseInterface(f, "StringService")
	if err != nil {
		t.Errorf("could not get interface func signatures: %v", err)
	}
	buf := bytes.NewBuffer([]byte{})
	gen := generator.NewGenerator([]generator.Template{
		&template.EndpointsTemplate{},
	}, fs, generator.NewWriterStrategy(buf))
	err = gen.Generate()
	if err != nil {
		t.Errorf("unable to generate: %v", err)
	}
	if buf.String() != string(out[:]) {
		t.Errorf("Got:\n\n%s\n\nExpected:\n\n%s", buf.String(), string(out[:]))
	}
}
