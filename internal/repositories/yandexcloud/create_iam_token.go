package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type iamTokenResponse struct {
	IamToken  string    `json:"iamToken"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type iamTokenRequest struct {
	YandexPassportOauthToken string `json:"yandexPassportOauthToken"`
}

func (y *YandexCloud) fillIamToken() error {
	reqStruct := iamTokenRequest{
		YandexPassportOauthToken: y.config.YandexCloud.OauthToken,
	}

	bodyBytes, err := y.sendHttpRequest(
		context.Background(),
		http.MethodPost,
		"https://iam.api.cloud.yandex.net/iam/v1/tokens",
		reqStruct,
	)

	if err != nil {
		return err
	}

	resParsed := iamTokenResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return err
	}

	y.iamToken = resParsed.IamToken

	return nil
}
