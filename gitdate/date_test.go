package gitdate

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func Test_formatDate(t *testing.T) {
	_time, err := formatDate("8:18")
	assert.Empty(t, err)
	assert.Equal(t, _time.Hour(), 8)
	assert.Equal(t, _time.Minute(), 18)

	_time, err = formatDate("8:18:22")
	assert.Empty(t, err)
	assert.Equal(t, _time.Hour(), 8)
	assert.Equal(t, _time.Minute(), 18)
	assert.Equal(t, _time.Second(), 22)

	_time, err = formatDate("x8:18:22  ")
	assert.Empty(t, err)
	assert.Equal(t, _time.Hour(), 8)
	assert.Equal(t, _time.Minute(), 18)
	assert.Equal(t, _time.Second(), 22)
}
