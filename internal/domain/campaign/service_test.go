package campaign

import (
	"goapi/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	NewCampaign := contract.NewCampaign{
		Name:    "Test A",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}

	id, err := service.Create(NewCampaign)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	NewCampaign := contract.NewCampaign{
		Name:    "Test A",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != NewCampaign.Name ||
			campaign.Content != NewCampaign.Content ||
			len(campaign.Contacts) != len(NewCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(NewCampaign)

	repositoryMock.AssertExpectations(t)
}
