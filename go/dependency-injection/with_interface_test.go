package dependencyinjection

import "testing"

// A simple fake struct used only for testing
type MockDatabase struct{}

// Satisfies the DataStore interface
func (m *MockDatabase) CountUsers() int {
	return 42 // We control the exact result without network calls
}

func TestGetUserCount(t *testing.T) {
	mockDB := &MockDatabase{}           // Create the mock
	service := &UserService{DB: mockDB} // Inject the mock

	count := service.GetUserCount() // Call the real logic with fake data

	if count != 42 {
		t.Errorf("Expected 42 users, got %d", count)
	}
	// This test runs instantly and reliably because no network calls were made.
}
