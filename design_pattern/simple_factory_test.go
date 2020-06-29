package dp

import "testing"

func TestSimpleFactory(t *testing.T) {
	animal := NewAnimal(DogType)
	animal.Say()
}
