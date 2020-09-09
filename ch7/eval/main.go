package main

import (
	"fmt"
	"math"
)

// An Expr is an arithmetic expression
type Expr interface {
	// Eval returns the value of this Expr in the Environment env.
	Eval(env Env) float64
}

// A Var identifies a variable, e.g., x.
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// A literal is a numberic constant, e.g., 3.141
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// A unary represents a unary operator expression, e.g., -x
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Errorf("unsupported unary operator: %q", u.op))
}

// A binary represents a binary opeartor expression, e.g., x+y
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Errorf("unsupported binary opeartor: %q", b.op))
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Errorf("unsupported function call: %s", c.fn))
}

type Env map[Var]float64

func main() {

}
