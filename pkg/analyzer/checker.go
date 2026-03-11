package analyzer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func checkUppercase(result *Result) {
	r, size := utf8.DecodeRuneInString(result.Log)
	if size == 0 || !unicode.IsLetter(r) || !unicode.IsUpper(r) {
		return
	}
	result.Messages = append(result.Messages,
		"Log messages should start with a lowercase letter")
	result.Log = strings.ToLower(string(r)) + result.Log[size:]
}

func checkLanguage(result *Result) {
	for _, r := range result.Log {
		if r > 127 && unicode.IsLetter(r) {
			result.Messages = append(result.Messages,
				"Log messages must be in English only")
			return
		}
	}
}

func checkSpecialSymbols(result *Result) {
	for _, r := range result.Log {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			result.Messages = append(result.Messages,
				"Log messages should not contain special characters or emojis")
			return
		}
	}
}

func checkSensitive(result *Result) {
	lower := strings.ToLower(result.Log)
	for _, kw := range CurrentConfig.Sensitive {
		if strings.Contains(lower, kw) {
			result.Messages = append(result.Messages,
				"Log messages should not contain potentially sensitive data")
			return
		}
	}
}
