package scan

import (
	"fmt"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	reader := strings.NewReader(" (\t foo123 bar: baz.)\n")
	scanner := NewScanner(reader)

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := LeftParen, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Identifier, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := "foo123", token.Value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Identifier, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := "bar:", token.Value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Identifier, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := "baz.", token.Value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := RightParen, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestDot(t *testing.T) {
	reader := strings.NewReader(".")
	scanner := NewScanner(reader)

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Dot, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestEmptyInput(t *testing.T) {
	reader := strings.NewReader("")
	scanner := NewScanner(reader)

	if token, err := scanner.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Eof, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestIllegalInput(t *testing.T) {
	reader := strings.NewReader("@")
	scanner := NewScanner(reader)

	if _, err := scanner.NextToken(); err == nil {
		t.Fatalf("expected err")
	} else if expected, actual := "scan: illegal rune: 64", err.Error(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReaderError(t *testing.T) {
	reader := &mockReader{}
	scanner := NewScanner(reader)

	if _, err := scanner.NextToken(); err == nil {
		t.Fatalf("expected err")
	} else if expected, actual := "mock reader", err.Error(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestTokenString(t *testing.T) {
	reader := strings.NewReader("'(x . y)")
	scanner := NewScanner(reader)

	if expected, actual := "Apostrophe '", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "LeftParen (", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "Identifier x", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "Dot .", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "Identifier y", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "RightParen )", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "Eof ", mustScan(scanner).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "Unknown ?", (&Token{-1, "?"}).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func mustScan(scanner *Scanner) *Token {
	if token, err := scanner.NextToken(); err != nil {
		panic("mustScan failed: " + err.Error())
	} else {
		return token
	}
}

type mockReader struct {
}

func (m *mockReader) Read(p []byte) (n int, err error) {
	fmt.Printf("Read %v", len(p))
	return 0, fmt.Errorf("mock reader")
}
