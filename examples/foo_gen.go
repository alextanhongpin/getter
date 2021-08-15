// Code generated by reader, DO NOT EDIT.
package examples

import "database/sql"

func (f Foo) Str() string {
	return f.str
}

func (f Foo) I() int {
	return f.i
}

func (f Foo) I32() int32 {
	return f.i32
}

func (f Foo) I64() int64 {
	return f.i64
}

func (f Foo) B() bool {
	return f.b
}

func (f Foo) SliceStr() []string {
	return f.sliceStr
}

func (f Foo) SliceInt() []int {
	return f.sliceInt
}

func (f Foo) SliceBool() []bool {
	return f.sliceBool
}

func (f Foo) BoolsByString() map[string][]bool {
	return f.boolsByString
}

func (f Foo) BoolByString() map[string]bool {
	return f.boolByString
}

func (f Foo) FooTypesByString() map[string][]FooType {
	return f.fooTypesByString
}

func (f Foo) FooTypePtrsByString() map[string][]*FooType {
	return f.fooTypePtrsByString
}

func (f Foo) StringByFooType() map[FooType]string {
	return f.stringByFooType
}

func (f Foo) StringByPtrFooType() map[*FooType]string {
	return f.stringByPtrFooType
}

func (f Foo) StrPtr() *string {
	return f.strPtr
}

func (f Foo) IntPtr() *int {
	return f.intPtr
}

func (f Foo) BoolPtr() *bool {
	return f.boolPtr
}

func (f Foo) Imported() sql.NullString {
	return f.imported
}

func (f Foo) AliasPtr() *FooType {
	return f.aliasPtr
}

func (f Foo) Alias() FooType {
	return f.alias
}
