// Code generated by github.com/alextanhongpin/getter, DO NOT EDIT.

package examples

import (
	"encoding/json"
	uuid "github.com/google/uuid"
)

func (i Imported) ID() uuid.UUID {
	return i.id
}

func (i Imported) Extra() json.RawMessage {
	return i.extra
}
