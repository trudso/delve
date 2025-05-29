package engine

import (
	"fmt"
	"os"
)

type ReplicationPrimitive[T comparable] struct {
	id            string
	originalValue T
	value         *T

	// Replication if true, the authority should replicate the value to the connections
	shouldReplicate bool

	// Replication if true, the value is replicated to connections, and never overwritten
	IsAuthority func() bool
}

func (r ReplicationPrimitive[T]) GetId() string {
	return r.id
}

// SetFromAuthority will only set the value if the value is from the authority
func (r *ReplicationPrimitive[T]) SetFromAuthority(value T) {
	if r.IsAuthority() {
		return
	}
	r.originalValue = value
	*r.value = value
}

func (r ReplicationPrimitive[T]) IsChanged() bool {
	return r.originalValue != *r.value
}

func (r ReplicationPrimitive[T]) ShouldReplicate() bool {
	return r.shouldReplicate
}

func (r ReplicationPrimitive[T]) BuildChangeSet() map[string]any {
	if !r.IsChanged() || !r.ShouldReplicate() {
		return nil
	}
	return map[string]any{r.id: *r.value}
}

func (r ReplicationPrimitive[T]) BuildSnapshot() map[string]any {
	return map[string]any{r.id: *r.value}
}

func (r ReplicationPrimitive[T]) ApplyDataSet(dataSet map[string]any) {
	if d, found := dataSet[r.id]; found {
		*r.value = d.(T)
	} else {
		fmt.Fprintln(os.Stderr, "No primitive data found for:", r.id)
	}
}

func (r *ReplicationPrimitive[T]) ResetToChanged() {
	r.originalValue = *r.value
}

func NewReplicationPrimitive[T comparable](id string, value *T, shouldReplicate bool, isAuthority func() bool) *ReplicationPrimitive[T] {
	return &ReplicationPrimitive[T]{
		id:              id,
		originalValue:   *value,
		value:           value,
		shouldReplicate: shouldReplicate,
		IsAuthority:     isAuthority,
	}
}
