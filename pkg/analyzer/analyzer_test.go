package analyzer_test

import (
	"testing"

	"github.com/zkqw3r/loglinter/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.Analyzer, "sample")
}
