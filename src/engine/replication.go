package engine

type AuthorityFunc func() bool

// ---------------- primitive ------------- //
type ReplicationPrimitive[T comparable] struct {
	id string
	originalValue T
	value *T

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

func (r ReplicationPrimitive[T]) BuildDataSet(fullSnapshot bool) map[string]any {
	if !fullSnapshot && (!r.IsChanged() || !r.ShouldReplicate()) {
		return nil
	}
	return map[string]any{r.id: *r.value}	
} 

func (r *ReplicationPrimitive[T]) ResetToChanged() {
	r.originalValue = *r.value	
}

func NewReplicationPrimitive[T comparable](id string, value *T, shouldReplicate bool, isAuthority AuthorityFunc) *ReplicationPrimitive[T] {
	return &ReplicationPrimitive[T]{
		id: id,
		originalValue: *value,
		value: value,
		shouldReplicate: shouldReplicate,
		IsAuthority: isAuthority,
	}
}

// --------------- Collection -------------- //
type Replicatable interface {
	ResetToChanged()
	IsChanged() bool
	BuildDataSet(fullSnapshot bool) map[string]any
}

type ReplicationCollection struct {
	id string
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

func (r ReplicationCollection) BuildDataSet(fullSnapshot bool) map[string]any {
	if !fullSnapshot && !r.IsChanged() {
		return nil
	}
	data := map[string]any{}
	for _, element := range r.elements {
		var ds = buildDataSet( element, fullSnapshot )
		for k, v := range ds {
			data[k] = v	
		}
	}

	return map[string]any{r.id: data}
}

func (r ReplicationCollection) AddCollection( other ReplicationCollection ) {
	r.elements = append(r.elements, other.elements...)
}

func (r *ReplicationCollection) AddElement( element Replicatable ) {
	r.elements = append(r.elements, element)
}

func NewReplicationCollection(id string, elements []Replicatable) ReplicationCollection {
	return ReplicationCollection{
		id: id,
		elements: elements,
	}
}

func buildDataSet( replicatable Replicatable, fullSnapshot bool) map[string]any {
	return replicatable.BuildDataSet(fullSnapshot)
}

func BuildChangeSet( replicatable Replicatable ) map[string]any {
	return replicatable.BuildDataSet(false)
}

func BuildSnapshot( replicatable Replicatable ) map[string]any {
	return replicatable.BuildDataSet(true)
}

func ResetToChanged( replicatable Replicatable) {
	replicatable.ResetToChanged()
}
