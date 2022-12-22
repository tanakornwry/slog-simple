package main

import (
	"net"
	"os"

	"golang.org/x/exp/slog"
)

var opts = slog.HandlerOptions{
	// AddSource: true,
	// Level:     slog.DebugLevel,
	// ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
	// 	if a.Key == slog.TimeKey {
	// 		return slog.Attr{}
	// 	}
	// 	return a
	// },
}

func slogAsText() {
	slog.SetDefault(slog.New(opts.NewTextHandler(os.Stdout)))
	logging()
}

func slogAsJSON() {
	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stdout)))
	logging()
}

func logging() {
	slog.Debug("checkpoint A")
	slog.Info("hello")
	slog.Warn("Be careful")
	slog.Error("oops", net.ErrClosed, "status", 500)

	slog.LogAttrs(slog.ErrorLevel, "oops", slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	slog.Info("Test group",
		slog.Group("request",
			slog.String("url", "/foo"),
			slog.String("referer", "/bar")))
}

func main() {
	slogAsText()
	slogAsJSON()
}
