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

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email   string
		wantRes bool
	}{
		{
			email:   "john.snow@example.com",
			wantRes: true,
		},
		{
			email:   "jane.snow@example.co.uk",
			wantRes: true,
		},
		{
			email:   "jane.snow!!!!*(*(&@example.co.uk",
			wantRes: false,
		},
		{
			email:   "invalid_email",
			wantRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			gotRes := IsValidEmail(tc.email)
			if gotRes != tc.wantRes {
				t.Errorf("IsValidEmail(%v) = %v, want %v", tc.email, gotRes, tc.wantRes)
			}
		})
	}
}
