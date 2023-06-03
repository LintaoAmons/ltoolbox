package functions

import (
	"strings"
	"unicode"
)

func shouldInsertSpace(i int, r, previousRune rune) bool {
	return unicode.IsUpper(r) && i > 0 && !unicode.IsSpace(previousRune) &&
		// This condition checks if the previous rune is not an uppercase letter.
		// If the previous rune is an uppercase letter, it indicates that the current
		// uppercase letter is part of an abbreviation or a sequence of uppercase letters, and there's no need to insert a space.
		!unicode.IsUpper(previousRune)
}

func insertSpacesBeforeUpper(s string) string {
	var result []rune
	previousRune := ' '

	for i, r := range s {
		if shouldInsertSpace(i, r, previousRune) {
			result = append(result, ' ')
		}
		result = append(result, r)
		previousRune = r
	}
	return string(result)
}

func replaceNonAlphanumericWithSpace(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return ' '
	}
	return r
}

func processWords(s string, mode string) string {
	words := strings.Fields(s)

	switch mode {
	case "constant":
		for i := range words {
			words[i] = strings.ToUpper(words[i])
		}
		return strings.Join(words, "_")
	case "snake":
		return strings.Join(words, "_")
	case "camel":
		for i := range words {
			if i > 0 {
				words[i] = strings.Title(strings.ToLower(words[i]))
			} else {
				words[i] = strings.ToLower(words[i])
			}
		}
		return strings.Join(words, "")
	case "kebab":
		for i := range words {
			words[i] = strings.ToLower(words[i])
		}
		return strings.Join(words, "-")
	default:
		panic("Not supported")
	}
}

func ToConstantCase(s string) string {
	s = insertSpacesBeforeUpper(s)
	s = strings.Map(replaceNonAlphanumericWithSpace, s)
	return processWords(s, "constant")
}

func ToCamelCase(s string) string {
	s = insertSpacesBeforeUpper(s)
	s = strings.Map(replaceNonAlphanumericWithSpace, s)
	return processWords(s, "camel")
}

func ToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	s = insertSpacesBeforeUpper(s)
	s = strings.Map(replaceNonAlphanumericWithSpace, s)
	return processWords(s, "kebab")
}

func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	// Find the number of leading underscores
	leadingUnderscores := ""
	for _, r := range s {
		if r == '_' {
			leadingUnderscores += "_"
		} else {
			break
		}
	}

	// Find the number of trailing underscores
	trailingUnderscores := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '_' {
			trailingUnderscores += "_"
		} else {
			break
		}
	}

	// Insert spaces before uppercase letters in camelCase words, starting after the leading underscores and ending before the trailing underscores
	s = s[:len(leadingUnderscores)] + insertSpacesBeforeUpper(s[len(leadingUnderscores):len(s)-len(trailingUnderscores)]) + s[len(s)-len(trailingUnderscores):]

	// Replace any non-alphanumeric characters with a space
	spaceReplaced := strings.Map(replaceNonAlphanumericWithSpace, s)

	// Split the string into words using spaces as delimiters
	words := strings.Fields(spaceReplaced)

	// Convert each word to lowercase
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	words[0] = leadingUnderscores + words[0]
	words[len(words)-1] = words[len(words)-1] + trailingUnderscores

	// Join the words together with an underscore ('_')
	return strings.Join(words, "_")
}
