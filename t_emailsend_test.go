package wrpgomail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tested with: https://github.com/camptocamp/docker_smtp
func TestEmailSend(t *testing.T) {
	cfg := EmailServerConfig{
		URI:      "0.0.0.0",
		Port:     2525,
		User:     "",
		Password: "",
	}

	s, errNew := NewEmailServer(cfg)
	if assert.Nil(t, errNew) {
		em := EmailData{
			Important:   true,
			Subject:     "Test Embed",
			From:        "john@loco.com",
			To:          []string{"mary@loco.com"},
			MessageHTML: `Hi! <img src="cid:x.png" alt="My image" /> Second: <img src="cid:y.png" alt="My image" />`,
			Embedded:    []string{"x.png", "y.png"},
		}

		assert.Nil(t, s.SendEmail(em))
	}
}
