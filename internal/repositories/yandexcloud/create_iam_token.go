package yandexcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
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
	y.logger.Info("Создаем iam token для доступа к yandex cloud")

	reqStruct := iamTokenRequest{
		YandexPassportOauthToken: y.config.YandexCloud.OauthToken,
	}

	body, err := json.Marshal(reqStruct)
	if err != nil {
		y.logger.Error(err.Error())
		return err
	}

	newCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(newCtx, http.MethodPost, "https://iam.api.cloud.yandex.net/iam/v1/tokens", bytes.NewBuffer(body))
	if err != nil {
		y.logger.Error(err.Error())
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		y.logger.Error(err.Error())
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := errors.New("не удалось создать iam token")
		y.logger.Error(err.Error())
		return err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		y.logger.Error(err.Error())
		return err
	}

	resParsed := iamTokenResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return err
	}

	y.iamToken = resParsed.IamToken
	y.logger.Info("Iam token успешно создан!")

	return nil
}
