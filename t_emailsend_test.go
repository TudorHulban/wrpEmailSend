package wrpgomail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testing with Mail Hog
func TestEmailSend(t *testing.T) {
	cfg := EmailServerConfig{
		URI:      "0.0.0.0",
		Port:     1025,
		User:     "",
		Password: "",
	}

	s, errNew := NewEmailServer(cfg)
	if assert.Nil(t, errNew) {
		em := EmailData{
			From:        "john@loco.com",
			To:          []string{"mary@loco.com"},
			MessageHTML: "hi <img src=\"cid:y.png\" alt=\"My image\" />",
			Attachments: []string{"x.png"},
			Embedded:    []string{"y.png"},
		}

		assert.Nil(t, s.SendEmail(em))
	}
}
