package dependencyinjection

// 1. Define an Interface for the required behavior
type DataStore interface {
	CountUsers() int
}

type UserService struct {
	DB DataStore // Depends on an interface, not a concrete type
}

// The function is now agnostic to whether it's talking to a real DB or a fake one
func (s *UserService) GetUserCount() int {
	count := s.DB.CountUsers()
	return count
}
