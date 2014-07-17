package gorecaptcha

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type recaptcha struct {
	privateKey string
	verifyURL  string
}

func (re *recaptcha) makeVerifyRequest(
	remoteIP string,
	challenge string,
	response string,
) (string, error) {
	resp, err := http.PostForm(re.verifyURL, url.Values{
		"privatekey": {re.privateKey},
		"remoteip":   {remoteIP},
		"challenge":  {challenge},
		"response":   {response},
	})

	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// NewRecaptcha is reCaptcha client constructor
func NewRecaptcha(privateKey string) *recaptcha {
	return &recaptcha{
		privateKey,
		"http://www.google.com/recaptcha/api/verify",
	}
}

// verify data via captcha server
// https://developers.google.com/recaptcha/docs/verify
func (re *recaptcha) Verify(
	remoteIP string,
	challenge string,
	response string,
) (recaptchaResponse, error) {

	body, err := re.makeVerifyRequest(remoteIP, challenge, response)
	resp := recaptchaResponse{}

	if err != nil {
		return resp, err
	}

	lines := strings.Split(string(body), "\n")

	if lines[0] == "true" {
		resp.status = true
	} else {
		resp.status = false
		resp.err = parseErrorLine(lines[1])
	}

	return resp, nil
}
