// example2.go
package main

import (
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*parser.BaseCalcListener
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, p.Start())
}
