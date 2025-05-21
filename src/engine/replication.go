package engine

import (
	"fmt"
	"os"
)

// interfaces
type AuthorityFunc func() bool

type Replicatable interface {
	ResetToChanged()
	IsChanged() bool
	//NewInstance() Replicatable
	BuildChangeSet() map[string]any
	BuildSnapshot() map[string]any
	ApplyDataSet(dataSet map[string]any)
}

// ---------------- primitive ------------- //
type ReplicationPrimitive[T comparable] struct {
	id            string
	originalValue T
	value         *T

	// Replication if true, the authority should replicate the value to the connections
	shouldReplicate bool

	// Replication if true, the value is replicated to connections, and never overwritten
	IsAuthority AuthorityFunc
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

func NewReplicationPrimitive[T comparable](id string, value *T, shouldReplicate bool, isAuthority AuthorityFunc) *ReplicationPrimitive[T] {
	return &ReplicationPrimitive[T]{
		id:              id,
		originalValue:   *value,
		value:           value,
		shouldReplicate: shouldReplicate,
		IsAuthority:     isAuthority,
	}
}

// --------------- Collection -------------- //
type ReplicationCollection struct {
	id       string
	elements []Replicatable
}

func (r ReplicationCollection) IsChanged() bool {
	for _, element := range r.elements {
		if element.IsChanged() {
			return true
		}
	}
	return false
}

func (r ReplicationCollection) ResetToChanged() {
	for _, element := range r.elements {
		element.ResetToChanged()
	}
}

func (r ReplicationCollection) BuildChangeSet() map[string]any {
	if !r.IsChanged() {
		return nil
	}
	data := map[string]any{}
	for _, element := range r.elements {
		var ds = BuildChangeSet(element)
		for k, v := range ds {
			data[k] = v
		}
	}

	return map[string]any{r.id: data}
}

func (r ReplicationCollection) BuildSnapshot() map[string]any {
	data := map[string]any{}
	for _, element := range r.elements {
		var ds = BuildSnapshot(element)
		for k, v := range ds {
			data[k] = v
		}
	}

	return map[string]any{r.id: data}
}

// TODO[mt]: how to work sub collections ?
func (r ReplicationCollection) ApplyDataSet(dataSet map[string]any) {
	if value, found := dataSet[r.id]; found {
		ds := value.(map[string]any)
		for key, element := range r.elements {
			fmt.Println("Applying", key, "to", element)
			ApplyDataSet(element, ds)
		}
	} else {
		fmt.Fprintln( os.Stderr, "No data found for:", r.id)	
	}
}

func (r ReplicationCollection) AddCollection(other ReplicationCollection) {
	r.elements = append(r.elements, other.elements...)
}

func (r *ReplicationCollection) AddElement(element Replicatable) {
	r.elements = append(r.elements, element)
}

func NewReplicationCollection(id string, elements []Replicatable) ReplicationCollection {
	return ReplicationCollection{
		id:       id,
		elements: elements,
	}
}

func BuildChangeSet(replicatable Replicatable) map[string]any {
	return replicatable.BuildChangeSet()
}

func BuildSnapshot(replicatable Replicatable) map[string]any {
	return replicatable.BuildSnapshot()
}

func ResetToChanged(replicatable Replicatable) {
	replicatable.ResetToChanged()
}

func ApplyDataSet(replicatable Replicatable, dataSet map[string]any) {
	replicatable.ApplyDataSet(dataSet)
}
