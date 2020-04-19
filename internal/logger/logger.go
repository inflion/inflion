// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package logger

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
	WarnLevel
	FatalLevel
)

type Configuration struct {
	Level Level
}

type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Error(args ...interface{})

	Fatal(args ...interface{})

	DebugWith(msg string, keysAndValues ...interface{})

	InfoWith(msg string, keysAndValues ...interface{})

	ErrorWith(msg string, keysAndValues ...interface{})

	FatalWith(msg string, keysAndValues ...interface{})
}
