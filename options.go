package timberjack

import (
	option2 "github.com/wangyougui/timberjack.v2/option"
	"time"
)

const (
	optkeyClock            = "clock"
	optkeyMaxAge           = "max-age"
	optKeyMaxSize          = "max-size"
	optkeyRotationTime     = "rotation-time"
	optkeyMaxBackups       = "max-backup"
	optKeyCompress         = "compress"
	optKeyCompression      = "compression"
	optKeyLocalTime        = "local-time"
	optKeyRotationInterval = "rotation-interval"
	optKeyRotateAtMinutes  = "rotate-at-minutes"
	optKeyBackupTimeFormat = "backup-time-format"
)

type option struct {
	name  string
	value interface{}
}

func (o *option) Name() string       { return o.name }
func (o *option) Value() interface{} { return o.value }

const optSpecificationSet = `opt-specification-set`

// WithSpecification allows you to specify a custom specification set
func WithSpecificationSet(ds SpecificationSet) Option {
	return &option{
		name:  optSpecificationSet,
		value: ds,
	}
}

type optSpecificationPair struct {
	name     byte
	appender Appender
}

const optSpecification = `opt-specification`

// WithSpecification allows you to create a new specification set on the fly,
// to be used only for that invocation.
func WithSpecification(b byte, a Appender) Option {
	return &option{
		name: optSpecification,
		value: &optSpecificationPair{
			name:     b,
			appender: a,
		},
	}
}

// WithMilliseconds is similar to WithSpecification, and specifies that
// the Strftime object should interpret the pattern `%b` (where b
// is the byte that you specify as the argument)
// as the zero-padded, 3 letter milliseconds of the time.
func WithMilliseconds(b byte) Option {
	return WithSpecification(b, Milliseconds())
}

// WithMicroseconds is similar to WithSpecification, and specifies that
// the Strftime object should interpret the pattern `%b` (where b
// is the byte that you specify as the argument)
// as the zero-padded, 3 letter microseconds of the time.
func WithMicroseconds(b byte) Option {
	return WithSpecification(b, Microseconds())
}

// WithUnixSeconds is similar to WithSpecification, and specifies that
// the Strftime object should interpret the pattern `%b` (where b
// is the byte that you specify as the argument)
// as the unix timestamp in seconds
func WithUnixSeconds(b byte) Option {
	return WithSpecification(b, UnixSeconds())
}

// WithClock creates a new Option that sets a clock
// that the RotateLogs object will use to determine
// the current time.
//
// By default rotatelogs.Local, which returns the
// current time in the local time zone, is used. If you
// would rather use UTC, use rotatelogs.UTC as the argument
// to this option, and pass it to the constructor.
func WithClock(c Clock) Option {
	return option2.NewOption(optkeyClock, c)
}

// WithLocation creates a new Option that sets up a
// "Clock" interface that the RotateLogs object will use
// to determine the current time.
//
// This optin works by always returning the in the given
// location.
func WithLocation(loc *time.Location) Option {
	return option2.NewOption(optkeyClock, clockFn(func() time.Time {
		return time.Now().In(loc)
	}))
}

// WithMaxAge creates a new Option that sets the
// max age of a log file before it gets purged from
// the file system.
func WithMaxAge(d int) Option {
	return option2.NewOption(optkeyMaxAge, d)
}

// WithRotationTime creates a new Option that sets the
// time between rotation.
func WithRotationTime(d time.Duration) Option {
	return option2.NewOption(optkeyRotationTime, d)
}

// WithMaxBackups creates a new Option that sets the
// log file size between rotation.
func WithMaxBackups(s int) Option {
	return option2.NewOption(optkeyMaxBackups, s)
}

// WithCompress creates a new Option that sets the
// number of files should be kept before it gets
// purged from the file system.
func WithCompress(n bool) Option {
	return option2.NewOption(optKeyCompress, n)
}

func WithMaxSize(maxSize int) Option {
	return option2.NewOption(optKeyMaxSize, maxSize)
}

func WithCompression(compression string) Option {
	return option2.NewOption(optKeyCompression, compression)
}

func WithLocalTime(localTime bool) Option {
	return option2.NewOption(optKeyLocalTime, localTime)
}

func WithRotationInterval(d time.Duration) Option {
	return option2.NewOption(optKeyRotationInterval, d)
}

func WithRotateAtMinutes(d []int) Option {
	return option2.NewOption(optKeyRotateAtMinutes, d)
}

func WithBackupTimeFormat(format string) Option {
	return option2.NewOption(optKeyBackupTimeFormat, format)
}
