package service

import (
	"dummy-cv-form/internal/model"
	"fmt"
	"time"
)

func (s *Service) GetEducations(profileCode int64) ([]*model.Education, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	educationInfo, err := s.redis.GetEducationsFromRedis(profileCode)
	if (err == nil) && (educationInfo != nil) {
		return educationInfo, nil
	}

	existEducations, err := s.repo.GetEducationsByProfileCode(profileCode)
	if err != nil {
		return nil, err
	}

	for _, education := range existEducations {
		if err = s.redis.SetEducationToRedis(education); err != nil {
			fmt.Printf("failed set to redis data education with id %d\n", education.ID)
		}
	}

	return existEducations, nil
}

func (s *Service) CreateEducation(education *model.Education) (*model.Education, error) {
	profile, err := s.GetProfile(education.ProfileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, education.ProfileCode)
	}

	education.CreatedAt = time.Now()
	education.UpdatedAt = time.Now()
	education.ID, err = s.repo.CreateNewEducation(education)
	if err != nil {
		return nil, err
	}

	if err = s.redis.SetEducationToRedis(education); err != nil {
		fmt.Printf("failed set to redis data education with id %d\n", education.ID)
	}

	return education, nil
}

func (s *Service) DeleteEducation(educationID, profileCode int64) error {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return err

	} else if profile == nil {
		return fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	_, err = s.repo.GetEducationByID(educationID)
	if err != nil {
		return fmt.Errorf("%s%d", model.EducationErr01, educationID)
	}

	err = s.repo.SoftDeleteEducation(educationID)
	if err != nil {
		return err
	}

	return s.redis.DeleteEducationFromRedis(profileCode, educationID)
}
