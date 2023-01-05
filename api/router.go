package api

import (
	"modoo-diary-api/api/diary"
	"modoo-diary-api/api/kakao"
	"modoo-diary-api/api/login"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Post("/api/kakao", kakao.PostKakaoHandler)
	app.Get("/api/diaries", diary.GetDiaryList)
	app.Post("/api/diaries", diary.PostDiary)
	app.Post("/api/login/token", login.PostRequestToken)
	app.Post("/api/login", login.PostLogin)
}
