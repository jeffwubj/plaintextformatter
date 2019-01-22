package plaintextformatter

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// PlainTextFormatter is similar to logrus.TextFormatter but display text directly
// which is suitable for console disply
type PlainTextFormatter struct {
	TimestampFormat string
	ShowLevel       bool
	ShowTime        bool
}

// Format the log entry. Implements logrus.Formatter.
func (f *PlainTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	prefixFieldClashes(data)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339Nano
	}

	data["time"] = entry.Time.Format(timestampFormat)
	data["message"] = entry.Message
	data["severity"] = entry.Level.String()
	output := fmt.Sprintf("%s", data["message"])
	if f.ShowTime {
		output = fmt.Sprintf("%s %s", data["time"], output)
	}
	if f.ShowLevel {
		output = fmt.Sprintf("%s %s", data["severity"], output)
	}

	serialized := []byte(output)

	return append(serialized, '\n'), nil
}

func prefixFieldClashes(data logrus.Fields) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
	}

	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
	}

	if l, ok := data["level"]; ok {
		data["fields.level"] = l
	}
}
