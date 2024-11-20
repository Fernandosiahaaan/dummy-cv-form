package redis

import (
	"encoding/json"
	"fmt"
	"profiles-service/internal/model"
)

func (r *Redis) SetEducationToRedis(education *model.Education) error {
	educationJson, err := json.Marshal(education)
	if err != nil {
		return fmt.Errorf("failed convert education to json")
	}

	keyEducationInfo := fmt.Sprintf("%s:%d:%d", PrefixKeyEducations, education.ProfileCode, education.ID)
	err = r.Redis.Set(r.Ctx, keyEducationInfo, educationJson, sessionTime).Err()
	if err != nil {
		return fmt.Errorf("error save education to redis. err = %s", err.Error())
	}

	return nil
}

func (r *Redis) GetEducationFromRedis(profile_code, idEducation int64) (education *model.Education, err error) {
	keyEducationInfo := fmt.Sprintf("%s:%d:%d", PrefixKeyEducations, profile_code, idEducation)
	educationJson, err := r.Redis.Get(r.Ctx, keyEducationInfo).Result()
	if err != nil {
		return nil, fmt.Errorf("failed get education from redis. err := %v", err)
	}
	err = json.Unmarshal([]byte(educationJson), &education)
	if err != nil {
		return nil, fmt.Errorf("failed convert education from json")
	}

	return education, nil
}

func (r *Redis) GetEducationsFromRedis(profile_code int64) ([]*model.Education, error) {
	keysPattern := fmt.Sprintf("%s:%d:*", PrefixKeyEducations, profile_code)

	keys, err := r.Redis.Keys(r.Ctx, keysPattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys from redis with pattern %s. err: %v", keysPattern, err)
	}

	// MGET all of data in keys
	educationsJson, err := r.Redis.MGet(r.Ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get multiple educations from redis. err: %v", err)
	}

	var educations []*model.Education
	for _, educationJson := range educationsJson {
		if educationJson == nil {
			continue
		}
		var education model.Education
		err := json.Unmarshal([]byte(educationJson.(string)), &education)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal education json: %v", err)
		}
		educations = append(educations, &education)
	}

	return educations, nil
}

func (r *Redis) DeleteEducationFromRedis(profile_code, educationID int64) error {
	keyEducationInfo := fmt.Sprintf("%s:%d:%d", PrefixKeyEducations, profile_code, educationID)

	_, err := r.Redis.Del(r.Ctx, keyEducationInfo).Result()
	if err != nil {
		return fmt.Errorf("failed to delete education from redis. err := %v", err)
	}

	return nil
}
