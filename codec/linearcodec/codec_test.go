// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package linearcodec

import (
	"testing"

	"github.com/dim4egster/qmallgo/codec"
)

func TestVectors(t *testing.T) {
	for _, test := range codec.Tests {
		c := NewDefault()
		test(c, t)
	}
}

func TestMultipleTags(t *testing.T) {
	for _, test := range codec.MultipleTagsTests {
		c := New([]string{"tag1", "tag2"}, defaultMaxSliceLength)
		test(c, t)
	}
}
