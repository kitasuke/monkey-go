package evaluator

import (
	"github.com/kitasuke/monkey-go/ast"
	"github.com/kitasuke/monkey-go/object"
	"github.com/kitasuke/monkey-go/token"
)

var (
	Null  = &object.Null{}
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: val}
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement)

		if result != nil && result.Type() == object.ReturnValueObj {
			return result
		}
	}

	return result
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return Null
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case token.Bang:
		return evalBangOperatorExpression(right)
	case token.Minus:
		return evalMinusPrefixOperatorExpression(right)
	default:
		return Null
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == token.Equal:
		return nativeBoolToBooleanObject(left == right)
	case operator == token.NotEqual:
		return nativeBoolToBooleanObject(left != right)
	default:
		return Null
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case True:
		return False
	case False:
		return True
	case Null:
		return True
	default:
		return False
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.IntegerObj {
		return Null
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case token.Plus:
		return &object.Integer{Value: leftValue + rightValue}
	case token.Minus:
		return &object.Integer{Value: leftValue - rightValue}
	case token.Asterisk:
		return &object.Integer{Value: leftValue * rightValue}
	case token.Slash:
		return &object.Integer{Value: leftValue / rightValue}
	case token.LessThan:
		return nativeBoolToBooleanObject(leftValue < rightValue)
	case token.GreaterThan:
		return nativeBoolToBooleanObject(leftValue > rightValue)
	case token.Equal:
		return nativeBoolToBooleanObject(leftValue == rightValue)
	case token.NotEqual:
		return nativeBoolToBooleanObject(leftValue != rightValue)
	default:
		return Null
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case Null:
		return false
	case True:
		return true
	case False:
		return false
	default:
		return true
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	} else {
		return False
	}
}
