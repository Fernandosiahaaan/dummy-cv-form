package service

import (
	"dummy-cv-form/internal/model"
	"fmt"
	"time"
)

func (s *Service) GetProfile(profileCode int64) (*model.Profile, error) {
	profileInfo, err := s.redis.GetProfileFromRedis(profileCode)
	if (err == nil) && (profileInfo != nil) {
		return profileInfo, nil
	}

	existProfile, err := s.repo.GetProfileByCode(profileCode)
	if err != nil {
		return nil, err
	} else if existProfile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	if err = s.redis.SetProfileToRedis(existProfile); err != nil {
		fmt.Printf("failed set to redis data profiles with code %d\n", existProfile.ProfileCode)
	}

	return existProfile, nil
}

func (s *Service) CreateNewProfile(profile *model.Profile) (int64, error) {
	var err error
	existProfile, _ := s.repo.GetProfileByEmail(profile.Email)
	if existProfile != nil {
		return 0, fmt.Errorf("%s%d", model.ProfileCodeErr02, existProfile.ProfileCode)
	}

	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()
	profile.ProfileCode, err = s.repo.CreateNewProfile(profile)
	if err != nil {
		return 0, err
	}

	if err = s.redis.SetProfileToRedis(profile); err != nil {
		return 0, err
	}
	return profile.ProfileCode, nil
}

func (s *Service) UpdateProfile(profileCode int64, profile *model.Profile) (*int64, error) {
	existProfile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if existProfile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	profile.UpdatedAt = time.Now()
	err = s.repo.UpdateProfileByCode(profileCode, profile)
	if err != nil {
		return nil, err
	}

	profile.ProfileCode = profileCode
	if err = s.redis.SetProfileToRedis(profile); err != nil {
		return nil, err
	}
	return &profileCode, nil
}
