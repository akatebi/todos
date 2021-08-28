package main

import (
	"io/ioutil"
	"log"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vektah/gqlparser/v2/validator"
	_ "github.com/vektah/gqlparser/v2/validator/rules"
)

func LoadSchema(str ...*ast.Source) (*ast.Schema, *gqlerror.Error) {
	return validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
}

func MustLoadSchema(str ...*ast.Source) *ast.Schema {
	s, err := validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	b, err := ioutil.ReadFile("../back/graph/schema.graphql")
	if err != nil {
		panic(err)
	}
	source := []*ast.Source{{Name: "graph/schema.graphql", Input: string(b), BuiltIn: false}}
	schema := MustLoadSchema(source...)
	log.Printf("#### %+v", schema)
}
