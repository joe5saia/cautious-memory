package store

import (
	"github.com/joe5saia/cautious-memory/internal/model"
)

// GetSubfield retrieves a single profile by its ID, including subfields
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
