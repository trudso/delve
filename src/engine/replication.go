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

// SetFromAuthority will only set the value if the value is from the authority
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

func (r *Replicatable[T]) Get() *T {
	return &r.value
}

func (r Replicatable[T]) IsChanged() bool {
	return r.originalValue != r.value
}

func (r Replicatable[T]) ShouldReplicate() bool {
	return r.shouldReplicate
}

func NewReplicatable[T comparable](value T, shouldReplicate bool, isAuthority AuthorityFunc) Replicatable[T] {
	return Replicatable[T]{
		originalValue: value,
		value: value,
		shouldReplicate: shouldReplicate,
		IsAuthority: isAuthority,
	}
}
