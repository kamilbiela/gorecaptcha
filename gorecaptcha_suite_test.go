package gorecaptcha

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGorecaptcha(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gorecaptcha Suite")
}

type MockRecaptcha struct {
	recaptcha
}

var _ = Describe("Gorecaptcha", func() {

	Context("parseErrorLine", func() {
		It("Should parse status line to constant", func() {
			Expect(parseErrorLine("invalid-site-private-key")).To(Equal(ErrInvalidSitePrivateKey))
			Expect(parseErrorLine("somethinig not supported")).To(Equal(ErrUnknown))
		})
	})

	Context("recaptcha", func() {
		It("NewRecaptcha", func() {
			key := "some-private-key"
			r := NewRecaptcha(key)

			Expect(r.privateKey).To(Equal(key))
		})

		It("Verify: return false and error on error response", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				txt, _ := ioutil.ReadFile("test/resp_err_invalid_key.txt")
				fmt.Fprintln(w, string(txt))
			}))
			defer ts.Close()

			r := NewRecaptcha("key")
			r.verifyURL = ts.URL
			resp, _ := r.Verify("", "", "")
			Expect(resp.err).To(Equal(ErrInvalidSitePrivateKey))
			Expect(resp.status).To(Equal(false))
		})

		It("Verify: return true and no error on success response", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				txt, _ := ioutil.ReadFile("test/response_success.txt")
				fmt.Fprintln(w, string(txt))
			}))
			defer ts.Close()

			r := NewRecaptcha("key")
			r.verifyURL = ts.URL
			resp, _ := r.Verify("", "", "")
			Expect(resp.err).To(BeNil())
			Expect(resp.status).To(Equal(true))
		})
	})
})
