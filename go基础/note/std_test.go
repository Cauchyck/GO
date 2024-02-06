package note

import (
	"errors"
	"testing"
)

func TestIsNotNagative(log *testing.T) {
	err := errors.New("Is Negatibve")
	if IsNotNagative(0) {
		log.Log("OK")
	} else {
		log.Error(err)
	}
	if IsNotNagative(1) {
		log.Log("OK")
	} else {
		log.Error(err)
	}
}
