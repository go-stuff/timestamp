package timestamp

// This file implements operations on google.protobuf.Timestamp.

import (
	"time"

	"github.com/go-stuff/timestamp/pb"
)

// Timestamp converts a google.protobuf.Timestamp proto to a time.Time.
// It returns an error if the argument is invalid.
//
// Unlike most Go functions, if Timestamp returns an error, the first return value
// is not the zero time.Time. Instead, it is the value obtained from the
// time.Unix function when passed the contents of the Timestamp, in the UTC
// locale. This may or may not be a meaningful time; many invalid Timestamps
// do map to valid time.Times.
//
// A nil Timestamp returns an error. The first return value in that case is
// undefined.
func Timestamp(ts *pb.Timestamp) time.Time {
	if ts == nil {
		return time.Unix(0, 0)
	}
	return time.Unix(ts.Time, 0)
}
