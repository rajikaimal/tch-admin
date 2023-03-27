package handlers

import (
	"errors"
	"testing"
)

func TestValidateCommonStudentReqBody(t *testing.T) {
	// valid teacher email(s)
	q1 := CommonStudentReqQuery{Teacher: []string{"teacher1@example.com"}}
	err := q1.ValidateCommonStudentReqBody()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// empty teacher email(s)
	q2 := CommonStudentReqQuery{Teacher: []string{""}}
	err = q2.ValidateCommonStudentReqBody()
	expectedErr := errors.New("Invalid Teacher's Email")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Unexpected error. Got: %v, expected: %v", err, expectedErr)
	}

	// invalid teacher email(s)
	q3 := CommonStudentReqQuery{Teacher: []string{"invalid_email"}}
	err = q3.ValidateCommonStudentReqBody()
	expectedErr = errors.New("Invalid Teacher's Email")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Unexpected error. Got: %v, expected: %v", err, expectedErr)
	}
}

func TestValidateRegisterReqBody(t *testing.T) {
	// Valid Register Request Body
	reqBody1 := RegisterReqBody{
		Teacher:  "teacher@example.com",
		Students: []string{"student1@example.com", "student2@example.com"},
	}
	err := reqBody1.ValidateRegisterReqBody()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Empty Teacher Email in Register Request Body
	reqBody2 := RegisterReqBody{
		Teacher:  "",
		Students: []string{"student1@example.com"},
	}
	err = reqBody2.ValidateRegisterReqBody()
	if err == nil || err.Error() != "Teacher's Email is required" {
		t.Errorf("Expected error: Teacher's Email is required")
	}

	// Empty Students Email List in Register Request Body
	reqBody3 := RegisterReqBody{
		Teacher:  "teacher@example.com",
		Students: []string{},
	}
	err = reqBody3.ValidateRegisterReqBody()
	if err == nil || err.Error() != "Student Email(s) are required" {
		t.Errorf("Expected error: Student Email(s) are required")
	}

	// Invalid Teacher Email in Register Request Body
	reqBody4 := RegisterReqBody{
		Teacher:  "invalid_email",
		Students: []string{"student1@example.com", "student2@example.com"},
	}
	err = reqBody4.ValidateRegisterReqBody()
	if err == nil || err.Error() != "Invalid Teacher's Email" {
		t.Errorf("Expected error: Invalid Teacher's Email")
	}
}

func TestValidateSuspendReqBody(t *testing.T) {
	// valid email
	reqBody1 := &SuspendReqBody{Email: "user@example.com"}

	err := reqBody1.ValidateSuspendReqBody()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// invalid email
	reqBody2 := &SuspendReqBody{Email: "invalid email"}

	err = reqBody2.ValidateSuspendReqBody()

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Error() != "Invalid Email" {
		t.Errorf("Expected error message: 'Invalid Email', but got: '%s'", err.Error())
	}

	// empty email
	reqBody3 := &SuspendReqBody{Email: ""}

	err = reqBody3.ValidateSuspendReqBody()

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Error() != "Email is required" {
		t.Errorf("Expected error message: 'Email is required', but got: '%s'", err.Error())
	}
}

func TestValidateRetrieveNotificationReqBody(t *testing.T) {
	// empty notification text
	reqBody1 := &RetrieveNotificationReqBody{Teacher: "teacher@example.com", Notification: ""}

	err := reqBody1.ValidateRetrieveNotificationReqBody()

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Error() != "Notifications Text is required" {
		t.Errorf("Expected error message: 'Notifications Text is required', but got: '%s'", err.Error())
	}

	// empty Teacher
	reqBody2 := &RetrieveNotificationReqBody{Teacher: "", Notification: "notification text"}

	err = reqBody2.ValidateRetrieveNotificationReqBody()

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Error() != "Teacher's Email is required" {
		t.Errorf("Expected error message: 'Teacher's Email is required', but got: '%s'", err.Error())
	}

	reqBody3 := &RetrieveNotificationReqBody{Teacher: "teacher@example.com", Notification: "Hello @student1@gmail.com"}

	err = reqBody3.ValidateRetrieveNotificationReqBody()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
