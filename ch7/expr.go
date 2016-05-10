package eval

import (
	"strconv"
	"math"
)

type Env map[Var]float64

// TODO String()

type Expr interface {
	Evaluate(env Env) float64
	String() string
}

// Variable
type Var string

func (v Var) Evaluate(env Env) float64 {
	return env[v]
}

func (v Var) String() string {
	return string(v);
}

// literal
type literal float64

func (l literal) Evaluate(env Env) float64 {
	return float64(l)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', 6, 64)
}

// unary
type unary struct {
	op rune // + -
	expr Expr
}

func (u unary) Evaluate(env Env) float64 {
	if (u.op == '+') {
		return u.expr.Evaluate(env)
	} else if (u.op == '-') {
		return u.expr.Evaluate(env) * (-1)
	}
	panic("unknown operator: " + string(u.op))
}

func (u unary) String() string {
	return string(u.op) + u.expr.String()
}

// binary
type binary struct {
	op rune // + - * /
	l, r Expr
}

func (b binary) Evaluate(env Env) float64 {
	var lv = b.l.Evaluate(env)
	var rv = b.r.Evaluate(env)

	if (b.op == '+') {
		return lv + rv;
	} else if (b.op == '-') {
		return lv - rv;
	} else if (b.op == '*') {
		return lv * rv;
	} else if (b.op == '/') {
		return lv / rv;
	}
	panic("unknown operator: " + string(b.op))
}

func (b binary) String() string {
	return b.l.String() + string(b.op) + b.r.String()
}

// call (sin, sqrt, pow)
type call struct {
	op string
	expr Expr
}

func (c call) Evaluate(env Env) float64 {
	var ev = c.expr.Evaluate(env)

	if (c.op == "sin") {
		return math.Sin(ev);
	} else if (c.op == "sqrt") {
		return math.Sqrt(ev);
	} else if (c.op == "pow") {
		return math.Pow(ev, 2);
	}
	panic("unknown operator " + c.op)
}

func (c call) String() string {
	return c.op + "(" + c.expr.String() + ")"
}
