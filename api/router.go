package api

import (
	"modoo-diary-api/api/diary"
	"modoo-diary-api/api/kakao"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Post("/api/kakao", kakao.PostKakaoHandler)
	app.Get("/api/diaries", diary.GetDiaryList)
}
