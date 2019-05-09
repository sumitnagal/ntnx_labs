// © 2019 Imhotep Software LLC. All rights reserved.

package wc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWords(t *testing.T) {
	uu := []struct {
		path  string
		count int
		err   error
	}{
		{"assets/4moons.txt", 2503, nil},
	}

	for _, u := range uu {
		wc, err := Words(u.path)
		assert.Equal(t, u.err, err)
		assert.Equal(t, u.count, wc)
	}
}

func TestCountWords(t *testing.T) {
	uu := []struct {
		line  string
		count int
	}{
		{"", 0},
		{"Hello World", 2},
		{"Hello World and then some", 5},
		{"😁", 0},
		{"😁 And there you go", 4},
		{multiLine, 110},
		{`It's "Boff" an' it's "Wham" un'erstan'?`, 6},
	}

	for _, u := range uu {
		assert.Equal(t, u.count, countWords(u.line))
	}
}

var multiLine = `
Oh, I'm Popeye the Sailor Man,
I'm Popeye the Sailor Man.
I'm strong to the finich
Cause I eats me spinach.
I'm Popeye the Sailor Man.

I'm one tough Gazookus
Which hates all Palookas
Wot ain't on the up and square.
I biffs 'em and buffs 'em
And always out roughs 'em
But none of 'em gets nowhere.

If anyone dares to risk my "Fisk",
It's "Boff" an' it's "Wham" un'erstan'?
So keep "Good Be-hav-or"
That's your one life saver
With Popeye the Sailor Man.

I'm Popeye the Sailor Man,
I'm Popeye the Sailor Man.
I'm strong to the finich
Cause I eats me spinach.
I'm Popeye the Sailor Man.
`
