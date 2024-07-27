package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goapi/internal/contract"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func Test_CampaignsPost_Should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "hi everyone",
		Emails:  []string{"teste@email.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("342", nil)
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()
	_, status, err := handler.CreateCampaign(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaignsPost_Should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "hi everyone",
		Emails:  []string{"teste@email.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()
	_, _, err := handler.CreateCampaign(rr, req)

	assert.NotNil(err)

}
