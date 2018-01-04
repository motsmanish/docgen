package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tdewolff/minify"
	mcss "github.com/tdewolff/minify/css"
	mhtml "github.com/tdewolff/minify/html"
	mjs "github.com/tdewolff/minify/js"
)

var (
	in          string
	out         string
	buildOutput = &cobra.Command{
		Use:   "build",
		Short: "Build html documentation from postman collection",
		Long:  `Build html documentation from postman collection`,
		Run:   buildAngGenerateFile,
	}
)

func init() {
	buildOutput.PersistentFlags().StringVarP(&in, "in", "i", "", "postman collection file relative path")
	buildOutput.PersistentFlags().StringVarP(&out, "out", "o", "", "output file relative path")
}

func buildAngGenerateFile(cmd *cobra.Command, args []string) {
	if in == "" {
		log.Println("You must provide a input file name!")
		return
	}

	if out == "" {
		log.Println("You must provide a output file name!")
		return
	}
	if _, err := os.Stat(in); os.IsNotExist(err) {
		log.Println("Invalid file path!")
		return
	}
	buf := readJSON(in)
	if !strings.HasSuffix(out, ".html") {
		out = out + ".html"
	}
	m := minify.New()
	m.AddFunc("text/html", mhtml.Minify)
	m.AddFunc("text/css", mcss.Minify)
	m.AddFunc("text/javascript", mjs.Minify)
	b, err := m.Bytes("text/html;text/css;text/javascript", buf.Bytes())
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(out, b, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Documentation successfully generated to %s", out)
}
