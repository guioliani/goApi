package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	Id        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedOn time.Time
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:        "1",
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
