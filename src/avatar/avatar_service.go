package avatar

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

type GenerateUserAvatarer interface {
	GenerateUserAvatar(userId string) (*string, error)
}

type AvatarService struct {
	GenerateUserAvatarer
}

type DefaultAvatarService struct {
	TasURL string
	ApiKey string
}

type GenerateUserAvatarResponse struct {
	AvatarURL string `json:"avatarURL"`
}

func (svc *DefaultAvatarService) GenerateUserAvatar(userId string) (*string, error) {
	client := resty.New()

	resp, err := client.R().SetHeader("X-API-KEY", svc.ApiKey).Post(svc.TasURL)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var response GenerateUserAvatarResponse

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &response.AvatarURL, nil
}
