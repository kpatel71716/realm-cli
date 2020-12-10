package whoami

import (
	"bytes"
	"testing"

	"github.com/10gen/realm-cli/internal/cli"
	"github.com/10gen/realm-cli/internal/utils/test/assert"
	"github.com/10gen/realm-cli/internal/utils/test/mock"
)

func TestHandler(t *testing.T) {
	t.Run("Handler should run as a noop", func(t *testing.T) {
		cmd := &command{}

		err := cmd.Handler(nil, nil, nil)
		assert.Nil(t, err)
	})
}

func TestFeedback(t *testing.T) {
	t.Run("Feedback should print the auth details", func(t *testing.T) {
		for _, tc := range []struct {
			description string
			setup       func(t *testing.T, profile *cli.Profile)
			test        func(t *testing.T, output string)
		}{
			{
				description: "with no user logged in",
				test: func(t *testing.T, output string) {
					assert.Equal(t, "01:23:45 UTC INFO  No user is currently logged in\n", output)
				},
			},
			{
				description: "with a user that has no active session",
				setup: func(t *testing.T, profile *cli.Profile) {
					profile.SetUser("username", "my-super-secret-key")
				},
				test: func(t *testing.T, output string) {
					assert.Equal(t, "01:23:45 UTC INFO  The user, username, is not currently logged in\n", output)
				},
			},
			{
				description: "with a user fully logged in",
				setup: func(t *testing.T, profile *cli.Profile) {
					profile.SetUser("username", "my-super-secret-key")
					profile.SetSession("accessToken", "refreshToken")
				},
				test: func(t *testing.T, output string) {
					assert.Equal(t, "01:23:45 UTC INFO  Currently logged in user: username (**-*****-******-key)\n", output)
				},
			},
		} {
			t.Run(tc.description, func(t *testing.T) {
				profile := mock.NewProfile(t)

				if tc.setup != nil {
					tc.setup(t, profile)
				}

				buf := new(bytes.Buffer)
				ui := mock.NewUI(mock.UIOptions{}, buf)

				cmd := &command{}
				err := cmd.Feedback(profile, ui)
				assert.Nil(t, err)

				tc.test(t, buf.String())
			})
		}
	})
}