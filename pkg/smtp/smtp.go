package smtp

import (
	"net/smtp"
	"os"
)

func SendMail(email string, token string) {
	id := os.Getenv("EMAIL_ID")
	pw := os.Getenv("EMAIL_PW")
	auth := smtp.PlainAuth("", id, pw, "smtp.gmail.com")
	from := id
	to := []string{email}

	// 메시지 작성
	headerSubject := "Subject: [제목] 모두의 일기 인증\r\n"
	headerBlank := "\r\n"
	body := "요청하신 로그인 token은 다음과 같습니다.\r\n" + token
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
}
