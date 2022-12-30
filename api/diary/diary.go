package diary

import (
	"modoo-diary-api/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetDiaryList(c *fiber.Ctx) error {
	type ResponseBody struct {
		DiaryId   int       `json:"diaryId"`
		Contents  string    `json:"contents"`
		CreatedAt time.Time `json:"createdAt"`
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
