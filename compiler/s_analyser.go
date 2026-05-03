package compiler

func Evaluate(n Node) (int64, error) {
	switch n := n.(type) {

	case *IntegerLiteral:
		return n.Value, nil
	case *InfixExpression:
		l, err := Evaluate(n.Left)
		if err != nil {
			return 0, err
		}
		r, err := Evaluate(n.Right)
		if err != nil {
			return 0, err
		}

		switch n.Operator {

		case "+":
			return l + r, nil

		case "-":
			return l - r, nil
		case "*":
			return l * r, nil
		case "/":
			if r == 0 {
				return 0, &DivisionByZero{msg: "Semantic Error, Division by Zero!"}
			}
			return l / r, nil
		}
	}
	return 0, nil
}

type DivisionByZero struct {
	msg string
}

func (e *DivisionByZero) Error() string {
	return e.msg
}
