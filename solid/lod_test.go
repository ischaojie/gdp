package solid

import "testing"

func TestLOD(t *testing.T) {
	var girls []Girl
	for i := 0; i < 20; i++ {
		girls = append(girls, Girl{})
	}

	teacher := Teacher{}
	groupLeader := NewGroupLeader(girls)
	teacher.Command(groupLeader)
}
