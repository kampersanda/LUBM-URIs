package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/knakk/rdf"
)

func main() {
	mode := flag.String("mode", "", "parse or merge")
	inpath := flag.String("inpath", ".", "help")
	outpath := flag.String("outpath", ".", "help")
	numUniv := flag.Int("num", math.MaxInt32, "numUniv")
	flag.Parse()

	if *mode == "parse" {
		fns, err := filepath.Glob(*inpath + "/*.owl")
		if err != nil {
			panic(err)
		}
		for _, fn := range fns {
			parseUris(fn)
		}
	}

	if *mode == "merge" {
		mergeUris(*inpath, *outpath, *numUniv)
	}
}

func parseUris(fn string) {
	uris := make([]string, 0)

	{
		file, err := os.Open(fn)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		dec := rdf.NewTripleDecoder(file, rdf.RDFXML)
		for triple, err := dec.Decode(); err != io.EOF; triple, err = dec.Decode() {
			uris = appendUris(uris, triple)
		}
		sort.Strings(uris)
	}

	file, err := os.Create(fn + ".uri")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(uris[0] + "\n")
	for i, uri := range uris[1:] {
		if uris[i] != uri {
			file.WriteString(uri + "\n")
		}
	}
}

func mergeUris(inpath, outpath string, numUniv int) {
	fns, err := filepath.Glob(inpath + "/*.owl.uri")
	if err != nil {
		panic(err)
	}

	if numUniv < math.MaxInt32 {
		tmp := make([]string, 0, len(fns))
		for _, fn := range fns {
			idx, err := strconv.Atoi(getIdx(fn))
			if err != nil {
				panic(err)
			}
			if idx < numUniv {
				tmp = append(tmp, fn)
			}
		}
		fns, tmp = tmp, fns
	}

	urisMap := make(map[string]struct{})

	for _, fn := range fns {
		func() {
			file, err := os.Open(fn)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				urisMap[scanner.Text()] = struct{}{}
			}
		}()
	}

	uris := make([]string, 0, len(urisMap))
	for key := range urisMap {
		uris = append(uris, key)
	}
	sort.Strings(uris)

	var outFn string
	if numUniv == math.MaxInt32 {
		outFn = fmt.Sprintf("%s/dataset.txt", outpath)
	} else {
		outFn = fmt.Sprintf("%s/dataset.%d.txt", outpath, numUniv)
	}
	file, err := os.Create(outFn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, uri := range uris {
		file.WriteString(uri + "\n")
	}
}

func getIdx(fn string) string {
	rep := regexp.MustCompile(`\d+`)
	return rep.FindString(fn)
}

func format(term rdf.Term) string {
	s := term.String()
	if s == "" {
		return ""
	}
	if term.Type() == rdf.TermIRI {
		return fmt.Sprintf("<%s>", s)
	}
	if term.Type() == rdf.TermLiteral {
		return fmt.Sprintf("\"%s\"", s)
	}
	return ""
}

func appendUris(uris []string, triple rdf.Triple) []string {
	subj := format(triple.Subj)
	if subj != "" {
		uris = append(uris, subj)
	}
	pred := format(triple.Pred)
	if pred != "" {
		uris = append(uris, pred)
	}
	obj := format(triple.Obj)
	if obj != "" {
		uris = append(uris, obj)
	}
	return uris
}
