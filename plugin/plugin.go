package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/mitchellh/mapstructure"
	"github.com/zkqw3r/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglinter", New)
}

func New(settings any) (register.LinterPlugin, error) {
	cfg := analyzer.CurrentConfig
	if settings != nil {
		if err := mapstructure.Decode(settings, &cfg); err != nil {
			return nil, err
		}
		analyzer.CurrentConfig = cfg
	}
	return &plugin{}, nil
}

type plugin struct{}

func (*plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (*plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
