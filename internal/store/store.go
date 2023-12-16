package store

import (
	"github.com/joe5saia/cautious-memory/internal/model"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// GetProfile retrieves a single profile by its ID, including subfields
func (s *Store) GetProfile(id uint) (*model.Profile, error) {
	var profile model.Profile
	err := s.db.Preload("Subfields").First(&profile, id).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetProfiles retrieves all profiles along with their subfields
func (s *Store) GetProfiles() ([]model.Profile, error) {
	var profiles []model.Profile
	if err := s.db.Preload("Subfields").Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

// CreateProfile creates a new profile in the database
func (s *Store) CreateProfile(profile *model.Profile) error {
	return s.db.Create(profile).Error
}

// UpdateProfile updates an existing profile in the database
func (s *Store) UpdateProfile(id uint, updatedProfile *model.Profile) error {
	var profile model.Profile
	if err := s.db.First(&profile, id).Error; err != nil {
		return err
	}

	return s.db.Model(&profile).Updates(updatedProfile).Error
}

// DeleteProfile deletes a profile from the database based on the provided ID
func (s *Store) DeleteProfile(id uint) error {
	return s.db.Delete(&model.Profile{}, id).Error
}

// GetProfile retrieves a single profile by its ID, including subfields
func (s *Store) GetSubfield(id uint) (*model.Subfield, error) {
	var subfield model.Subfield
	err := s.db.First(&subfield, id).Error
	if err != nil {
		return nil, err
	}
	return &subfield, nil
}

// GetSubfields retrieves all subfields
func (s *Store) GetSubfields() ([]model.Subfield, error) {
	var subfields []model.Subfield
	if err := s.db.Find(&subfields).Error; err != nil {
		return nil, err
	}
	return subfields, nil
}

// CreateSubfield creates a new subfield
func (s *Store) CreateSubfield(subfield *model.Subfield) error {
	return s.db.Create(subfield).Error
}

// UpdateSubfield updates an existing subfield based on ID
func (s *Store) UpdateSubfield(id uint, updatedSubfield *model.Subfield) error {
	var subfield model.Subfield
	if err := s.db.First(&subfield, id).Error; err != nil {
		return err
	}

	return s.db.Model(&subfield).Updates(updatedSubfield).Error
}

// DeleteSubfield deletes a subfield based on ID
func (s *Store) DeleteSubfield(id uint) error {
	return s.db.Delete(&model.Subfield{}, id).Error
}
