package bandname

import (
	"fmt"
	"math/rand"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Bandname() string {
	return petname.Generate(2, "-")
}

func Bandemail() string {
	return fmt.Sprintf("%s@example.com", Bandname())
}
