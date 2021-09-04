# reader

Generate getters for golang using tags.

By default, getters for all private fields will be generated. To ignore, add `get:"-"` tag.


## Installation

```bash
$ go install github.com/alextanhongpin/getter
```

## Using CLI

```bash
$ getter \
	-in=path/to/your/input/file.go \
	-type=YourStructType1,YourStructType2 \ 	    # Accepts multiple struct names.
	-out=optional/path/to/your/output/dir/or/file.go \  # Generated based on struct name with _gen.go suffix if not provided.
	-prefix=OptionalPrefix                              # Adds a custom prefix, e.g. Get, by default no prefix will be added.
```

## Using go generate

```go
package main

//go:generate getter -type=User -prefix=Get
type User struct {
	id      uuid.UUID
	name    string  `get:"RealName"`        // With custom field name.
	remarks *string `get:"-"`               // Ignore field.
	account Account `get:",inline,Account"` // Inlines struct, with custom prefix.
}

type Account struct {
	id int64
}
```

## Features

- generate getters with custom prefix
- tags to ignore field, rename field, inline struct, add inline prefix
