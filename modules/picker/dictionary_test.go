// © 2019 Imhotep Software LLC. All rights reserved.

package picker_test

import (
	"testing"

	"github.com/gopherland/labs/picker"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	l, err := picker.LoadDefault("musicians")

	assert.Nil(t, err)
	assert.Equal(t, 5, len(l))
}

func TestPicker(t *testing.T) {
	l, err := picker.LoadDefault("musicians")
	w := picker.Pick(l)

	assert.Nil(t, err)
	assert.NotEqual(t, "", w)
}

func TestLoadMissing(t *testing.T) {
	_, err := picker.LoadDefault("actors")
	assert.EqualError(t, err, "unable to load dictionary `assets/actors.txt")
}

func TestLoadCustomFail(t *testing.T) {
	_, err := picker.Load("zorg", "actors")
	assert.EqualError(t, err, "unable to load dictionary `zorg/actors.txt")
}
