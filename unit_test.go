package ethgo

import (
	"fmt"
	"math/big"
	"testing"
)

func TestValue(t *testing.T) {
	println(fmt.Sprintf("%d", NewValue(big.NewFloat(0.001), Ada)))
}
