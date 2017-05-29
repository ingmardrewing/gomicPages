package db

import "testing"

func TestQueryRowSingleValue(t *testing.T) {
	Initialize()

	var actual int
	db.QueryRow("SELECT MAX(pageNumber) FROM pages").Scan(&actual)
	expected := 87
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestQueryRowMultipleValues(t *testing.T) {
	Initialize()

	var (
		actual_id    int
		actual_title string
	)
	db.QueryRow("SELECT id, title FROM pages LIMIT 1").
		Scan(
			&actual_id,
			&actual_title)
	expected_id := 1
	expected_title := "#1 A Step in the dark"
	if actual_id != expected_id {
		t.Errorf("Expected id to be %d, but got %d",
			expected_id, actual_id)
	}
	if actual_title != expected_title {
		t.Errorf("Expected title to be %s, but got %s",
			expected_title, actual_title)
	}
}
