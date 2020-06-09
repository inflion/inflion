// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package flow

import "testing"

func TestFlow(t *testing.T) {
	f := NewOpsFlow(MockRecipeReader{})
	_, err := f.Run()
	if err != nil {
		t.Error(err)
	}
}
