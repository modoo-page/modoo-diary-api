package diary

import (
	"modoo-diary-api/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetDiaryList(c *fiber.Ctx) error {
	type ResponseBody struct {
		DiaryId   int
		Contents  string
		CreatedAt time.Time
	}
	var responseBody []ResponseBody
	diaryList, err := database.SelectDiaryListTop10()
	if err != nil {
		return c.SendStatus(500)
	}
	for _, diary := range diaryList {
		var temp ResponseBody
		temp.DiaryId = diary.DiaryId
		temp.Contents = diary.DiaryContent
		temp.CreatedAt = diary.Diary.CreatedAt
		responseBody = append(responseBody, temp)
	}
	return c.JSON(responseBody)
}
