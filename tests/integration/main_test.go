// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"testing"

	"github.com/btwiuse/etcd/v3/pkg/testutil"
	"github.com/btwiuse/etcd/v3/server/embed"
)

func TestMain(m *testing.M) {
	embed.SetupGrpcLoggingForTest(true)
	testutil.MustTestMainWithLeakDetection(m)
}
