package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type EmailVerification struct {
	Id        int
	Email     string    `orm:"size(100)"`
	Code      string    `orm:"size(100)"`
	ExpiresAt time.Time `orm:"type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func SetVerifyCode(email string, verifyCode string) error {
	o := orm.NewOrm()
	expiresAt := time.Now().Add(time.Hour)
	// Check if record with given email exists
	var emailVerification EmailVerification
	err := o.QueryTable("email_verification").Filter("email", email).One(&emailVerification)
	if err == nil {
		// Record exists, update Code and ExpiresAt fields
		emailVerification.Code = verifyCode
		emailVerification.ExpiresAt = expiresAt
		_, err = o.Update(&emailVerification)
		return err
	} else {
		// Record does not exist, create new record
		emailVerification = EmailVerification{
			Email:     email,
			Code:      verifyCode,
			ExpiresAt: expiresAt,
		}
		_, err = o.Insert(&emailVerification)
		return err
	}
}

func CheckVerifyCode(email string, verifyCode string) (string, error) {
	// Retrieve the EmailVerification row for the given email
	emailVerification := EmailVerification{Email: email}
	if err := orm.NewOrm().Read(&emailVerification, "Email"); err != nil {
		return "", err
	}

	// Check if the verification code is correct and not expired
	if emailVerification.Code != verifyCode {
		return "invalid verification code", nil
	}

	if time.Now().Before(emailVerification.ExpiresAt) {
		return "", nil
	} else {
		return "verification code expired", nil
	}
}
