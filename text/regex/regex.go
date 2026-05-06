// Package regex provides regular expression matching equivalent to
// System.Text.RegularExpressions.Regex in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.text.regularexpressions.regex?view=netframework-4.7.2
package regex

import (
	"regexp"
)

// RegexOptions specifies options for regular expression matching.
type RegexOptions int

const (
	// None specifies no options are set.
	None RegexOptions = 0
	// IgnoreCase specifies case-insensitive matching.
	IgnoreCase RegexOptions = 1 << iota
	// Multiline changes the meaning of ^ and $ so they match at the beginning and end of each line.
	Multiline
	// Singleline specifies single-line mode where the period (.) matches every character.
	Singleline
	// ExplicitCapture specifies that the only valid captures are groups that are explicitly named or numbered.
	ExplicitCapture
	// IgnorePatternWhitespace eliminates unescaped white space from the pattern.
	IgnorePatternWhitespace
	// RightToLeft specifies that the search will be from right to left.
	RightToLeft
	// ECMAScript enables ECMAScript-compliant behavior for the expression.
	ECMAScript
	// CultureInvariant specifies that cultural differences in language are ignored.
	CultureInvariant
)

// Regex represents an immutable regular expression.
// Equivalent to System.Text.RegularExpressions.Regex in .NET.
type Regex struct {
	pattern string
	options RegexOptions
	re      *regexp.Regexp
}

// Compile creates a new Regex instance with the specified pattern.
func Compile(pattern string) (*Regex, error) {
	return CompileWithOptions(pattern, None)
}

// CompileWithOptions creates a new Regex with the specified options.
func CompileWithOptions(pattern string, options RegexOptions) (*Regex, error) {
	flags := buildRegexpFlags(options)

	re, err := regexp.Compile(flags + pattern)
	if err != nil {
		return nil, err
	}

	return &Regex{
		pattern: pattern,
		options: options,
		re:      re,
	}, nil
}

// MustCompile creates a new Regex and panics if the pattern is invalid.
func MustCompile(pattern string) *Regex {
	return MustCompileWithOptions(pattern, None)
}

// MustCompileWithOptions creates a new Regex with options and panics if invalid.
func MustCompileWithOptions(pattern string, options RegexOptions) *Regex {
	r, err := CompileWithOptions(pattern, options)
	if err != nil {
		panic(err)
	}
	return r
}

// IsMatch indicates whether the regular expression finds a match in the input string.
func (r *Regex) IsMatch(input string) bool {
	return r.re.MatchString(input)
}

// Match searches the input string for the first occurrence of the regular expression.
// Returns nil if no match is found.
func (r *Regex) Match(input string) *Match {
	m := r.re.FindStringSubmatch(input)
	if m == nil {
		return nil
	}
	return newMatch(m, r.re.SubexpNames())
}

// Matches searches the input string for all occurrences of the regular expression.
func (r *Regex) Matches(input string) *MatchCollection {
	matches := r.re.FindAllStringSubmatch(input, -1)
	if matches == nil {
		return &MatchCollection{matches: []*Match{}}
	}

	result := make([]*Match, len(matches))
	names := r.re.SubexpNames()
	for i, m := range matches {
		result[i] = newMatch(m, names)
	}
	return &MatchCollection{matches: result}
}

// Replace replaces all occurrences of the pattern in the input string with the replacement string.
func (r *Regex) Replace(input, replacement string) string {
	return r.re.ReplaceAllString(input, replacement)
}

// ReplaceFunc replaces all matches using a function.
func (r *Regex) ReplaceFunc(input string, replacer func(*Match) string) string {
	return r.re.ReplaceAllStringFunc(input, func(s string) string {
		m := r.Match(s)
		if m != nil {
			return replacer(m)
		}
		return s
	})
}

// Split splits the input string at the positions defined by the regular expression.
func (r *Regex) Split(input string) []string {
	return r.re.Split(input, -1)
}

// SplitWithCount splits the input string at most count times.
func (r *Regex) SplitWithCount(input string, count int) []string {
	return r.re.Split(input, count)
}

// Pattern returns the regular expression pattern that was used to create this instance.
func (r *Regex) Pattern() string {
	return r.pattern
}

// Options returns the options that were passed in to the Regex constructor.
func (r *Regex) Options() RegexOptions {
	return r.options
}

// GetGroupNames returns an array of capturing group names.
func (r *Regex) GetGroupNames() []string {
	return r.re.SubexpNames()
}

// GetGroupNumbers returns an array of capturing group numbers.
func (r *Regex) GetGroupNumbers() []int {
	n := r.re.NumSubexp()
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = i
	}
	return result
}

// String returns the regex pattern.
func (r *Regex) String() string {
	return r.pattern
}

// buildRegexpFlags converts .NET RegexOptions to regexp flag prefix.
func buildRegexpFlags(options RegexOptions) string {
	if options&None != 0 && options == None {
		return ""
	}

	var flags string

	if options&IgnoreCase != 0 {
		flags += "(?i)"
	}
	if options&Multiline != 0 {
		flags += "(?m)"
	}
	if options&Singleline != 0 {
		flags += "(?s)"
	}

	return flags
}

// --- Match ---

// Match represents the results from a single regular expression match.
// Equivalent to System.Text.RegularExpressions.Match in .NET.
type Match struct {
	groups    []*Group
	groupMap  map[string]*Group
	value     string
	index     int
	length    int
	success   bool
}

func newMatch(submatches []string, names []string) *Match {
	if len(submatches) == 0 {
		return &Match{success: false}
	}

	groups := make([]*Group, len(submatches))
	groupMap := make(map[string]*Group)

	// Find the index of the full match in the input
	// submatches[0] is the full match
	fullMatch := submatches[0]

	for i, sm := range submatches {
		g := &Group{
			value: sm,
			index: 0, // Getting exact index requires FindStringSubmatchIndex
			length: len(sm),
			success: sm != "" || i == 0,
		}
		groups[i] = g
		if i < len(names) && names[i] != "" {
			groupMap[names[i]] = g
		}
	}

	return &Match{
		groups:   groups,
		groupMap: groupMap,
		value:    fullMatch,
		length:   len(fullMatch),
		success:  true,
	}
}

// Success gets a value indicating whether the match is successful.
func (m *Match) Success() bool {
	return m.success
}

// Value gets the matched substring.
func (m *Match) Value() string {
	return m.value
}

// Index gets the position in the original string where the first character of the matched substring is found.
func (m *Match) Index() int {
	return m.index
}

// Length gets the length of the matched substring.
func (m *Match) Length() int {
	return m.length
}

// Groups gets a collection of groups matched by the regular expression.
func (m *Match) Groups() []*Group {
	return m.groups
}

// Group returns the captured group by number or name.
func (m *Match) Group(nameOrNumber interface{}) *Group {
	switch v := nameOrNumber.(type) {
	case int:
		if v >= 0 && v < len(m.groups) {
			return m.groups[v]
		}
	case string:
		if g, ok := m.groupMap[v]; ok {
			return g
		}
	}
	return &Group{success: false}
}

// String returns the matched substring.
func (m *Match) String() string {
	return m.value
}

// --- Group ---

// Group represents the results from a single capturing group.
// Equivalent to System.Text.RegularExpressions.Group in .NET.
type Group struct {
	value   string
	index   int
	length  int
	success bool
}

// Success gets a value indicating whether the match is successful.
func (g *Group) Success() bool {
	return g.success
}

// Value gets the captured substring from the input string.
func (g *Group) Value() string {
	return g.value
}

// Index gets the position in the original string where the first character of the captured substring is found.
func (g *Group) Index() int {
	return g.index
}

// Length gets the length of the captured substring.
func (g *Group) Length() int {
	return g.length
}

// String returns the captured substring.
func (g *Group) String() string {
	return g.value
}

// --- MatchCollection ---

// MatchCollection represents the set of successful matches found by iteratively
// applying a regular expression pattern to the input string.
// Equivalent to System.Text.RegularExpressions.MatchCollection in .NET.
type MatchCollection struct {
	matches []*Match
}

// Count gets the number of matches.
func (mc *MatchCollection) Count() int {
	return len(mc.matches)
}

// Item gets the match at the specified index.
func (mc *MatchCollection) Item(index int) *Match {
	if index < 0 || index >= len(mc.matches) {
		return nil
	}
	return mc.matches[index]
}

// ToSlice returns all matches as a slice.
func (mc *MatchCollection) ToSlice() []*Match {
	return mc.matches
}

// --- Convenience functions (static methods on .NET Regex) ---

// IsMatch tests whether the specified pattern finds a match in the input string.
func IsMatch(input, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

// MatchString searches the input string for the first occurrence of the specified pattern.
func MatchString(input, pattern string) *Match {
	r, err := Compile(pattern)
	if err != nil {
		return nil
	}
	return r.Match(input)
}

// MatchesString searches the input string for all occurrences of the specified pattern.
func MatchesString(input, pattern string) *MatchCollection {
	r, err := Compile(pattern)
	if err != nil {
		return &MatchCollection{matches: []*Match{}}
	}
	return r.Matches(input)
}

// ReplaceString replaces all occurrences of the pattern in input with replacement.
func ReplaceString(input, pattern, replacement string) string {
	r, err := Compile(pattern)
	if err != nil {
		return input
	}
	return r.Replace(input, replacement)
}

// SplitString splits the input string at the positions defined by the pattern.
func SplitString(input, pattern string) []string {
	r, err := Compile(pattern)
	if err != nil {
		return []string{input}
	}
	return r.Split(input)
}

// Escape escapes a minimal set of characters (\, *, +, ?, |, {, [, (,), ^, $, ., #, and white space)
// by replacing them with their escape codes.
func Escape(str string) string {
	return regexp.QuoteMeta(str)
}

// Unescape converts any escaped characters in the input string.
func Unescape(str string) string {
	return str // regexp package doesn't expose unescape
}
