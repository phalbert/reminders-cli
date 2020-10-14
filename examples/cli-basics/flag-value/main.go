package main

import (
	"fmt"
	"flag"
	"time"
	"strings"
)

type idsFlag []string 

func (ids idsFlag) String() string {
	return strings.Join(ids, ",")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
} 

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s, created on %s", p.name, p.born.String())
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	var ids idsFlag
	var p person


	flag.Var(&ids, "id", "the id to add to the list")
	flag.Var(&p, "name", "the name of the person")
	flag.Parse()
	fmt.Println(ids, p)
}
