package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@gmail.com", "email2@gmail.com"}
)

func Test_Create_MustValidateNameEmpty(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)
	assert.Equal("Name is required", err.Error())

}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("sdjahsdjkhasdjkashdjkhasjdhassadasd", content, contacts)
	assert.Equal("The Name must have a maximum of 30 characters", err.Error())
}
