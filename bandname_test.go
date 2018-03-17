package bandname

import (
	"fmt"
	"testing"
)

func TestBandname(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Bandname())
	}
}
