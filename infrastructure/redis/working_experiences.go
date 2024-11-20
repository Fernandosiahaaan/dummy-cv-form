package redis

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
)

func (r *Redis) SetWorkingExperiencesToRedis(profileCode int64, workingExperiences *model.WorkingExperience) error {
	workingExperiencesJson, err := json.Marshal(workingExperiences)
	if err != nil {
		return fmt.Errorf("failed to convert working experiences to json")
	}

	keyWorkingExperiencesInfo := fmt.Sprintf("%s:%d", PrefixKeyWorkExperiences, profileCode)

	err = r.Redis.Set(r.Ctx, keyWorkingExperiencesInfo, workingExperiencesJson, sessionTime).Err()
	if err != nil {
		return fmt.Errorf("error saving working experiences to redis. err: %s", err.Error())
	}

	return nil
}

func (r *Redis) GetWorkingExperiencesFromRedis(profileCode int64) (workExp *model.WorkingExperience, err error) {
	keyWorkExpInfo := fmt.Sprintf("%s:%d", PrefixKeyWorkExperiences, profileCode)
	workExpJson, err := r.Redis.Get(r.Ctx, keyWorkExpInfo).Result()
	if err != nil {
		return nil, fmt.Errorf("failed get work experiences from redis. err := %v", err)
	}
	err = json.Unmarshal([]byte(workExpJson), &workExp)
	if err != nil {
		return nil, fmt.Errorf("failed convert work experiences from json")
	}

	return workExp, nil
}
