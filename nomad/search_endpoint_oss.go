// +build !pro,!ent

package nomad

import (
	"fmt"

	memdb "github.com/hashicorp/go-memdb"
	"github.com/hashicorp/nomad/nomad/state"
	"github.com/hashicorp/nomad/nomad/structs"
)

// getEnterpriseMatch is a no-op in oss since there are no enterprise objects.
func getEnterpriseMatch(match interface{}) (id string, ok bool) {
	return "", false
}

// getEnterpriseResourceIter is used to retrieve an iterator over an enterprise
// only table.
func getEnterpriseResourceIter(context structs.Context, namespace, prefix string, ws memdb.WatchSet, state *state.StateStore) (memdb.ResultIterator, error) {
	// If we have made it here there
	return nil, fmt.Errorf("context must be one of %v or 'all' for all contexts; got %q", ossContexts, context)
}
