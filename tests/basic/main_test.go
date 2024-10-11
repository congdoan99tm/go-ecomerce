package basic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	//"github.com/stretchr/testify"
)

func TestAddOne(t *testing.T) {
	//var (
	//	input  = 1
	//	output = 2
	//)
	//actual := AddOne(1)
	//if actual != output {
	//	t.Errorf("AddOne(%d) , input %d, expected: %d, actual: %d", input, input, output, actual)
	//}
	assert.Nil(t, nil, nil)
	assert.Equal(t, AddOne(1), 3, fmt.Sprintf("Expected AddOne(1) to return %d, but got %d", 2, AddOne(1)))
}
