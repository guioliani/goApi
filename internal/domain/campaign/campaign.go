package campaign

import (
	internalerrors "goapi/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string `validate:"required,max=30"`
	Content   string
	Contacts  []Contact
	CreatedOn time.Time
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}

	err := internalerrors.ValidateStruct(*campaign)
	if err != nil {
		return nil, err
	}

	return campaign, err
}
