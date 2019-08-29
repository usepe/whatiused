package domain

// Projects struct
type Projects interface {
	GetByID(ID string) *Project
	All() []*Project
}
