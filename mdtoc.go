package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

const HeadingType = "*ast.Heading"

func PrintToC(md []byte, maxLevel int) {
	extentsion := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extentsion)
	doc := p.Parse(md)
	for _, c := range doc.GetChildren() {
		h, ok := c.(*ast.Heading)
		if !ok {
			continue
		}
		if h.Level > maxLevel {
			continue
		}

		pre := ""
		for i := h.Level - 1; i > 0; i-- {
			pre += "  "
		}
		pre += "- "

		hid := strings.TrimSpace(string(h.HeadingID))
		hText := ""

		for _, c := range h.GetChildren() {
			if c.AsLeaf() == nil {
				continue
			}
			leaf := c.AsLeaf()
			if len(leaf.Content) > 0 && len(leaf.Literal) > 0 {
				hText += string(leaf.Content) + fmt.Sprintf("`%s`", string(leaf.Literal))
			} else {
				hText += string(leaf.Content)
				hText += string(leaf.Literal)
			}
		}
		fmt.Printf("%s[%s](#%s)\n", pre, hText, hid)
	}
}

func main() {
	depth := flag.Int("d", 2, "Headers level to include in ToC up to")
	flag.Usage = func() {
		fmt.Fprint(
			flag.CommandLine.Output(),
			"Usage:\nmdtoc [OPTIONS] <filepath> \n\nOPTIONS:\n",
		)
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Please provide a file path.")
	}
	fp := args[0]

	md, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't open file at: %s. Error: %v", fp, err))
	}
	PrintToC(md, *depth)
}
