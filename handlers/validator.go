package handlers

import (
	"fmt"

	"github.com/rajikaimal/tch-admin/utils"
)

func (q *CommonStudentReqQuery) ValidateCommonStudentReqBody() error {
	if len(q.Teacher) == 0 {
		return fmt.Errorf("Teacher's Email is required")
	}

	for _, email := range q.Teacher {
		if utils.IsValidEmail(email) == false {
			return fmt.Errorf("Invalid Teacher's Email")
		}
	}

	return nil
}

func (b *RegisterReqBody) ValidateRegisterReqBody() error {
	if b.Teacher == "" {
		return fmt.Errorf("Teacher's Email is required")
	}

	if utils.IsValidEmail(b.Teacher) == false {
		return fmt.Errorf("Invalid Teacher's Email")
	}

	return nil
}

func (b *SuspendReqBody) ValidateSuspendReqBody() error {
	if b.Email == "" {
		return fmt.Errorf("Email is required")
	}

	if utils.IsValidEmail(b.Email) == false {
		return fmt.Errorf("Invalid email")
	}

	return nil
}

func (b *RetrieveNotificationReqBody) ValidateRetrieveNotificationReqBody() error {
	if b.Teacher == "" {
		return fmt.Errorf("Teacher's Email is required")
	}

	if b.Notification == "" {
		return fmt.Errorf("Notifications Text is required")
	}

	return nil
}
