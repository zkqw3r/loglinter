package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      "checks log messages for style violations",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if !isLog(pass, call) {
			return
		}
		msg, pos := getMessage(pass, call)
		if msg == "" || pos == token.NoPos {
			return
		}
		result := &Result{Log: msg}

		if CurrentConfig.CheckUppercase {
			checkUppercase(result)
		}
		if CurrentConfig.CheckLanguage {
			checkLanguage(result)
		}
		if CurrentConfig.CheckSpecialSymbols {
			checkSpecialSymbols(result)
		}
		if CurrentConfig.CheckSensitive {
			checkSensitive(result)
		}

		if len(result.Messages) == 0 {
			return
		}

		pass.Report(analysis.Diagnostic{
			Pos:     pos,
			Message: strings.Join(result.Messages, "; "),
			SuggestedFixes: []analysis.SuggestedFix{{
				Message: "apply fix",
				TextEdits: []analysis.TextEdit{{
					Pos:     pos + 1,
					End:     pos + 1 + token.Pos(len(msg)),
					NewText: []byte(result.Log),
				}},
			}},
		})
	})

	return nil, nil
}
