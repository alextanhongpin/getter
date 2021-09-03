package loader

import (
	"flag"
	"fmt"
	"go/types"
	"os"
	"strings"
)

// StructField for the example below.
//type Foo struct {
//  Name sql.NullString `json:"name"
//}
type StructField struct {
	CustomName string `example:"CustomName"` // Defaults to name if no custom name is provided.
	// Name of the struct field.
	Name string `example:"Name"`

	// Useful when the output directory doesn't match the existing ones.
	PkgPath string `example:"github.com/alextanhongpin/go-codegen/test"`
	PkgName string `example:"test"`

	Exported bool `example:"true"`
	Tag      *Tag `example:"get:'Renamed'"` // To ignore getter.
	*Type
	Ordinal int
}

type Type struct {
	Name         string `example:"NullString"`
	PkgPath      string `example:"database/sql"`
	IsPointer    bool
	IsCollection bool // Whether it's an array or slice.
	IsStruct     bool
	IsMap        bool
	MapKey       *Type
	MapValue     *Type
	T            types.Type
	E            types.Type
}

// NewType recursively checks for the field type.
func NewType(orityp types.Type) (*Type, error) {
	var isPointer, isCollection, isStruct, isMap bool
	var fieldPkgPath, fieldType string
	var mapKey, mapValue *Type
	var err error
	typ := orityp

	switch t := typ.(type) {
	case *types.Slice:
		isCollection = true
		typ = t.Elem()
	case *types.Array:
		isCollection = true
		typ = t.Elem()
	case *types.Map:
		isMap = true
		var err error
		mapKey, err = NewType(t.Key())
		if err != nil {
			return nil, err
		}
		mapValue, err = NewType(t.Elem())
		if err != nil {
			return nil, err
		}
	}

	// In case the slice or array is pointer, we take the elem again.
	switch t := typ.(type) {
	case *types.Pointer:
		isPointer = true
		typ = t.Elem()
	}

	switch t := typ.(type) {
	case *types.Struct:
		isStruct = true
		if err != nil {
			return nil, err
		}
	case *types.Named:
		obj := t.Obj()
		fieldPkgPath = obj.Pkg().Path()
		fieldType = obj.Name()

		switch v := t.Underlying().(type) {
		case *types.Struct:
			isStruct = true
			typ = v
		default:
			typ = v
		}
	default:
		fieldType = t.String()
	}

	return &Type{
		Name:         fieldType,
		PkgPath:      fieldPkgPath,
		IsCollection: isCollection,
		IsPointer:    isPointer,
		IsMap:        isMap,
		IsStruct:     isStruct,
		MapKey:       mapKey,
		MapValue:     mapValue,
		T:            orityp,
		E:            typ,
	}, nil
}

type Option struct {
	In         string
	Out        string
	PkgName    string
	PkgPath    string
	Prefix     string
	StructName string
	Type       *Type
}

type Generator func(opt Option) error

func New(fn Generator) error {
	structPtr := flag.String("type", "", "the target struct name")
	inp := flag.String("in", os.Getenv("GOFILE"), "the input file, defaults to the file with the go:generate comment")
	outp := flag.String("out", "", "the output directory")
	prefixp := flag.String("prefix", "", "the generated field name prefix, e.g. Get")
	flag.Parse()

	in := fullPath(*inp)
	pkg := loadPackage(packagePath(in)) // github.com/your-github-username/your-pkg.
	pkgPath := pkg.PkgPath              // Specify the config packages.NeedName to get this value.
	pkgName := pkg.Name                 // main

	// Allows -type=Foo,Bar
	structNames := strings.Split(*structPtr, ",")
	for _, structName := range structNames {
		out := FileNameFromTypeName(*inp, *outp, structName)
		obj := pkg.Types.Scope().Lookup(structName)
		if obj == nil {
			panic(fmt.Errorf("loader: struct %s not found", structName))
		}

		// Check if it is a declared typed.
		if _, ok := obj.(*types.TypeName); !ok {
			panic(fmt.Errorf("loader: %v is not a named type", obj))
		}

		// Check if the type is a struct.
		structType, ok := obj.Type().Underlying().(*types.Struct)
		if !ok {
			panic(fmt.Errorf("loader: %v is not a struct", obj))
		}

		t, err := NewType(structType)
		if err != nil {
			return err
		}

		if err := fn(Option{
			PkgName:    pkgName,
			PkgPath:    pkgPath,
			Prefix:     *prefixp,
			Out:        out,
			In:         in,
			StructName: structName,
			Type:       t,
		}); err != nil {
			return err
		}
	}
	return nil
}

func ExtractStructFields(structType *types.Struct) (map[string]StructField, error) {
	fields := make(map[string]StructField)
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		tag, err := NewTag(structType.Tag(i))
		if err != nil {
			return nil, err
		}

		name := field.Name()
		if tag != nil {
			if tag.Skip {
				continue
			}
			if tag.Name != "" {
				name = tag.Name
			}
		}

		t, err := NewType(field.Type())
		if err != nil {
			return nil, err
		}

		fields[name] = StructField{
			CustomName: name,
			Name:       field.Name(),
			PkgPath:    field.Pkg().Path(),
			Exported:   field.Exported(),
			Type:       t,
			Tag:        tag,
			Ordinal:    i,
		}
	}

	return fields, nil
}
