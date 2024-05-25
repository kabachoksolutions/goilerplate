package loggingbp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"path"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	var layer string
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		if a.Key == "layer" {
			layer = a.Value.String()
		} else {
			fields[a.Key] = a.Value.Any()
		}

		return true
	})

	if layer == "" {
		layer = "app"
	}

	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format(fmt.Sprintf("[%s]", time.DateTime))
	msg := color.CyanString(r.Message)

	_, file, line, _ := runtime.Caller(4)
	filePath := fmt.Sprintf("%s:%d", file, line)

	message := fmt.Sprintf("SOURCE %s LAYER %s", color.GreenString(path.Base(filePath)), color.RedString(layer))

	if len(b) > 3 {
		message += " " + string(b)
	}

	h.l.Println(timeStr, level, msg, message)

	return nil
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}
	return h
}
