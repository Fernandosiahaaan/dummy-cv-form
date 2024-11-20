package service

import (
	"profiles-service/internal/model"
)

func (s *Service) GetEducation(profile_code int64) ([]*model.Education, error) {
	educationInfo, err := s.redis.GetEducationsFromRedis(profile_code)
	if (err == nil) && (educationInfo != nil) {
		return educationInfo, nil
	}

	existEducations, err := s.repo.GetEducationsByProfileCode(profile_code)
	if err != nil {
		return nil, err
	}
	return existEducations, nil
}

func (s *Service) CreateEducation(education *model.Education) (*model.Education, error) {
	var err error

	education.ID, err = s.repo.CreateNewEducation(education)
	if err != nil {
		return nil, err
	}

	if err = s.redis.SetEducationToRedis(education); err != nil {
		return nil, err
	}
	return education, nil
}

func (s *Service) DeleteEducation(educationID, profile_code int64) error {
	err := s.repo.SoftDeleteEducation(educationID)
	if err != nil {
		return err
	}

	return s.redis.DeleteEducationFromRedis(profile_code, educationID)
}
