// Code generated by github.com/alextanhongpin/getter, DO NOT EDIT.

package examples

import "database/sql"

func (b Bar) Foo() Foo {
	return b.foo
}

func (b Bar) Version() int64 {
	return b.version
}

func (f Foo) String() string {
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

func (f Foo) SliceStrp() []*string {
	return f.sliceStrp
}

func (f Foo) SliceIntp() []*int {
	return f.sliceIntp
}

func (f Foo) SliceBoolp() []*bool {
	return f.sliceBoolp
}

func (f Foo) SliceBoolFoo() []Foo {
	return f.sliceBoolFoo
}

func (f Foo) SliceBoolFoop() []*Foo {
	return f.sliceBoolFoop
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

func (f Foo) Remarks() *string {
	return f.fizz.Remarks
}

func (f Foo) NestedBar() *Bar {
	return f.fizz.bar
}

func (f Foo) BuzzItName() string {
	return f.buzz.Name
}

func (f Foo) BuzzItAge() int64 {
	return f.buzz.Age
}
