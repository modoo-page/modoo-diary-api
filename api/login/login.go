package login

import (
	"log"
	"modoo-diary-api/database"
	"modoo-diary-api/pkg/random"
	smtp "modoo-diary-api/pkg/smtp"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

func PostRequestToken(c *fiber.Ctx) (err error) {
	type RequestBody struct {
		Email string `json:"email"`
	}
	var requestBody RequestBody
	err = c.BodyParser(&requestBody)
	if err != nil {
		return c.SendStatus(400)
	}
	validEmail, _ := regexp.Compile(`^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]`)
	if !validEmail.MatchString(requestBody.Email) {
		return c.SendStatus(400)
	}
	err = database.InsertUser(requestBody.Email)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	createdToken := random.RandSeq(10)
	err = database.UpdateToken(requestBody.Email, null.NewString(createdToken, true), null.NewTime(time.Now().Add(30*time.Minute), true))
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	go func(email string, token string) {
		smtp.SendMail(email, token)
	}(requestBody.Email, createdToken)
	return c.SendStatus(200)
}

func PostLogin(c *fiber.Ctx) (err error) {
	type RequestBody struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}
	type ResponseBody struct {
		UserToken string `json:"userToken"`
	}
	var requestBody RequestBody
	var responseBody ResponseBody
	err = c.BodyParser(&requestBody)
	if err != nil {
		return c.SendStatus(400)
	}
	validEmail, _ := regexp.Compile(`^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]`)
	if !validEmail.MatchString(requestBody.Email) {
		return c.SendStatus(400)
	}
	user, err := database.SelectLoginToken(requestBody.Email, requestBody.Token)
	if err != nil {
		log.Println(err)
		return c.SendStatus(403)
	}
	responseBody.UserToken = user.UserToken
	return c.JSON(responseBody)
}
