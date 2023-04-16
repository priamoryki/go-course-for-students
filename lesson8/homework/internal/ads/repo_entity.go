package ads

type RepoEntityInterface interface {
	GetID() int64
	SetID(ID int64)
	HasName(name string) bool
}

type RepoEntity struct {
	ID int64
}

func (s *RepoEntity) GetID() int64 {
	return s.ID
}

func (s *RepoEntity) SetID(ID int64) {
	s.ID = ID
}

func (s *RepoEntity) HasName(_ string) bool {
	return false
}
