package utils

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestIsValidEmail(t *testing.T) {
	t.Parallel()
	t.Run("email", func(t *testing.T) {
		email := gofakeit.Email()
		expected := true

		require.Equal(t, expected, IsValidEmail(email))
	})
	t.Run("no valid email", func(t *testing.T) {
		email := gofakeit.BeerHop()
		expected := false

		require.Equal(t, expected, IsValidEmail(email))
	})
}
