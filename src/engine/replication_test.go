package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthority(t *testing.T) {
	authority := func() bool {
		return true
	}

	sut := Replicatable[int]{
		originalValue: 0,
		value:         0,
		shouldReplicate: true,
		IsAuthority:    authority,
	}

	// test attempting setting when already authority
	sut.SetFromAuthority(1)
	assert.Equal(t, sut.Get(), 0)
	assert.False(t, sut.IsChanged())
	assert.True(t, sut.ShouldReplicate())

	sut.Set(1)
	assert.Equal(t, sut.Get(), 1)
	assert.True(t, sut.IsChanged())
	assert.True(t, sut.ShouldReplicate())
}
