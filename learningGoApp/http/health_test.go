package http_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	uri    string
	svr    *httptest.Server
	router *mux.Router
)

var _ = Describe("Health", func() {
	// TODO: implement tests for health endpoint?

	BeforeEach(func() {
		router := mux.NewRouter().StrictSlash(true)

		// AddRoutes(router)
		svr = httptest.NewServer(router)
		uri = svr.URL + "/health"
	})

	AfterEach(func() {
		svr.Close()
	})

	It("Returns 200", func() {
		req, _ := http.NewRequest(http.MethodGet, uri, nil)

		res, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		Expect(res.StatusCode).To(Equal(http.StatusOK))
	})

})
