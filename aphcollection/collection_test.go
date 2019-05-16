package aphcollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	s := []string{"foo", "bar", "jar", "mar", "lar", "sar"}
	n := Remove(s, []string{"foo", "bar"}...)
	assert := assert.New(t)
	assert.Len(n, 4)
}

func TestIndex(t *testing.T) {
	s := []string{"foo", "bar", "jar", "mar", "lar", "sar"}
	assert := assert.New(t)
	assert.Less(Index(s, "lora"), 0, "should return less than zero")
	assert.Equal(Index(s, "jar"), 2, "should return index 2")
	assert.Equal(Index(s, "sar"), 5, "should return index 5")
}

func TestContains(t *testing.T) {
	s := []string{"foo", "bar", "jar", "mar", "lar", "sar"}
	assert := assert.New(t)
	assert.True(Contains(s, "lar"), true, "element lar should be present")
	assert.True(Contains(s, "bar"), true, "element bar should be present")
	assert.False(Contains(s, "nota"), true, "element nota should not be present")
}
