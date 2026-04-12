// Package storage gère le stockage
package main

import (
	"fmt"
	"io"
	"sort"
)

type Stringer interface {
	String() string
}
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (p Persons) Len() int           { return len(p) }
func (p Persons) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p Persons) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Person) String() string {
	return fmt.Sprintf("%s (%d ans)", p.Name, p.Age)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Writer(p []byte) (n int, err error)
}

func Copier(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

func main() {
	persons := Persons{
		Person{"Andy", 27},
		Person{"Micha", 20},
		Person{"Rosy", 26},
	}
	sort.Sort(persons)
	fmt.Println(persons)

}
