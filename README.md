Easy to use, tested reCaptcha library. Allows to check send data against reCaptcha service. 
There are other libs to do this, but my returns errors n RecaptchaResponse object so You know what went wrong ;).

# Usage of lib

[![Build Status](https://travis-ci.org/kamilbiela/gorecaptcha.svg?branch=master)](https://travis-ci.org/kamilbiela/gorecaptcha)

```
recaptcha := gorecaptcha.NewRecaptcha("recaptcha private key")

captchaResponse, err := recaptcha.Verify(
	ip,
	challenge_from_request_form,
	response_from_request_form,
)

if err != nil {
	// error (http request error for example)
	fmt.Println(err)
}

fmt.Println(captchaResponse.Status) // true or false
fmt.Println(captchaResponse.Err) // actual reason if false
```
