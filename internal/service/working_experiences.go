package service

import (
	"dummy-cv-form/internal/model"
)

func (s *Service) GetWorkingExperiences(profileCode int64) (*model.WorkingExperience, error) {
	workingExperiences, err := s.redis.GetWorkingExperiencesFromRedis(profileCode)
	if err == nil && workingExperiences != nil {
		return workingExperiences, nil
	}

	workingExperiences, err = s.repo.GetWorkingExperienceByProfileCode(profileCode)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWorkingExperiencesToRedis(profileCode, workingExperiences); err != nil {
		return nil, err
	}

	return workingExperiences, nil
}

func (s *Service) CreateWorkingExperience(profileCode int64, workingExperience *model.WorkingExperience) (*model.WorkingExperience, error) {
	var err error
	workingExperience.ProfileCode = profileCode

	workingExperience.ID, err = s.repo.CreateNewWorkingExperience(workingExperience)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWorkingExperiencesToRedis(profileCode, workingExperience); err != nil {
		return nil, err
	}

	return workingExperience, nil
}

func (s *Service) UpdateWorkingExperience(profileCode int64, workingExperience *model.WorkingExperience) (*model.WorkingExperience, error) {
	err := s.repo.UpdateWorkingExperienceByProfileCode(workingExperience)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWorkingExperiencesToRedis(profileCode, workingExperience); err != nil {
		return nil, err
	}

	return workingExperience, nil
}
