package main

import (
	"fmt"
	"math"
)

func main() {

	var tests = []struct {
		input string
		expr  Expr
		env   Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"sqrt(A / pi)", 
		call {
			op: "sqrt",
			expr: binary {
				op: '/',
				l: Var("A"),
				r: Var("pi") }},
		 Env{"A": 87616, "pi": math.Pi}, "167"},
		// {"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		// {"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		// {"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		// {"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		// {"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		// {"-1 + -x", Env{"x": 1}, "-2"},
		// {"-1 - x", Env{"x": 1}, "-2"},
	}

	for _, test := range tests {
		got := fmt.Sprintf("%.6g", test.expr.Evaluate(test.env))
		fmt.Printf("%v\t%v\t%v\t%v\n", test.input, test.expr, test.want, got)
	}
}