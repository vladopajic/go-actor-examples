package example

import "testing"

func TestRegisteredAliases(t *testing.T) {
	t.Parallel()

	tests := []string{
		"1",
		"01",
		"01_producer_consumer",
		"producer_consumer",
		"01-producer-consumer",
		"01_producer_consumer.go",
		"11",
		"11_http_service",
		"http_service",
	}

	for _, name := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if _, exists := registeredExamples[normalizeName(name)]; !exists {
				t.Fatalf("example alias %q is not registered", name)
			}
		})
	}
}

func TestNamesReturnsCanonicalNames(t *testing.T) {
	t.Parallel()

	names := Names()
	if len(names) != 11 {
		t.Fatalf("expected 11 canonical names, got %d: %v", len(names), names)
	}

	if names[0] != "01_producer_consumer" {
		t.Fatalf("expected first name to be 01_producer_consumer, got %q", names[0])
	}

	if names[len(names)-1] != "11_http_service" {
		t.Fatalf("expected last name to be 11_http_service, got %q", names[len(names)-1])
	}
}
