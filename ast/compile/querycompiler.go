package compile

import "avidbound.com/zego/ast"

type QueryCompiler interface {
	Compile(q ast.Body) (ast.Body, error)
}

type queryCompiler struct {
}

func (c queryCompiler) Compile(q ast.Body) (ast.Body, error) {
	return nil, nil
}
