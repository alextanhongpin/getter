// Code generated by access, DO NOT EDIT.
package examples

import uuid "github.com/google/uuid"

func (b Bar) ID() uuid.UUID {
	return b.id
}

func (b *Bar) SetID(id uuid.UUID) *Bar {
	b.id = id
	return b
}

func (b Bar) Name() string {
	return b.name
}

func (b *Bar) SetName(name string) *Bar {
	b.name = name
	return b
}

func (b Bar) Age() int64 {
	return b.age
}

func (b *Bar) SetAge(age int64) *Bar {
	b.age = age
	return b
}

func (b Bar) URL() string {
	return b.url
}

func (b *Bar) SetURL(url string) *Bar {
	b.url = url
	return b
}
