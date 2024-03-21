func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"https://www.youtube.com/watch?v=12345678901", "12345678901", nil},
		{"https://www.youtube.com/embed/12345678901", "12345678901", nil},
		{"https://youtu.be/12345678901", "12345678901", nil},
		{"https://www.youtube.com/shorts/12345678901", "12345678901", nil},
		{"12345678901", "12345678901", nil},
		{"http://www.example.com/?v=12345678901", "", ErrInvalidCharactersInVideoID},
		{"1234567", "", ErrVideoIDMinLength},
	}

	for _, test := range tests {
		result, err := ExtractVideoID(test.input)
		if result != test.expected {
			t.Errorf("ExtractVideoID(%s) = %s, expected %s", test.input, result, test.expected)
		}
		if err != test.err {
			if err == nil || test.err == nil || err.Error() != test.err.Error() {
				t.Errorf("ExtractVideoID(%s) error = %v, expected %v", test.input, err, test.err)
			}
		}
	}
}