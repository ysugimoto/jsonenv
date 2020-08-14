package jsonenv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ysugimoto/jsonenv"
)

func TestJsonEnv(t *testing.T) {
	t.Run("Should return error if env key does not exist", func(t *testing.T) {
		os.Setenv("jsonenv-test01", `{"EXAMPLE":"VALUE"}`)
		defer func() {
			assert.NoError(t, os.Unsetenv("jsonenv-test01"))
		}()
		assert.Error(t, jsonenv.Extract("jsonenv-test00"))
	})

	t.Run("Should return error if unmarshal JSON fails", func(t *testing.T) {
		os.Setenv("jsonenv-test02", `{"EXAMPLE":"VALUE"`)
		defer func() {
			assert.NoError(t, os.Unsetenv("jsonenv-test02"))
		}()
		assert.Error(t, jsonenv.Extract("jsonenv-test02"))
	})

	t.Run("Should success with exact env value", func(t *testing.T) {
		os.Setenv("jsonenv-test03", `{"EXAMPLE":"VALUE"}`)
		defer func() {
			assert.NoError(t, os.Unsetenv("jsonenv-test03"))
		}()
		assert.NoError(t, jsonenv.Extract("jsonenv-test03"))
		assert.Equal(t, "VALUE", os.Getenv("EXAMPLE"))
	})
}
