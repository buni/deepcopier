package deepcopier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Status string

type Request struct {
	Status *string
}
type RequestStr struct {
	Status string
}
type Entity struct {
	Status Status
}

func TestBasicTypeToUserDefinedType(t *testing.T) {
	t.Run("successfully copy string to user defined string type", func(t *testing.T) {
		str := "active"
		request := &Request{
			Status: &str,
		}

		entity := &Entity{
			Status: "",
		}

		err := Copy(request).To(entity)
		assert.NoError(t, err)
		assert.EqualValues(t, str, entity.Status)
	})
	t.Run("string type is nil, user defined type should remain empty", func(t *testing.T) {
		request := &Request{}

		entity := &Entity{
			Status: "",
		}

		err := Copy(request).To(entity)
		assert.NoError(t, err)
		assert.Empty(t, entity.Status)
	})
	t.Run("successfully copy string to user defined string type", func(t *testing.T) {
		str := "active"
		request := &RequestStr{
			Status: str,
		}

		entity := &Entity{
			Status: "",
		}

		err := Copy(request).To(entity)
		assert.NoError(t, err)
		assert.EqualValues(t, str, entity.Status)
	})
}
