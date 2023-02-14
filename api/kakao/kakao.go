package kakao

import (
	"fmt"
	"log"
	"modoo-diary-api/database"
	"modoo-diary-api/pkg/discord"
	"modoo-diary-api/pkg/random"
	"modoo-diary-api/pkg/slack"
	smtp "modoo-diary-api/pkg/smtp"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type KakaoRequest struct {
	UserRequest struct {
		User struct {
			Id string `json:"id"`
		} `json:"user"`
	} `json:"userRequest"`
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
	case "changeNickname":
		return postChangeNickname(c)
	default:
		return postFailMethod(c, "method")
	}
}
func postReadDiary(c *fiber.Ctx) error {
	diaryList, err := database.SelectDiaryListTop10()
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "db")
	}

	result := ""
	for _, diary := range diaryList {
		result += fmt.Sprintf("%s  %s\n", diary.Nickname, diary.Diary.CreatedAt.Format("2006-01-02 15:04"))
		result += diary.DiaryContent + "\n"
		result += "==========\n"
	}

	return c.Type("application/json").JSON(makeSimpleText(result))
}
func postReadMyDiary(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "body")
	}
	user, err := database.SelectUserByKakaoId(kakaoRequest.UserRequest.User.Id)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "login")
	}
	diaryList, err := database.SelectDiaryListTop10ByUserId(user.User.UserId)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "db")
	}

	result := ""
	for _, diary := range diaryList {
		result += fmt.Sprintf("%s\n", diary.CreatedAt.Format("2006-01-02 15:04"))
		result += diary.DiaryContent + "\n"
		result += "==========\n"
	}

	return c.Type("application/json").JSON(makeSimpleText(result))
}
func postWriteDiary(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "body")
	}
	user, err := database.SelectUserByKakaoId(kakaoRequest.UserRequest.User.Id)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "로그인이 필요합니다")
	}
	text, ok := kakaoRequest.Action.Params["text"].(string)
	if !ok {
		return postFailMethod(c, "text param")
	}
	if text == "취소" {
		return c.Type("application/json").JSON(makeSimpleText("취소 됐습니다"))

	}
	err = database.InsertDiary(user.User.UserId, text)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "db insert")
	}
	discord.SendWebHook(user.Nickname + ": " + text)
	slack.SendWebHook(user.Nickname + ": " + text)
	return c.Type("application/json").JSON(makeSimpleText("일기 작성이 완료됐습니다"))
}
func postLogin(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "body")
	}

	_, err = database.SelectUserByKakaoId(kakaoRequest.UserRequest.User.Id)
	if err != gorm.ErrRecordNotFound {
		return postFailMethod(c, "이미 로그인 돼 있습니다")
	}
	email, ok := kakaoRequest.Action.Params["email"].(string)
	if !ok {
		return postFailMethod(c, "param email")
	}
	validEmail, _ := regexp.Compile(`^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]`)
	if !validEmail.MatchString(email) {
		return postFailMethod(c, "param email 형식이 맞지 않습니다")
	}
	token, ok := kakaoRequest.Action.Params["auth_token"].(string)
	if !ok {
		return postFailMethod(c, "param token")
	}
	user, err := database.SelectLoginToken(email, token)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "login token")
	}
	err = database.InsertKakaoAuth(kakaoRequest.UserRequest.User.Id, user.UserId)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "insert kakao token")
	}
	database.UpdateToken(email, null.NewString("", false), null.NewTime(time.Now(), false))
	return c.Type("application/json").JSON(makeSimpleText("로그인 성공"))
}
func postLogout(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "body")
	}
	err = database.DeleteKakaoAuth(kakaoRequest.UserRequest.User.Id)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "delete")
	}
	return c.Type("application/json").JSON(makeSimpleText("logout"))
}
func postRequestToken(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "body")
	}
	email, ok := kakaoRequest.Action.Params["email"].(string)
	if !ok {
		return postFailMethod(c, "param email")
	}
	validEmail, _ := regexp.Compile(`^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]`)
	if !validEmail.MatchString(email) {
		return postFailMethod(c, "param email 형식이 맞지 않습니다")
	}
	err = database.InsertUser(email)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "db insert")
	}
	createdToken := random.RandSeq(10)
	err = database.UpdateToken(email, null.NewString(createdToken, true), null.NewTime(time.Now().Add(30*time.Minute), true))
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "db update")
	}

	smtp.SendMail(email, createdToken)
	return c.Type("application/json").JSON(makeSimpleText("이메일로 토큰 정보를 보내드렸습니다.\n확인 후 입력해주세요."))
}
func postChangeNickname(c *fiber.Ctx) (err error) {
	var kakaoRequest KakaoRequest
	err = c.BodyParser(&kakaoRequest)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "body")
	}
	nickname, ok := kakaoRequest.Action.Params["nickname"].(string)
	if !ok {
		return postFailMethod(c, "param nickname")
	}
	user, err := database.SelectUserByKakaoId(kakaoRequest.UserRequest.User.Id)
	if err != nil {
		return postFailMethod(c, "로그인이 필요합니다")
	}
	err = database.UpdateNickname(user.User.UserId, nickname)
	if err != nil {
		log.Println(err)
		return postFailMethod(c, "db update")
	}

	return c.Type("application/json").JSON(makeSimpleText("nickname이 변경됐습니다"))
}
func postFailMethod(c *fiber.Ctx, message string) (err error) {
	str := string(c.Body())
	fmt.Println(str)
	return c.Type("application/json").JSON(makeSimpleText("fail: " + message))
}
