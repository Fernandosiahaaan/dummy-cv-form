package redis

import (
	"encoding/json"
	"fmt"
	"profiles-service/internal/model"
)

func (r *Redis) SetEmploymentToRedis(employment *model.Employment) error {
	employmentJson, err := json.Marshal(employment)
	if err != nil {
		return fmt.Errorf("failed convert employment to json")
	}

	keyEmploymentInfo := fmt.Sprintf("%s:%d:%d", PrefixKeyEmployments, employment.ProfileCode, employment.ID)
	err = r.Redis.Set(r.Ctx, keyEmploymentInfo, employmentJson, sessionTime).Err()
	if err != nil {
		return fmt.Errorf("error save employment to redis. err = %s", err.Error())
	}

	return nil
}

func (r *Redis) GetEmploymentFromRedis(profile_code, idEmployment int64) (employment *model.Employment, err error) {
	keyEmploymentInfo := fmt.Sprintf("%s:%d:%d", PrefixKeyEmployments, profile_code, idEmployment)
	EmploymentJson, err := r.Redis.Get(r.Ctx, keyEmploymentInfo).Result()
	if err != nil {
		return nil, fmt.Errorf("failed get employment from redis. err := %v", err)
	}
	err = json.Unmarshal([]byte(EmploymentJson), &employment)
	if err != nil {
		return nil, fmt.Errorf("failed convert employment from json")
	}

	return employment, nil
}

func (r *Redis) GetEmploymentsFromRedis(profile_code int64) ([]*model.Employment, error) {
	keysPattern := fmt.Sprintf("%s:%d:*", PrefixKeyEmployments, profile_code)

	keys, err := r.Redis.Keys(r.Ctx, keysPattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys from redis with pattern %s. err: %v", keysPattern, err)
	}

	// MGET all of data in keys
	employmentsJson, err := r.Redis.MGet(r.Ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get multiple employments from redis. err: %v", err)
	}

	var employments []*model.Employment
	for _, employmentJson := range employmentsJson {
		if employmentJson == nil {
			continue
		}
		var employment model.Employment
		err := json.Unmarshal([]byte(employmentJson.(string)), &employment)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal employment json: %v", err)
		}
		employments = append(employments, &employment)
	}

	return employments, nil
}

func (r *Redis) DeleteEmploymentFromRedis(profile_code, employmentID int64) error {
	keyEmploymentInfo := fmt.Sprintf("%s:%d:%d", PrefixKeyEmployments, profile_code, employmentID)

	_, err := r.Redis.Del(r.Ctx, keyEmploymentInfo).Result()
	if err != nil {
		return fmt.Errorf("failed to delete employment from redis. err := %v", err)
	}

	return nil
}
