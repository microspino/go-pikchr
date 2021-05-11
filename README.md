# go-pikchr
Use [pikchr](https://github.com/drhsqlite/pikchr) (made by Dr. R. Hipp) directly from Go

Given a file with pikchr dsl (a pic-like domain specific language to describe diagrams) it will generate the corresponding svg file.

#### Example

Running this on your shell

    go run . example.txt 
    
...Will generate the graph below

![example](./example.svg)

[BSD Zero Clause License](https://spdx.org/licenses/0BSD.html)