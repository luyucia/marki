package app

import "testing"

func TestNameHandle(t *testing.T) {

	name,suffix := handleName("aaa.md")

	t.Log(name)
	t.Log(suffix)
}
