package service

import "testing"

func TestValidateAndFormatString(t *testing.T) {
	s := &ShrtnrService{}

	t.Run("https valid", func(t *testing.T) {
		expected := "https://google.com/"
		got, err := s.ValidateAndFormatLink("https://google.com/")
		if err != nil {
			t.Errorf("Got error on valid URL: https://google.com/ %s", err)
		}
		if expected != got {
			t.Errorf("Expected %s; got %s", expected, got)
		}
	})

	t.Run("http valid", func(t *testing.T) {
		expected := "http://google.com/"
		got, err := s.ValidateAndFormatLink("http://google.com/")
		if err != nil {
			t.Errorf("Got error on valid URL: http://google.com/ %s", err)
		}
		if expected != got {
			t.Errorf("Expected %s; got %s", expected, got)
		}
	})

	t.Run("no scheme valid", func(t *testing.T) {
		expected := "https://google.com"
		got, err := s.ValidateAndFormatLink("google.com")
		if err != nil {
			t.Errorf("Got error on valid URL: google.com %s", err)
		}
		if expected != got {
			t.Errorf("Expected %s; got %s", expected, got)
		}
	})

	t.Run("unsupported scheme", func(t *testing.T) {
		_, err := s.ValidateAndFormatLink("dsn://google.com")
		if err == nil {
			t.Errorf("Expected to fail on invalid link dsn://google.com")
		}
	})

	t.Run("WTF scheme", func(t *testing.T) {
		_, err := s.ValidateAndFormatLink("asdfasdfhn https://google.com")
		if err == nil {
			t.Errorf("Expected to fail on invalid link asdfasdfhn https://google.com")
		}
	})
}
