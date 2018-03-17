package bandname

import (
	"github.com/dustinkirkland/golang-petname"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Bandname() string {
	return petname.Generate(2, "-")
}
