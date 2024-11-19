package redis

import (
	"encoding/json"
	"fmt"
	"profiles-service/internal/model"
)

func (r *Redis) SetProfileToRedis(profiles *model.Profile) error {
	profileJson, err := json.Marshal(profiles)
	if err != nil {
		return fmt.Errorf("failed convert profiles to json")
	}

	keyProfileInfo := fmt.Sprintf("%s:%d", PrefixKeyProfiles, profiles.ProfileCode)
	err = r.Redis.Set(r.Ctx, keyProfileInfo, profileJson, model.SessionTime).Err()
	if err != nil {
		return fmt.Errorf("error save profile to redis. err = %s", err.Error())
	}
	return nil
}

func (r *Redis) GetProfileFromRedis(profileCode int64) (profile *model.Profile, err error) {
	keyProfileInfo := fmt.Sprintf("%s:%d", PrefixKeyProfiles, profileCode)
	profileJson, err := r.Redis.Get(r.Ctx, keyProfileInfo).Result()
	if err != nil {
		return nil, fmt.Errorf("failed get login info from redis. err := %v", err)
	}
	err = json.Unmarshal([]byte(profileJson), &profile)
	if err != nil {
		return nil, fmt.Errorf("failed convert data login info from json")
	}
	return profile, nil
}
