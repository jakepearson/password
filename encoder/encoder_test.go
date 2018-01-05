package encoder

import (
	"testing"
)

var hashTests = []struct {
	password string
	expected string
}{
	{"angryMonkey", "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="},
	{"", "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg=="},
}

func TestHashAndEncode(t *testing.T) {
	for _, test := range hashTests {
		actual := HashAndEncode(test.password)
		if actual != test.expected {
			t.Errorf("'%v' not equal to '%v' for '%s'", actual, test.expected, test.password)
		}
	}
}