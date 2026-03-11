package sample

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	sugar.Info("server started")                // OK
	sugar.Infow("Server started", "port", 8080) // want `Log messages should start with a lowercase letter`
	sugar.Infow("запуск сервера", "port", 8080) // want `Log messages must be in English only`
	sugar.Infow("starting!", "port", 8080)      // want `Log messages should not contain special characters or emojis`
	sugar.Infow("starting 🚀", "port", 8080)     // want `Log messages should not contain special characters or emojis`
	sugar.Infow("123 password", "port", 8080)   // want `Log messages should not contain potentially sensitive data`

	slog.Info("Server started") // want `Log messages should start with a lowercase letter.`
	slog.Info("запуск сервера") // want `Log messages must be in English only.`
	slog.Info("starting!")      // want `Log messages should not contain special characters or emojis.`

	logger, _ = zap.NewProduction()
	logger.Info("Server started") // want `Log messages should start with a lowercase letter.`
}
