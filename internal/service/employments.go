package service

import (
	"dummy-cv-form/internal/model"
)

func (s *Service) GetEmployment(profile_code int64) ([]*model.Employment, error) {
	employmentsInfo, err := s.redis.GetEmploymentsFromRedis(profile_code)
	if (err == nil) && (employmentsInfo != nil) {
		return employmentsInfo, nil
	}

	existEmployments, err := s.repo.GetEmploymentsByProfileCode(profile_code)
	if err != nil {
		return nil, err
	}
	return existEmployments, nil
}

func (s *Service) CreateEmployment(employment *model.Employment) (*model.Employment, error) {
	var err error

	employment.ID, err = s.repo.CreateNewEmployment(employment)
	if err != nil {
		return nil, err
	}

	if err = s.redis.SetEmploymentToRedis(employment); err != nil {
		return nil, err
	}
	return employment, nil
}

func (s *Service) DeleteEmployment(employmentID, profile_code int64) error {
	err := s.repo.SoftDeleteEmployment(employmentID)
	if err != nil {
		return err
	}

	return s.redis.DeleteEmploymentFromRedis(profile_code, employmentID)
}
