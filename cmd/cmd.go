package main

import (
	"fmt"

	"github.com/brdgme/lordsofvegas"
	"github.com/brdgme/render"
)

func main() {
	g := &lordsofvegas.Game{}
	if err, _ := g.Start(2); err != nil {
		panic(err)
	}
	output, err := render.Term(g.Template(), g)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
