package datastore2

import (
	"github.com/jbenet/go-ipfs/Godeps/_workspace/src/github.com/jbenet/go-datastore"
)

// ClaimThreadSafe claims that a Datastore is threadsafe, even when
// it's type does not guarantee this. Use carefully.
type ClaimThreadSafe struct {
	datastore.Datastore
}

var _ datastore.ThreadSafeDatastore = ClaimThreadSafe{}

func (ClaimThreadSafe) IsThreadSafe() {}
