package functions_test

import (
	"testing"

	. "github.com/LintaoAmons/toolbox/functions"
)

func TestToCamelCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello_world_example", "helloWorldExample"},
		{"AAA_BBB_CCC", "aaaBbbCcc"},
		{"_starting_with_underscore", "startingWithUnderscore"},
		{"ending_with_underscore_", "endingWithUnderscore"},
		{"__double__underscore__", "doubleUnderscore"},
		{"all_lowercase", "allLowercase"},
		{"ALL_UPPERCASE", "allUppercase"},
		{"aaa bbb ccc", "aaaBbbCcc"},
		{"", ""}, // Empty string
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := ToCamelCase(tc.input)
			if actual != tc.expected {
				t.Errorf("ToCamelCase(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestToConstantCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello_world_example", "HELLO_WORLD_EXAMPLE"},
		{"aaaBbbCcc", "AAA_BBB_CCC"},
		{"starting with space", "STARTING_WITH_SPACE"},
		{"ending with space ", "ENDING_WITH_SPACE"},
		{"double  spaces", "DOUBLE_SPACES"},
		{"all_lowercase", "ALL_LOWERCASE"},
		{"ALL_UPPERCASE", "ALL_UPPERCASE"},
		{"MixedCaseWith_Spaces", "MIXED_CASE_WITH_SPACES"},
		{"__double__underscore__", "DOUBLE_UNDERSCORE"},
		{"", ""}, // Empty string
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := ToConstantCase(tc.input)
			if actual != tc.expected {
				t.Errorf("ToConstantCase(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"helloWorldExample", "hello_world_example"},
		{"AAA_BBB_CCC", "aaa_bbb_ccc"},
		{"_startingWithUnderscore", "_starting_with_underscore"},
		{"endingWithUnderscore_", "ending_with_underscore_"},
		{"__double__underscore__", "__double_underscore__"},
		{"allLowercase", "all_lowercase"},
		{"ALL_UPPERCASE", "all_uppercase"},
		{"MixedCaseWith_Spaces", "mixed_case_with_spaces"},
		{"", ""}, // Empty string
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := ToSnakeCase(tc.input)
			if actual != tc.expected {
				t.Errorf("ToSnakeCase(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"helloWorldExample", "hello-world-example"},
		{"AAA_BBB_CCC", "aaa-bbb-ccc"},
		{"starting with space", "starting-with-space"},
		{"ending with space ", "ending-with-space"},
		{"double  spaces", "double-spaces"},
		{"allLowercase", "all-lowercase"},
		{"ALL_UPPERCASE", "all-uppercase"},
		{"MixedCaseWith_Spaces", "mixed-case-with-spaces"},
		{"__double__underscore__", "double-underscore"},
		{"", ""}, // Empty string
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := ToKebabCase(tc.input)
			if actual != tc.expected {
				t.Errorf("ToKebabCase(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
