package infrastracture

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"unicode"

	"gopkg.in/gomail.v2"
)

func IsValidEmail(email string) bool {

	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}


// func SendActivationEmail(email, token string) error {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "bereket.meng@gmail.com")
// 	m.SetHeader("To", email)
// 	m.SetHeader("Subject", "Account Activation")


// 	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://localhost:8080/users/verify-email?email=%s&token=%s\">here</a> to activate your account.", email, token))


// 	d := gomail.NewDialer("smtp.gmail.com", 587, "bereket.meng@gmail.com", "xjbs vduu hkjd lqlf")

	
// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}
// 	return nil
// }

func SendActivationEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "bereket.meng@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Account Activation")

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
					color: #333;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					background-color: #ffffff;
					padding: 20px;
					border-radius: 8px;
					box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
				}
				.header {
					text-align: center;
					padding-bottom: 20px;
				}
				.header h1 {
					color: #4CAF50;
				}
				.body-content {
					font-size: 16px;
					line-height: 1.6;
				}
				.otp {
					font-size: 24px;
					font-weight: bold;
					color: #4CAF50;
					text-align: center;
					margin-top: 20px;
					margin-bottom: 20px;
				}
				.footer {
					text-align: center;
					font-size: 14px;
					color: #888;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Account Activation</h1>
				</div>
				<div class="body-content">
					<p>Dear User,</p>
					<p>Thank you for signing up! To complete your registration, please use the following OTP code to activate your account:</p>
					<div class="otp">%s</div>
					<p>If you did not request this code, please ignore this email.</p>
				</div>
				<div class="footer">
					<p>Best regards,<br>Felagi</p>
				</div>
			</div>
		</body>
		</html>
	`, token)

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "bereket.meng@gmail.com", "xjbs vduu hkjd lqlf")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}


func GenerateOTP() (string, error) {
	digits := "0123456789"
	var otp string
	for i := 0; i < 6; i++ {
		randomByte := make([]byte, 1)
		_, err := rand.Read(randomByte)
		if err != nil {
			return "", err
		}
		randomByte[0] = randomByte[0] % byte(len(digits))
		otp += string(digits[randomByte[0]])
	}
	return otp, nil
}

func GenerateActivationToken() (string, error) {
	// Create a 32-byte random token
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	// Convert the token to a hex string
	return hex.EncodeToString(token), nil
}

func GenerateDeviceFingerprint(ip, userAgent string) string {
    data := ip + userAgent
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

func SendResetLink(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "bereket.meng@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Link")


	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://127.0.0.1:8080/users/reset-password/%s\">here</a> to reset your password.", token))


	d := gomail.NewDialer("smtp.gmail.com", 587, "bereket.meng@gmail.com", "xjbs vduu hkjd lqlf")

	
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}