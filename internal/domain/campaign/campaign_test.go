package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@gmail.com", "email2@gmail.com"}
)

func TestShouldCreateNewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNewCampaignIdIsNotNill(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func TestNewCampaignCreatedOnisNotNill(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func TestNewCampaignShouldNotCreateNameNil(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}
