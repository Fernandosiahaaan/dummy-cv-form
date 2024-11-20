package service

import (
	"dummy-cv-form/internal/model"
	"errors"
	"fmt"
	"time"
)

func (s *Service) GetProfile(profile_code int64) (*model.Profile, error) {
	profileInfo, err := s.redis.GetProfileFromRedis(profile_code)
	if (err == nil) && (profileInfo != nil) {
		return profileInfo, nil
	}

	existProfile, err := s.repo.GetProfileByCode(profile_code)
	if err != nil {
		return nil, err
	}
	return existProfile, nil
}

func (s *Service) CreateNewProfile(profile *model.Profile) (int64, error) {
	existProfile, err := s.repo.GetProfileByCode(profile.ProfileCode)
	if err != nil {
		return 0, err
	}

	if existProfile != nil {
		return 0, errors.New("profile already created")
	}

	profile.ProfileCode, err = s.repo.CreateNewProfile(profile)
	if err != nil {
		return 0, err
	}

	if err = s.redis.SetProfileToRedis(profile); err != nil {
		return 0, err
	}
	return profile.ProfileCode, nil
}

func (s *Service) UpdateProfile(profile_code int64, profile *model.Profile) (*int64, error) {
	existProfile, err := s.GetProfile(profile_code)
	if err != nil {
		return nil, err
	} else if existProfile == nil {
		return nil, fmt.Errorf("not exist profile")
	}

	profile.UpdatedAt = time.Now()
	err = s.repo.UpdateProfileByCode(profile_code, profile)
	if err != nil {
		return nil, err
	}

	profile.ProfileCode = profile_code
	if err = s.redis.SetProfileToRedis(profile); err != nil {
		return nil, err
	}
	return &profile_code, nil
}
