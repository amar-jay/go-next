package models

import (
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// MarshalUUID marshals a UUID to a string
func MarshallUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(id.String())
}

// UnmarshalUUID unmarshals a string to a UUID
func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	idxStr, ok := v.(string)
	if !ok {
		return uuid.Nil, errors.New("uuid must be a string")
	}

	return uuid.Parse(idxStr)
}
