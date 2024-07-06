package campaign

import (
	"errors"
	"goapi/internal/contract"
	internalerrors "goapi/internal/internalErrors"
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

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test A",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""
	_, err := service.Create(newCampaign)
	if err != nil {
		assert.Equal("Name is required", err.Error())
	} else {
		assert.Fail("Expected error but got nil")
	}

}

func Test_Create_ValidateNameSize(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = "adhasjdhaskjdhsdjdhasgdhjasgdhjsgdhjgas"
	_, err := service.Create(newCampaign)
	if err != nil {
		assert.Equal("The Name must have a maximum of 30 characters", err.Error())
	} else {
		assert.Fail("Expected error but got nil")
	}
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("erro to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
