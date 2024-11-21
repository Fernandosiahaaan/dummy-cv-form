package service

import (
	"dummy-cv-form/internal/model"
	"fmt"
)

func (s *Service) GetEmployments(profileCode int64) ([]*model.Employment, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	employmentsInfo, err := s.redis.GetEmploymentsFromRedis(profileCode)
	if (err == nil) && (employmentsInfo != nil) {
		return employmentsInfo, nil
	}

	existEmployments, err := s.repo.GetEmploymentsByProfileCode(profileCode)
	if err != nil {
		return nil, err
	}

	for _, employment := range existEmployments {
		if err = s.redis.SetEmploymentToRedis(employment); err != nil {
			fmt.Printf("failed set to redis data employment with id %d\n", employment.ID)
		}
	}

	return existEmployments, nil
}

func (s *Service) CreateEmployment(employment *model.Employment) (*model.Employment, error) {
	profile, err := s.GetProfile(employment.ProfileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, employment.ProfileCode)
	}

	employment.ID, err = s.repo.CreateNewEmployment(employment)
	if err != nil {
		return nil, err
	}

	if err = s.redis.SetEmploymentToRedis(employment); err != nil {
		fmt.Printf("failed set to redis data employment with id %d\n", employment.ID)
	}
	return employment, nil
}

func (s *Service) DeleteEmployment(employmentID, profile_code int64) error {
	profile, err := s.GetProfile(profile_code)
	if err != nil {
		return err

	} else if profile == nil {
		return fmt.Errorf("%s%d", model.ProfileCodeErr01, profile_code)
	}

	_, err = s.repo.GetEmploymentByID(employmentID)
	if err != nil {
		return fmt.Errorf("%s%d", model.EmploymentErr01, employmentID)
	}

	err = s.repo.SoftDeleteEmployment(employmentID)
	if err != nil {
		return err
	}

	return s.redis.DeleteEmploymentFromRedis(profile_code, employmentID)
}
