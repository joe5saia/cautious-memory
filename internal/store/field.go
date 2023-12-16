package store

import (
	"github.com/joe5saia/cautious-memory/internal/model"
)

// GetField retrieves a single profile by its ID, including Fields
func (s *Store) GetField(id uint) (*model.Field, error) {
	var Field model.Field
	err := s.db.First(&Field, id).Error
	if err != nil {
		return nil, err
	}
	return &Field, nil
}

// GetFields retrieves all Fields
func (s *Store) GetFields() ([]model.Field, error) {
	var Fields []model.Field
	if err := s.db.Find(&Fields).Error; err != nil {
		return nil, err
	}
	return Fields, nil
}

// CreateField creates a new Field
func (s *Store) CreateField(Field *model.Field) error {
	return s.db.Create(Field).Error
}

// UpdateField updates an existing Field based on ID
func (s *Store) UpdateField(id uint, updatedField *model.Field) error {
	var Field model.Field
	if err := s.db.First(&Field, id).Error; err != nil {
		return err
	}

	return s.db.Model(&Field).Updates(updatedField).Error
}

// DeleteField deletes a Field based on ID
func (s *Store) DeleteField(id uint) error {
	return s.db.Delete(&model.Field{}, id).Error
}
