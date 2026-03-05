package logx

import (
	"log/slog"
	"time"

	"github.com/av1ppp/timex"
)

type Kind = slog.Kind

const (
	KindAny Kind = iota
	KindBool
	KindDuration
	KindFloat64
	KindInt64
	KindString
	KindTime
	KindUint64
	KindGroup
	KindLogValuer
)

type Attr = slog.Attr

// String returns an Attr for a string value.
var String = slog.String

// Int8 converts an int8 to an int64 and returns
// an Attr with that value.
func Int8(key string, value int8) Attr {
	return Int64(key, int64(value))
}

// Int16 converts an int16 to an int64 and returns
// an Attr with that value.
func Int16(key string, value int16) Attr {
	return Int64(key, int64(value))
}

// Int32 converts an int32 to an int64 and returns
// an Attr with that value.
func Int32(key string, value int32) Attr {
	return Int64(key, int64(value))
}

// Int64 returns an Attr for an int64.
var Int64 = slog.Int64

// Int converts an int to an int64 and returns
// an Attr with that value.
var Int = slog.Int

// Uint8 converts an uint8 to an uint64 and returns
// an Attr with that value.
func Uint8(key string, value uint8) Attr {
	return Uint64(key, uint64(value))
}

// Uint16 converts an uint16 to an uint64 and returns
// an Attr with that value.
func Uint16(key string, value uint16) Attr {
	return Uint64(key, uint64(value))
}

// Uint32 converts an uint32 to an uint64 and returns
// an Attr with that value.
func Uint32(key string, value uint32) Attr {
	return Uint64(key, uint64(value))
}

// Uint64 returns an Attr for a uint64.
var Uint64 = slog.Uint64

// Uint converts an uint to an uint64 and returns
// an Attr with that value.
func Uint(key string, value uint) Attr {
	return Uint64(key, uint64(value))
}

// Float32 converts an float32 to an float64 and returns
// an Attr with that value.
func Float32(key string, value float32) Attr {
	return Float64(key, float64(value))
}

// Float64 returns an Attr for a floating-point number.
var Float64 = slog.Float64

// Bool returns an Attr for a bool.
var Bool = slog.Bool

// Time returns an Attr for a [time.Time].
// It discards the monotonic portion.
func Time(key string, value time.Time) Attr {
	return Attr{
		Key:   key,
		Value: slog.StringValue(value.Format(time.RFC3339)),
	}
}

// Duration returns an Attr for a [time.Duration].
var Duration = slog.Duration

func Durationx(key string, value timex.Duration) Attr {
	return slog.String(key, value.String())
}

var Group = slog.Group

// Any returns an Attr for the supplied value.
// See [AnyValue] for how values are treated.
var Any = slog.Any

func Module(name string) slog.Attr {
	return slog.String("module", name)
}

func App(name string) slog.Attr {
	return slog.String("app", name)
}

// Cause returns an Attr that represents the cause of the error.
// The Attr contains a string key "cause" with the error's message as its value.
func Cause(err error) slog.Attr {
	if err == nil {
		return slog.String("cause", "<nil>")
	}
	return slog.String("cause", err.Error())
}

var GroupValue = slog.GroupValue
