package kakao

import (
	"fmt"
	"golang-5252/database"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type KakaoRequest struct {
	User struct {
		Id string `json:"id"`
	} `json:"user"`
	Action struct {
		Params map[string]interface{}
	}
}
type KakaoResponse struct {
	Version  string `json:"version"`
	Template struct {
		Outputs []struct {
			SimpleText struct {
				Text string `json:"text"`
			} `json:"simpleText"`
		} `json:"outputs"`
	} `json:"template"`
}

func makeSimpleText(text string) (response KakaoResponse) {
	response.Version = "2.0"
	response.Template = struct {
		Outputs []struct {
			SimpleText struct {
				Text string "json:\"text\""
			} "json:\"simpleText\""
		} "json:\"outputs\""
	}{}
	response.Template.Outputs = append(response.Template.Outputs, struct {
		SimpleText struct {
			Text string "json:\"text\""
		} "json:\"simpleText\""
	}{})
	response.Template.Outputs[0].SimpleText = struct {
		Text string "json:\"text\""
	}{text}
	return
}
func PostKakaoHandler(c *fiber.Ctx) error {
	method := c.GetReqHeaders()["Method"]
	switch method {
	case "readDiary":
		return postReadDiary(c)
	case "writeDiary":
		return postWriteDiary(c)
	case "readMyDiary":
		return postReadMyDiary(c)
	case "requestToken":
		return postRequestToken(c)
	case "login":
		return postLogin(c)
	case "logout":
		return postLogout(c)
	default:
		return postFailMethod(c)
	}
}
func postReadDiary(c *fiber.Ctx) error {
	diaryList, err := database.SelectDiaryListTop10()
	if err != nil {
		return postFailMethod(c)
	}

	result := ""
	for _, diary := range diaryList {
		result += diary.DiaryContent + "\n"
	}

	return c.Type("application/json").JSON(makeSimpleText(result))
}
func postReadMyDiary(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		return postFailMethod(c)
	}
	userId, err := database.SelectUserIdByKakaoId(kakaoRequest.User.Id)
	if err != nil {
		return postFailMethod(c)
	}
	diaryList, err := database.SelectDiaryListTop10ByUserId(userId)
	if err != nil {
		return postFailMethod(c)
	}

	result := ""
	for _, diary := range diaryList {
		result += diary.DiaryContent + "\n"
	}

	return c.Type("application/json").JSON(makeSimpleText(result))
}
func postWriteDiary(c *fiber.Ctx) (err error) {
	return c.Type("application/json").JSON(makeSimpleText("write"))
}
func postLogin(c *fiber.Ctx) (err error) {
	return c.Type("application/json").JSON(makeSimpleText("login"))
}
func postLogout(c *fiber.Ctx) (err error) {
	return c.Type("application/json").JSON(makeSimpleText("logout"))
}
func postRequestToken(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		return postFailMethod(c)
	}
	email, ok := kakaoRequest.Action.Params["email"].(string)
	if !ok {
		return postFailMethod(c)
	}
	err = database.InsertUser(email)
	if err != nil {
		return postFailMethod(c)
	}
	createdToken := randSeq(10)
	err = database.UpdateToken(email, createdToken)
	if err != nil {
		return postFailMethod(c)
	}
	// TODO: Email 보내기
	return c.Type("application/json").JSON(makeSimpleText("이메일 요청이 완료됐습니다"))
}
func postFailMethod(c *fiber.Ctx) (err error) {
	str := string(c.Body())
	fmt.Println(str)
	return c.Type("application/json").JSON(makeSimpleText("fail"))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
