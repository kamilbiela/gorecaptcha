package gorecaptcha

import (
	"errors"
)

//
type recaptchaResponse struct {
	Status bool
	Err    error
}

var (
	// ErrInvalidSitePrivateKey error response type
	ErrInvalidSitePrivateKey = errors.New("Invalid site private key")
	// ErrInvalidRequestCookie error response type
	ErrInvalidRequestCookie = errors.New("Invalid request coookie")
	// ErrIncorrectCaptchaSol error response type
	ErrIncorrectCaptchaSol = errors.New("Incorrect captcha sol")
	// ErrCaptchaTimeout error response type
	ErrCaptchaTimeout = errors.New("Captcha timeout")
	// ErrRecaptchaNotReachable error response type
	ErrRecaptchaNotReachable = errors.New("Recaptcha not reachable")
	// ErrUnknown error response type
	ErrUnknown = errors.New("Recaptcha not supported error")
)

func parseErrorLine(line string) error {
	switch line {
	case "invalid-site-private-key":
		return ErrInvalidSitePrivateKey
	case "invalid-request-cookie":
		return ErrInvalidRequestCookie
	case "incorrect-captcha-sol":
		return ErrIncorrectCaptchaSol
	case "captcha-timeout":
		return ErrCaptchaTimeout
	case "recaptcha-not-reachable":
		return ErrRecaptchaNotReachable
	default:
		return ErrUnknown
	}
}
