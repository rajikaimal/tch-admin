package utils

import "testing"

func TestContains(t *testing.T) {
	slice := []string{"student1@gmail.com", "student2@gmail.com", "student3@gmail.com"}
	value := "student2@gmail.com"

	if !Contains(slice, value) {
		t.Errorf("Expected slice to contain value '%s', it didn't", value)
	}

	value = "student4@gmail.com"
	if Contains(slice, value) {
		t.Errorf("Expected slice not to contain value '%s', it did", value)
	}
}
