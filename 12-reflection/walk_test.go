package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{
				"Chris",
				30,
			},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

// package reflection
//
// import (
// 	"reflect"
// 	"testing"
// )
//
// func TestWakl(t *testing.T) {
// 	expected := "Chris"
// 	var got []string
//
// 	cases := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 	}
//
// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			var got []string
// 			walk(test.Input, func(input string) {
// 				got = append(got, input)
// 			})
// 		})
//
// 		if !reflect.DeepEqual(got, test.ExpectedCalls) {
// 			t.Errorf("got %v, want %v", got, test.ExpectedCalls)
// 		}
// 	}
//
// 	x := struct {
// 		Name string
// 	}{expected}
//
// 	walk(x, func(input string) {
// 		got = append(got, input)
// 	})
//
// 	if len(got) != 1 {
// 		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
// 	}
//
// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}
// }
