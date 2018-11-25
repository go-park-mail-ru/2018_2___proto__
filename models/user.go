package models

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

/*type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"id,omitempty valid:"-"`
	Nickname             string   `protobuf:"bytes,2,opt,name=Nickname,json=nickname,proto3" json:"nickname" valid:"required,login"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,json=password,proto3" json:"password,omitempty" valid:"required,pass"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,json=email,proto3" json:"email,omitempty" valid:"required,email"`
	Fullname             string   `protobuf:"bytes,5,opt,name=Fullname,json=fullname,proto3" json:"fullname,omitempty" valid:"name"`
	Avatar               string   `protobuf:"bytes,6,opt,name=Avatar,json=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}*/

func init() {
	// custom validator for password
	govalidator.CustomTypeTagMap.Set("pass", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		password, ok := i.(string)
		if !ok {
			return false
		}

		// very strong password check
		// at least 1 uppercase letter, 1 lowercase, 1 numerical and 1 special
		regexp := regexp.MustCompile(
			"^[[[:alnum:]]*[[:graph:]]]*.{8,20}$")
		correct := regexp.MatchString(password)
		if !correct {
			return false
		}
		return true
	}))
	// "?=.*[A-Z].)(?=.*[!@#$&*_])(?=.*[0-9])(?=.*[a-z]).{8}$"
	// "^[A-Z]+[0-9]+[a-zA-Z0-9_`!@#$%^&.?()-=+]*.{8,20}$"
	// custom validator for nickname
	govalidator.CustomTypeTagMap.Set("login", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		nickname, ok := i.(string)
		if !ok {
			return false
		}

		// only latin letters are allowed
		regexp := regexp.MustCompile(
			"^[a-zA-Z0-9_.]*.{3,20}$")
		correct := regexp.MatchString(nickname)
		if !correct {
			return false
		}
		if len(nickname) < 3 || len(nickname) > 16 {
			return false
		}
		return true
	}))

	// custom validator for fullname
	govalidator.CustomTypeTagMap.Set("name", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		fullname, ok := i.(string)
		if !ok {
			return false
		}

		// only latin letters are allowed
		regexp := regexp.MustCompile(
			"^[a-zA-Z]+$")
		correct := regexp.MatchString(fullname)
		if !correct {
			return false
		}

		if len(fullname) < 2 || len(fullname) > 32 {
			return false
		}
		return true
	}))
}
