// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logutil

import (
	"time"

	pb "github.com/youtube/vitess/go/vt/proto/logutil"
)

// This file contains a few functions to help with proto3.  proto3
// will eventually support timestamps, at which point we'll retire
// this.

// ProtoToTime converts a pb.Time to a time.Time.
//
// A nil pointer is like the empty timestamp.
func ProtoToTime(ts *pb.Time) time.Time {
	if ts == nil {
		// treat nil like the empty Timestamp
		return time.Unix(0, 0).UTC()
	}
	return time.Unix(ts.Seconds, int64(ts.Nanoseconds)).UTC()
}

// TimeToProto converts the time.Time to a pb.Time.
func TimeToProto(t time.Time) *pb.Time {
	seconds := t.Unix()
	nanos := int64(t.Sub(time.Unix(seconds, 0)))
	return &pb.Time{
		Seconds:     seconds,
		Nanoseconds: int32(nanos),
	}
}