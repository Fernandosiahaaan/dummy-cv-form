package service

import (
	"dummy-cv-form/internal/model"
	"fmt"
)

func (s *Service) GetWorkingExperiences(profileCode int64) (*model.WorkingExperience, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	workingExperience, err := s.redis.GetWorkingExperiencesFromRedis(profileCode)
	if err == nil && workingExperience != nil {
		return workingExperience, nil
	}

	workingExperience, err = s.repo.GetWorkingExperienceByProfileCode(profileCode)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWorkingExperiencesToRedis(profileCode, workingExperience); err != nil {
		fmt.Printf("failed set to redis data working experiences with id %d\n", workingExperience.ID)
	}

	return workingExperience, nil
}

func (s *Service) CreateWorkingExperience(profileCode int64, workingExperience *model.WorkingExperience) (*model.WorkingExperience, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	workingExperience.ProfileCode = profileCode

	workingExperience.ID, err = s.repo.CreateNewWorkingExperience(workingExperience)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWorkingExperiencesToRedis(profileCode, workingExperience); err != nil {
		fmt.Printf("failed set to redis data working experiences with id %d\n", workingExperience.ID)
	}

	return workingExperience, nil
}

func (s *Service) UpdateWorkingExperience(profileCode int64, workingExperience *model.WorkingExperience) (*model.WorkingExperience, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	err = s.repo.UpdateWorkingExperienceByProfileCode(workingExperience)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWorkingExperiencesToRedis(profileCode, workingExperience); err != nil {
		return nil, err
	}

	return workingExperience, nil
}
