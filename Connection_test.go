package picqer

import (
	"testing"
)

func TestGetPicklist(t *testing.T) {
	Picklists, err := GetPicklists()

	if err != nil {
		t.Fatal(err)
	}

	if len(Picklists) > 100 {
		t.Fatal("List is greather than 100 should be 100")
	}

	if len(Picklists) < 100 {
		t.Fatal("List is less than 100 should be 100")
	}

}
