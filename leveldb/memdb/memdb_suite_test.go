package memdb

import (
	"testing"

	"github.com/df-mc/golevelb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
