# LUBM-URIs: Generator of dataset of URI strings via Lehigh University Benchmark

Through this library, you can produce a dataset of URI strings via [Lehigh University Benchmark (LUBM)](http://swat.cse.lehigh.edu/projects/lubm/).

## Usage

### Compile the [UBA](http://swat.cse.lehigh.edu/projects/lubm/) data generator.

First of all, please visit directory `uba1.7` and compile the source files as follows.

```
$ javac -d classes src/edu/lehigh/swat/bench/uba/*.java
```

### Generate URI datasets

Please come back to the project root directory from `uba1.7`.

You can generate a dataset by extracting URIs from the RDF dataset generated for 10 universities, as follows.

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