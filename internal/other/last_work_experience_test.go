package other_test

import (
	"testing"

	"github.com/juniorpen01/dsa/internal/other"
)

func TestLastWorkExperience(t *testing.T) {
	cases := []struct {
		input    []string
		expected *string
	}{
		{[]string{"Software Engineer", "Data Analyst", "Project Manager"}, strPtr("Project Manager")},
		{[]string{"Intern", "Junior Developer"}, strPtr("Junior Developer")},
		{[]string{}, nil},
		{[]string{"CEO"}, strPtr("CEO")},
		{[]string{"Cashier", "Supervisor", "Manager", "Director"}, strPtr("Director")},
	}

	for _, c := range cases {
		other.LastWorkExperience(c.input)
		if actual, expected := other.LastWorkExperience(c.input), c.expected; expected != nil {
			if *actual != *expected {
				t.Errorf("Expected experience does not match actual experience: expected %d, got %d", expected, actual)
			}
		} else {
			if actual != expected {
				t.Errorf("Expected experience does not match actual experience: expected %d, got %d", expected, actual)
			}
		}
	}
}

func strPtr(str string) *string {
	return &str
}
