package gofiql

import "testing"

func TestFindToken(t *testing.T) {
	s := "name==Joe"
	i, b := findToken(&s)
	if i != -1 {
		t.Logf("Expected -1, got: %d", i)
		t.Fail()
	}
	bb := string(b)
	if string(b) != s {
		t.Logf("Expected same string in output, got: %s", bb)
		t.Fail()
	}

	s = "name==Joe,name=Tom"
	i, b = findToken(&s)
	if i == -1 {
		t.Logf("Expected != -1, got: %d", i)
		t.Fail()
	}
	bb = string(b)
	if string(b) != s {
		t.Logf("Expected same string in output, got: %s", bb)
		t.Fail()
	}

	s = "(name==Joe,name=Tom)"
	i, b = findToken(&s)
	if i == -1 {
		t.Logf("Expected != -1, got: %d", i)
		t.Fail()
	}
	bb = string(b)
	if string(b) == s {
		t.Logf("Expected different string in output, got: %s", bb)
		t.Fail()
	}
	s = s[1 : len(s)-1]
	if string(b) != s {
		t.Logf("Expected string trimmed of (), got: %s", bb)
		t.Fail()
	}

	s = "(((name==Joe,name=Tom)))"
	i, b = findToken(&s)
	if i == -1 {
		t.Logf("Expected != -1, got: %d", i)
		t.Fail()
	}
	bb = string(b)
	if string(b) == s {
		t.Logf("Expected different string in output, got: %s", bb)
		t.Fail()
	}
	s = s[3 : len(s)-3]
	if string(b) != s {
		t.Logf("Expected string trimmed of (), got: %s", bb)
		t.Fail()
	}

	s = "(((((name==Joe,name=Tom)));qty=gt=10))"
	i, b = findToken(&s)
	if i == -1 {
		t.Logf("Expected != -1, got: %d", i)
		t.Fail()
	}
	bb = string(b)
	if string(b) == s {
		t.Logf("Expected different string in output, got: %s", bb)
		t.Fail()
	}
	s = s[2 : len(s)-2]
	if string(b) != s {
		t.Logf("Expected string trimmed of (), got: %s", bb)
		t.Fail()
	}
}
