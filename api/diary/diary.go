package diary

import (
	"modoo-diary-api/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetDiaryList(c *fiber.Ctx) (err error) {
	type RequestQuery struct {
		Page int `query:"page"`
	}
	type ResponseBody struct {
		DiaryId   int       `json:"diaryId"`
		Author    string    `json:"author"`
		Contents  string    `json:"contents"`
		CreatedAt time.Time `json:"createdAt"`
	}
	var responseBody []ResponseBody

	var requestQuery RequestQuery
	c.QueryParser(&requestQuery)
	if requestQuery.Page == 0 {
		requestQuery.Page = 1
	}

	diaryList, err := database.SelectDiaryListByPaging(requestQuery.Page, 10)
	if err != nil {
		return c.SendStatus(500)
	}
	for _, diary := range diaryList {
		var temp ResponseBody
		temp.DiaryId = diary.DiaryId
		temp.Author = diary.Nickname
		temp.Contents = diary.DiaryContent
		temp.CreatedAt = diary.Diary.CreatedAt
		responseBody = append(responseBody, temp)
	}
	return c.JSON(responseBody)
}
