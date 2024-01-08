//nolint:gochecknoglobals,gochecknoinits
package log

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file

		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]

				break
			}
		}

		file = short

		return file + ":" + strconv.Itoa(line)
	}

	desiredTimeFormat := "01-02 15:04:05"

	zerolog.TimeFieldFormat = desiredTimeFormat
	zerolog.CallerFieldName = "loc"
	zerolog.LevelFieldName = "lvl"
	zerolog.TimestampFieldName = "at"
	zerolog.MessageFieldName = "msg"

	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: desiredTimeFormat}).
		Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return logger.Fatal()
}
