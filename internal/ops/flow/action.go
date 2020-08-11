package flow

// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

type Action struct {
	Type   string
	Params map[string]string
}

type Actions []Action

type ActionResult struct {
	Action     Action
	Outputs    map[string]string
	ExitStatus bool
}

type ActionResults []ActionResult

func (a ActionResults) isSuccess() bool {
	for _, ar := range a {
		if ar.ExitStatus == false {
			return false
		}
	}
	return true
}

func (a ActionResults) getExitMessage() string {
	if a.isSuccess() {
		return "success"
	}

	return "fail"
}
