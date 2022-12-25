package api

import (
	"golang-5252/api/kakao"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Post("/api/kakao", kakao.PostKakaoHandler)
}
