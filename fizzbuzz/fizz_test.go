package fizzbuzz

import (
	"strconv"
	"testing"

	"github.com/magiconair/properties/assert"
)


func Test_fizz(t *testing.T) {
	assert.Equal(t,"fizz",Convert(3))
}

func Test_buzz(t *testing.T) {
	assert.Equal(t,"buzz",Convert(5))
}

func Test_other(t *testing.T) {
	assert.Equal(t,"1",Convert(1))
}

func Test_fizzbuzz(t *testing.T) {
	assert.Equal(t,"fizzbuzz",Convert(15))
}

func Test_IT_fizzbuzz(t *testing.T) {
	var wants [100]string
	for i := 0; i < 100; i++  {
		if ((i+1) % 15 == 0) {
			wants[i] = "fizzbuzz"
		} else if ((i+1) % 3 == 0) {
			wants[i] = "fizz"
		} else if ((i+1) % 5 == 0) {
			wants[i] = "buzz"
		} else {
			wants[i] = strconv.Itoa(i+1)
		}
	}

	for i,want := range wants {
		assert.Equal(t,want,Convert(i+1))
	}
}