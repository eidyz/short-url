package helpers

import (
	"github.com/teris-io/shortid"
)

// ShortID - short id generator
var ShortID *shortid.Shortid

// InitShortID ---
func InitShortID() {
	generator, err := shortid.New(1, shortid.DefaultABC, 406026)
	if err != nil {
		panic(err)
	}
	
	ShortID = generator
	return
}
