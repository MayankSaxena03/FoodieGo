package models

var KeyPhoneLoginOTP = "otp:phone:?"
var KeyEmailLoginOTP = "otp:email:?"

type LoginSignupBody struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	OTP      string `json:"otp,omitempty" bson:"otp,omitempty"`
}
