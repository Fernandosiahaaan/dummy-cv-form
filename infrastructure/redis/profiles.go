package redis

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
)

func (r *Redis) SetProfileToRedis(profile *model.Profile) error {
	profileJson, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("failed convert profiles to json")
	}

	keyProfileInfo := fmt.Sprintf("%s:%d", PrefixKeyProfiles, profile.ProfileCode)
	err = r.Redis.Set(r.Ctx, keyProfileInfo, profileJson, sessionTime).Err()
	if err != nil {
		return fmt.Errorf("error save profile to redis. err = %s", err.Error())
	}

	return nil
}

func (r *Redis) GetProfileFromRedis(profileCode int64) (profile *model.Profile, err error) {
	keyProfileInfo := fmt.Sprintf("%s:%d", PrefixKeyProfiles, profileCode)
	profileJson, err := r.Redis.Get(r.Ctx, keyProfileInfo).Result()
	if err != nil {
		return nil, fmt.Errorf("failed get profile from redis. err := %v", err)
	}
	err = json.Unmarshal([]byte(profileJson), &profile)
	if err != nil {
		return nil, fmt.Errorf("failed convert profile from json")
	}

	return profile, nil
}
