package service_admin

import (
	"library/db"
	"library/model/model_admin"
	"regexp"
)

// 长度 6-18 字母开头 只能包含数字 字母下划线
const PasswordRegex = "^[a-zA-Z]\\w{5,17}$"

//4-16位字母,数字,汉字,下划线
const PhoneRegex = "^1(3\\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$"

//const NameRegex = "^([\u4e00-\u9fa5]{2,4})|([A-Za-z0-9_]{4,16})|([a-zA-Z0-9_\u4e00-\u9fa5]{3,16})$"
//const EmailRegex = "^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"

// QueryAdminByAdminID 查询管理员,根据管理员ID查询管理员信息，仅供内部使用，不作为api暴露
func QueryAdminByAdminID(AdminID int) model_admin.Admin {
	var admin model_admin.Admin
	if err := db.MySQL.First(&admin, AdminID).Error; err != nil {
		return model_admin.Admin{AdminID: -1}
	}
	return admin
}

func ValidAdmin(admin *model_admin.Admin) bool {
	if len(admin.Phone) == 0 || len(admin.Password) == 0 {
		return false
	}
	if m, _ := regexp.MatchString(PhoneRegex, admin.Phone); !m {
		return false
	}
	if m, _ := regexp.MatchString(PasswordRegex, admin.Password); !m {
		return false
	}
	return true
}
