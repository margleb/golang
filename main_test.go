package golang

import "testing"

// дублирует функции тестирования
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		result, err := devide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}
		if result != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, result)
		}
	}
}

//// тестируем деление
//func TestDivide(t *testing.T) {
//	_, err := devide(10.0, 1.0)
//	if err != nil {
//		t.Error("Got an error when should not have")
//	}
//}
//
//// тестируем плохое деление
//func TestBadDivide(t *testing.T) {
//	_, err := devide(10.0, 0)
//	if err == nil {
//		t.Error("Did not get an error when we should have")
//	}
//}
