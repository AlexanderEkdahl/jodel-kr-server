package main

import "testing"

func TestRepoCreateMessage(t *testing.T) {
	m := Message{
		Message: "Alex was here",
	}
	m = RepoCreateMessage(m)
	t.Logf("Created Message: %v", m)

	if m.ID == 0 {
		t.Errorf("Returned message should have a ID higher than zero")
	}
}
