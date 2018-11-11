package evaluator

import (
	"fmt"

	"github.com/kitasuke/monkey-go/object"
)

const (
	BuiltinFuncNameLen   = "len"
	BuiltinFuncNameFirst = "first"
	BuiltinFuncNameLast  = "last"
	BuiltinFuncNameRest  = "rest"
	BuiltinFuncNamePush  = "push"
	BuiltinFuncNamePuts  = "puts"
)

var builtins = map[string]*object.Builtin{
	BuiltinFuncNameLen: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to %q not supported, got %s", BuiltinFuncNameLen, args[0].Type())
			}
		},
	},
	BuiltinFuncNameFirst: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("argument to %q must be %s, got %s", BuiltinFuncNameFirst, object.ArrayObj, args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return Null
		},
	},
	BuiltinFuncNameLast: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("argument to %q must be %s, got %s", BuiltinFuncNameLast, object.ArrayObj, args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return Null
		},
	},
	BuiltinFuncNameRest: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("argument to %q must be %s, got %s", BuiltinFuncNameRest, object.ArrayObj, args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return Null
		},
	},
	BuiltinFuncNamePush: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("argument to %q must be %s, got %s", BuiltinFuncNamePush, object.ArrayObj, args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	BuiltinFuncNamePuts: {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return Null
		},
	},
}
