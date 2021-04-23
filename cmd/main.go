package main

import (
	"flag"
	"fmt"
	pkg "github.com/maskimko/BashHereDoc/pkg"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	all := flag.Bool("a", false, "list all HereDocument contents")
	token := flag.String("t", "", "get here document by token. "+
		"If token is not provided, the first heredoc occurrence is returned")
	flag.Parse()
	args := flag.Args()
	hereDocs := make(map[string][]byte)
	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("cannot read from StdIn Error %s", err.Error())
		}
		for k, v := range pkg.ParseHereDocs(data) {
			hereDocs[k] = v
		}
	}
	for _, f := range args {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatalf("cannot read from file %s Error %s", f, err.Error())
		}
		for k, v := range pkg.ParseHereDocs(data) {
			hereDocs[k] = v
		}

	}
	if all != nil && *all {
		for t, data := range hereDocs {
			fmt.Printf("%s:\n%s\n\n", t, string(data))
		}
		os.Exit(0)
	}
	if token != nil && *token != "" {
		data, ok := hereDocs[*token]
		if !ok {
			log.Fatalf("token %q has been not found in parsed data", *token)
		}
		fmt.Println(string(data))
		os.Exit(0)
	}
	for _, data := range hereDocs {
		fmt.Println(string(data))
		os.Exit(0)
	}
	os.Exit(2)
}
