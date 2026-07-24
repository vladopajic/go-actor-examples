package example

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type runFunc func()

var (
	registeredExamples = map[string]runFunc{}
	exampleNames       []string
)

func register(number int, name string, fn runFunc) {
	exampleNames = append(exampleNames, canonicalName(number, name))

	for _, alias := range aliases(number, name) {
		if _, exists := registeredExamples[alias]; exists {
			panic(fmt.Sprintf("example %q already registered", alias))
		}

		registeredExamples[alias] = fn
	}
}

// Run runs an example by name.
func Run(name string) bool {
	fn, exists := registeredExamples[normalizeName(name)]
	if !exists {
		return false
	}

	fn()

	return true
}

// Names returns all registered example names.
func Names() []string {
	names := append([]string(nil), exampleNames...)
	sort.Strings(names)

	return names
}

func normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.TrimSuffix(name, ".go")
	name = strings.ReplaceAll(name, "-", "_")

	return name
}

func canonicalName(number int, name string) string {
	return fmt.Sprintf("%02d_%s", number, normalizeName(name))
}

func aliases(number int, name string) []string {
	normalizedName := normalizeName(name)
	candidates := []string{
		strconv.Itoa(number),
		fmt.Sprintf("%02d", number),
		canonicalName(number, name),
		normalizedName,
	}

	aliases := make([]string, 0, len(candidates))
	seen := map[string]struct{}{}

	for _, alias := range candidates {
		if _, exists := seen[alias]; exists {
			continue
		}

		seen[alias] = struct{}{}

		aliases = append(aliases, alias)
	}

	return aliases
}
