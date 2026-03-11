package analyzer

type Config struct {
	CheckUppercase      bool     `mapstructure:"check_uppercase"`
	CheckLanguage       bool     `mapstructure:"check_language"`
	CheckSpecialSymbols bool     `mapstructure:"check_special_symbols"`
	CheckSensitive      bool     `mapstructure:"check_sensitive"`
	Sensitive           []string `mapstructure:"sensitive"`
}

var CurrentConfig = Config{
	CheckUppercase:      true,
	CheckLanguage:       true,
	CheckSpecialSymbols: true,
	CheckSensitive:      true,
	Sensitive: []string{
		"password", "passwd", "token", "secret",
		"api_key", "apikey", "auth", "key", "pwd",
		"pass", "token", "credential", "jwt",
		"private_key", "certificate", "cert",
		"credit_card", "bank_account",
	},
}

type Result struct {
	Messages []string
	Log      string
}

var zapReceivers = map[string]bool{
	"*go.uber.org/zap.Logger":        true,
	"go.uber.org/zap.Logger":         true,
	"*go.uber.org/zap.SugaredLogger": true,
	"go.uber.org/zap.SugaredLogger":  true,
}

var zapMethods = map[string]bool{
	"Info": true, "Infow": true, "Infof": true,
	"Error": true, "Errorw": true, "Errorf": true,
	"Warn": true, "Warnw": true, "Warnf": true,
	"Debug": true, "Debugw": true, "Debugf": true,
}

var slogReceivers = map[string]bool{
	"*log/slog.Logger": true,
	"log/slog.Logger":  true,
}

var slogMethods = map[string]bool{
	"Info": true, "Debug": true, "Warn": true, "Error": true,
	"InfoContext": true, "DebugContext": true,
	"WarnContext": true, "ErrorContext": true,
}
