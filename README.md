# LUBM-URIs: Generator of URI-string dataset via Lehigh University Benchmark

This library generates a dataset of URI strings via [Lehigh University Benchmark (LUBM)](http://swat.cse.lehigh.edu/projects/lubm/).

## Usage

### Compile the [UBA](http://swat.cse.lehigh.edu/projects/lubm/) data generator.

First of all, please compile the source files as the following command.

```
$ javac -d uba1.7/classes uba1.7/src/edu/lehigh/swat/bench/uba/*.java
```

### Generate URI datasets

You can generate a dataset by extracting URIs from the RDF dataset generated for 10 universities, as the following commands.

```
$ mkdir workspace
$ cd workspace
$ java -cp ../uba1.7/classes edu.lehigh.swat.bench.uba.Generator -univ 10 -onto http://swat.cse.lehigh.edu/onto/univ-bench.owl
$ go get github.com/knakk/rdf
$ go run ../main.go -mode parse
$ go run ../main.go -mode merge
```

If you want to apply more universities, please give the number to the argument of `-univ` (instead of 10). 

As a result (in the case of 10 universities), `dataset.txt` of URIs will be generated at `workspace`, as follows.

```
$ head dataset.txt 
"AssistantProfessor0"
"AssistantProfessor0@Department0.University0.edu"
"AssistantProfessor0@Department0.University1.edu"
"AssistantProfessor0@Department0.University2.edu"
"AssistantProfessor0@Department0.University3.edu"
"AssistantProfessor0@Department0.University4.edu"
"AssistantProfessor0@Department0.University5.edu"
"AssistantProfessor0@Department0.University6.edu"
"AssistantProfessor0@Department0.University7.edu"
"AssistantProfessor0@Department0.University8.edu"
$ wc dataset.txt 
  314870  314870 19239244 dataset.txt
```

## Licensing

This library is free software provided under [GNU General Public License v2.0 License](https://github.com/kampersanda/LUBM-URIs/blob/master/LICENSE), following the original License of [UBA 1.7](http://swat.cse.lehigh.edu/projects/lubm/).
