package adyen2

import (
	"log"
	"testing"
)

func TestFingerprint_String(t *testing.T) {
	fp := NewDF()
	log.Printf("fp: %v", fp.String())
}
