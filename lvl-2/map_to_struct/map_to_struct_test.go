package map_to_struct

import (
	"fmt"
	"reflect"
	"testing"
)

type Contacts struct {
	Email            string
	Phone            string
	AdditionalPhones map[string]string
}

type Person struct {
	Name     string
	Age      int64
	Contacts Contacts
}

type Person2 struct {
	Name     string
	Age      int64
	Contacts *Contacts
}
type Person3 struct {
	Name     string
	age      int64
	Contacts *Contacts
}

func TestFillStruct(t *testing.T) {
	data := make(map[string]interface{})
	data["Name"] = "John"
	data["Age"] = int64(23)
	contacts := make(map[string]interface{})
	contacts["Phone"] = "71234567890"
	contacts["Email"] = "e@mail.ru"
	contacts["AdditionalPhones"] = map[string]string{
		"Home": "74951111111",
		"Work": "74991234567",
	}
	data["Contacts"] = contacts
	valid := &Person{
		Name: "John",
		Age:  23,
		Contacts: Contacts{
			Phone: "71234567890",
			Email: "e@mail.ru",
			AdditionalPhones: map[string]string{
				"Home": "74951111111",
				"Work": "74991234567",
			},
		},
	}
	result := &Person{}
	err := FillStruct(result, data)
	if err != nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, valid)
	}
	if !reflect.DeepEqual(result, valid) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", result, valid)
	}
	// Error field does not exists
	data["Invalid"] = "Invalid"
	validError := fmt.Errorf("No such field: %s in obj", "Invalid")
	err = FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	// Error invalid type
	delete(data, "Invalid")
	validError = fmt.Errorf("Provided value type didn't match obj field type")
	data["Age"] = float64(12.21312)
	err = FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}

	// Valid inner struct is pointer
	data["Age"] = int64(23)
	valid2 := &Person2{
		Name: "John",
		Age:  23,
		Contacts: &Contacts{
			Phone: "71234567890",
			Email: "e@mail.ru",
			AdditionalPhones: map[string]string{
				"Home": "74951111111",
				"Work": "74991234567",
			},
		},
	}
	result2 := &Person2{}
	err = FillStruct(result2, data)
	if err != nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, valid2)
	}
	if !reflect.DeepEqual(result2, valid2) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", result2, valid2)
	}

	// Error unexported struct fields
	delete(data, "Age")
	data["age"] = 23
	result3 := &Person3{}
	validError = fmt.Errorf("Cannot set %s field value", "age")
	err = FillStruct(result3, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
