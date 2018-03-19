package multiconfig

import (
	"testing"
)

type ConfigSingle struct {
	TestField    string `toml:"test_field" default:"test"`
	AnotherField string `toml:"another_field" default:"another test"`
	Nested       Nested
}

type ConfigSlice struct {
	TestField    string `toml:"test_field" default:"test"`
	AnotherField string `toml:"another_field" default:"another test"`
	Nested       []Nested
}

type Nested struct {
	NestedField        string `toml:"nested_field" default:"nested"`
	AnotherNestedField string `toml:"another_nested_field"`
}

func TestDefaultsSingle(t *testing.T) {
	doSingleTest(t, "defaultsingle.toml")
	doSingleTest(t, "defaultsingle.json")
	doSingleTest(t, "defaultsingle.yaml")
}

func doSingleTest(t *testing.T, configFile string) {
	expected := ConfigSingle{
		TestField:    "foo",
		AnotherField: "another test",
		Nested: Nested{
			NestedField:        "nested",
			AnotherNestedField: "bar",
		}}

	m := NewWithPath("testdata/" + configFile)
	config := new(ConfigSingle)
	m.MustLoad(config)

	if config.TestField != expected.TestField {
		t.Errorf("Test file %s, TestField value is wrong: %s, want: %s", configFile, config.TestField, expected.TestField)
	}
	if config.AnotherField != expected.AnotherField {
		t.Errorf("Test file %s, AnotherField value is wrong: %s, want: %s", configFile, config.AnotherField, expected.AnotherField)
	}
	if config.Nested.NestedField != expected.Nested.NestedField {
		t.Errorf("Test file %s, NestedField value is wrong: %s, want: %s", configFile, config.Nested.NestedField, expected.Nested.NestedField)
	}
	if config.Nested.AnotherNestedField != expected.Nested.AnotherNestedField {
		t.Errorf("Test file %s, AnotherNestedField value is wrong: %s, want: %s", configFile, config.Nested.AnotherNestedField, expected.Nested.AnotherNestedField)
	}
}

func TestDefaultsSlice(t *testing.T) {
	doSliceTest(t, "defaultslice.toml")
	doSliceTest(t, "defaultslice.json")
	doSliceTest(t, "defaultslice.yaml")
}

func doSliceTest(t *testing.T, configFile string) {
	expected := ConfigSlice{
		TestField:    "foo",
		AnotherField: "another test",
		Nested: []Nested{Nested{
			NestedField:        "nested",
			AnotherNestedField: "bar",
		}, Nested{
			NestedField:        "quux",
			AnotherNestedField: "baz",
		}}}

	m := NewWithPath("testdata/" + configFile)
	config := new(ConfigSlice)
	m.MustLoad(config)

	if config.TestField != expected.TestField {
		t.Errorf("Test file %s, TestField value is wrong: %s, want: %s", configFile, config.TestField, expected.TestField)
	}
	if config.AnotherField != expected.AnotherField {
		t.Errorf("Test file %s, AnotherField value is wrong: %s, want: %s", configFile, config.AnotherField, expected.AnotherField)
	}
	if len(config.Nested) != len(expected.Nested) {
		t.Errorf("Test file %s, Nested length is wrong: %d, want: %d", configFile, len(config.Nested), len(expected.Nested))
	} else {
		for i := range config.Nested {
			if config.Nested[i].NestedField != expected.Nested[i].NestedField {
				t.Errorf("Test file %s, NestedField value is wrong: %s, want: %s", configFile, config.Nested[i].NestedField, expected.Nested[i].NestedField)
			}
			if config.Nested[i].AnotherNestedField != expected.Nested[i].AnotherNestedField {
				t.Errorf("Test file %s, AnotherNestedField value is wrong: %s, want: %s", configFile, config.Nested[i].AnotherNestedField, expected.Nested[i].AnotherNestedField)
			}
		}
	}
}
