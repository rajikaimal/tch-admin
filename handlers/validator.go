package handlers

import (
	"errors"

	"github.com/rajikaimal/tch-admin/utils"
)

func (q *CommonStudentReqQuery) ValidateCommonStudentReqBody() error {
	if len(q.Teacher) == 0 {
		return errors.New("Teacher's Email is required")
	}

	for _, email := range q.Teacher {
		if utils.IsValidEmail(email) == false {
			return errors.New("Invalid Teacher's Email")
		}
	}

	return nil
}

func (b *RegisterReqBody) ValidateRegisterReqBody() error {
	if b.Teacher == "" {
		return errors.New("Teacher's Email is required")
	}

	if len(b.Students) == 0 {
		return errors.New("Student Email(s) are required")
	}

	if utils.IsValidEmail(b.Teacher) == false {
		return errors.New("Invalid Teacher's Email")
	}

	return nil
}

func (b *SuspendReqBody) ValidateSuspendReqBody() error {
	if b.Email == "" {
		return errors.New("Email is required")
	}

	if utils.IsValidEmail(b.Email) == false {
		return errors.New("Invalid Email")
	}

	return nil
}

func (b *RetrieveNotificationReqBody) ValidateRetrieveNotificationReqBody() error {
	if b.Teacher == "" {
		return errors.New("Teacher's Email is required")
	}

	if b.Notification == "" {
		return errors.New("Notifications Text is required")
	}

	return nil
}
