package rpg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterRefresh(t *testing.T) {
	c := NewCharacter("Test", Attributes{Strength: 8, Agility: 9, Intelligence: 10})
	modifier := AttributesModification{
		Modification: Attributes{
			Strength:     3,
			Agility:      4,
			Intelligence: 5,
		},
	}

	// run twice to ensure no double increments
	c.Refresh([]Modifier{modifier})
	c.Refresh([]Modifier{modifier})

	assert.Equal(t, 11, c.Attributes.Current.Strength)
	assert.Equal(t, 13, c.Attributes.Current.Agility)
	assert.Equal(t, 15, c.Attributes.Current.Intelligence)
}
