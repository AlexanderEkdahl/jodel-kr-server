package main

import "testing"

func TestRepoCreateMessage(t *testing.T) {
	m := Message{
		Message: "Alex was here",
	}
	m, err := RepoCreateMessage(m)
	if err != nil {
		t.Errorf("There should not be an error when creating a message.")
	}

	t.Logf("Created Message: %v", m)

	if m.ID == 0 {
		t.Errorf("Returned message should have a ID higher than zero")
	}
}
