package db

import (
	"strings"
	"time"

	dbmodels "github.com/gangjun06/book-server/models/db"
	"github.com/gangjun06/book-server/utils"
	"github.com/jinzhu/gorm"
)

func CreateUser(email, name, password string) error {
	verifyCode := utils.CreateRandomString(8)

	hashedPassword := utils.HashAndSalt(password)
	result := utils.GetDB().Create(&dbmodels.User{Name: name, Email: email, Password: hashedPassword, VerifyCode: verifyCode, JoinAt: time.Now()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUserByID(id int) (*dbmodels.User, error) {
	data := &dbmodels.User{}
	result := utils.GetDB().Where("id = ?", id).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrItemNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func FindUserByEmail(email string) (*dbmodels.User, error) {
	data := &dbmodels.User{}
	loweremail := strings.ToLower(email)
	result := utils.GetDB().Where("email = ?", loweremail).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrItemNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func UpdatePassword(id int, password string) error {
	result := utils.GetDB().Model(&dbmodels.User{}).Where("id = ?", id).Update("password", utils.HashAndSalt(password))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckVerify(code string) (bool, error) {
	data := &dbmodels.User{}
	if result := utils.GetDB().Where("verify_code = ?", code).First(&data); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	if data.Verified {
		return false, ErrUserAlreadyVerified
	}

	if result := utils.GetDB().Model(&dbmodels.User{}).Where("verify_code = ?", code).Update("verified", true); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func SetPasswordVerifyCode(email, code string) error {
	data := &dbmodels.User{}
	result := utils.GetDB().Where("email = ?", email).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return ErrItemNotFound
	} else if result.Error != nil {
		return result.Error
	}

	currentTime := time.Now()
	if result := utils.GetDB().Model(&dbmodels.User{}).Where("email = ?", email).Update(&dbmodels.User{VerifyCodePassword: code, VerifyCodePasswordTime: &currentTime}); result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckPasswordVerify(email, code string) (bool, error) {
	if code == "" {
		return false, nil
	}
	data := &dbmodels.User{}
	result := utils.GetDB().Where("email = ? and verify_code_password = ?", email, code).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return false, ErrItemNotFound
	} else if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
