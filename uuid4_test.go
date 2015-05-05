package uuid4

import "testing"

func TestBytes(t *testing.T) {
	for i := 0; i < 100000; i++ {
		u := NewUUID()
		if u.DashString()[14] != '4' {
			t.Error("Byte 14 is not '4':", u.DashString())
		}
		b19 := u.DashString()[19]
		if b19 != '8' && b19 != '9' && b19 != 'a' && b19 != 'b' {
			t.Error("Byte 19 is not one of '89ab':", u.DashString())
		}
		t.Log(u.DashString())
	}
}
