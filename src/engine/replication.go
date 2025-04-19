package engine

type AuthorityFunc func() bool

type Replicatable[T comparable] struct {
	originalValue T
	value T

	// Replication if true, the authority should replicate the value to the connections
	shouldReplicate bool

	// Replication if true, the value is replicated to connections, and never overwritten
	IsAuthority AuthorityFunc
}

func (r *Replicatable[T]) SetFromAuthority(value T) {
	if r.IsAuthority() {
		return
	}
	r.originalValue = value
	r.value = value
}

func (r *Replicatable[T]) Set(value T) {
	if !r.IsAuthority() {
		return		
	}	
	r.value = value
} 

func (r Replicatable[T]) Get() T {
	return r.value
}

func (r Replicatable[T]) IsChanged() bool {
	return r.originalValue != r.value
}

func (r Replicatable[T]) ShouldReplicate() bool {
	return r.shouldReplicate
}
