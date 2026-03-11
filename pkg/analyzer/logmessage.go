package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func isLog(pass *analysis.Pass, call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.ObjectOf(sel.Sel)
	if obj == nil || obj.Pkg() == nil {
		return false
	}

	pkg := obj.Pkg().Path()

	sig, ok := obj.Type().(*types.Signature)
	if !ok {
		return false
	}

	if pkg == "log/slog" {
		if !slogMethods[sel.Sel.Name] {
			return false
		}
		if sig.Recv() == nil {
			return true
		}
		return slogReceivers[sig.Recv().Type().String()]
	}

	if strings.HasPrefix(pkg, "go.uber.org/zap") {
		if !zapMethods[sel.Sel.Name] {
			return false
		}
		if sig.Recv() == nil {
			return false
		}
		return zapReceivers[sig.Recv().Type().String()]
	}

	return false
}

func getMessage(pass *analysis.Pass, call *ast.CallExpr) (string, token.Pos) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", token.NoPos
	}
	obj := pass.TypesInfo.ObjectOf(sel.Sel)
	if obj == nil {
		return "", token.NoPos
	}
	sig, ok := obj.Type().(*types.Signature)
	if !ok {
		return "", token.NoPos
	}

	for i, arg := range call.Args {
		lit, ok := arg.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			continue
		}
		params := sig.Params()
		if params.Len() <= i {
			continue
		}
		basic, ok := params.At(i).Type().Underlying().(*types.Basic)
		if ok && basic.Kind() == types.String {
			return strings.Trim(lit.Value, `"`), lit.Pos()
		}
	}
	return "", token.NoPos
}
