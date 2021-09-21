package instagram

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type mockRoundTripper struct {
	mock.Mock
}

func (m *mockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	arg := m.Called(request)
	return arg.Get(0).(*http.Response), arg.Error(1)
}

func TestFetchProfile(t *testing.T) {
	t.Parallel()
	agent := http.DefaultClient
	rt := &mockRoundTripper{}
	html, _ := os.ReadFile("./mock.html")
	body := ioutil.NopCloser(bytes.NewBuffer(html))
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}
	rt.On("RoundTrip", mock.Anything).Return(resp, nil)
	agent.Transport = rt
	handler := New(agent)
	res, err := handler.RetrieveProfile("mgufrone")
	assert.Nil(t, err)
	assert.Equal(t, 0.15484515484515485, res.Rate())
	req, _ := rt.Calls[0].Arguments[0].(*http.Request)
	rt.AssertExpectations(t)
	assert.Equal(t, "https://www.instagram.com/mgufrone/", req.URL.String())
}
