package main

import (
	"io/ioutil"
	"log"

	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	_ "github.com/vektah/gqlparser/v2/validator/rules"
)

// func LoadSchema(str ...*ast.Source) (*ast.Schema, *gqlerror.Error) {
// 	return validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
// }

// func MustLoadSchema(str ...*ast.Source) *ast.Schema {
// 	s, err := validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }

func main() {
	b, err := ioutil.ReadFile("../back/graph/schema.graphql")
	if err != nil {
		panic(err)
	}
	source := []*ast.Source{{Name: "graph/schema.graphql", Input: string(b), BuiltIn: false}}
	schema := gqlparser.MustLoadSchema(source...)
	// if err1 != nil {
	// 	panic(err)
	// }
	// json, err := json.Marshal(schema)
	// if err != nil {
	// 	panic(err)
	// }
	// os.Stdout.Write(json)
	for _, v := range schema.Mutation.Fields {
		log.Printf("#### %+v, %+v", v.Name, v.Arguments[0])
	}
}
