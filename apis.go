package simple

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/techpartners-asia/simple-go/utils"
)

// simple
var (
	SimpleGetToken = utils.API{
		Url:    "/mbank-auth-main-service/token",
		Method: http.MethodPost,
	}
	SimpleCreateInvoice = utils.API{
		Url:    "/mbank-integration-gateway-service/integration/addInvoice/createInvoice",
		Method: http.MethodPost,
	}
	SimpleSendInvoiceToNumber = utils.API{
		Url:    "/mbank-integration-gateway-service/integration/addInvoice/assignInvoiceMobile",
		Method: http.MethodPost,
	}
	SimpleGetInvoice = utils.API{
		Url:    "/mbank-integration-gateway-service/integration/checkInvoice/merchant",
		Method: http.MethodGet,
	}
	SimpleChangeInvoiceStatus = utils.API{
		Url:    "mbank-integration-gateway-service/integration/invoiceDelivered",
		Method: http.MethodPost,
	}
)

func (p *simple) httpRequest(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
	authObj, authErr := p.authSimple()
	if authErr != nil {
		err = authErr
		return
	}
	p.loginObject = &authObj

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	baseUrl := p.baseUrl + api.Url + urlExt

	req, _ := http.NewRequest(api.Method, baseUrl, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+p.loginObject.Data.AccessToken)

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}

	return
}

// Authsimple [Login to simple]
func (s *simple) authSimple() (authRes SimpleGetTokenResponse, err error) {
	if s.loginObject != nil {
		now := time.Now().Unix()
		if now < *s.expire_in {
			authRes = *s.loginObject
			err = nil
			return
		}
	}
	url := s.baseUrl + SimpleGetToken.Url
	reqBody := SimpleGetTokenBody{
		GrantType: "client_credentials",
		DeviceID:  "test",
	}

	reqBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(SimpleGetToken.Method, url, bytes.NewReader(reqBytes))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.userName, s.password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return authRes, fmt.Errorf("%s-Simple auth response: %s", time.Now().Format(utils.TimeFormatYYYYMMDDHHMMSS), res.Status)
	}

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &authRes)

	defer res.Body.Close()
	expireD := time.Now().Unix() + int64(authRes.Data.ExpiresIn)
	s.expire_in = &expireD
	return authRes, nil
}
