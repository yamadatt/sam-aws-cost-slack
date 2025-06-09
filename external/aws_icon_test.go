package external

import "testing"

func TestGetIconURL(t *testing.T) {
	url := GetIconURL("Amazon DynamoDB")
	if url == "" {
		t.Errorf("expected icon url for DynamoDB")
	}

	unknown := GetIconURL("Unknown Service")
	if unknown != "" {
		t.Errorf("expected empty url for unknown service")
	}
}
