package redis

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
)

func (r *Redis) SetSkillToRedis(skill *model.Skill) error {
	skillJson, err := json.Marshal(skill)
	if err != nil {
		return fmt.Errorf("failed to convert skill to json")
	}

	keySkillInfo := fmt.Sprintf("%s:%d:%d", PrefixKeySkills, skill.ProfileCode, skill.ID)
	err = r.Redis.Set(r.Ctx, keySkillInfo, skillJson, sessionTime).Err()
	if err != nil {
		return fmt.Errorf("error saving skill to redis. err = %s", err.Error())
	}

	return nil
}

func (r *Redis) GetSkillsFromRedis(profile_code int64) ([]*model.Skill, error) {
	keysPattern := fmt.Sprintf("%s:%d:*", PrefixKeySkills, profile_code)

	keys, err := r.Redis.Keys(r.Ctx, keysPattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys from redis with pattern %s. err: %v", keysPattern, err)
	}

	skillsJson, err := r.Redis.MGet(r.Ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get multiple skills from redis. err: %v", err)
	}

	var skills []*model.Skill
	for _, skillJson := range skillsJson {
		if skillJson == nil {
			continue
		}
		var skill model.Skill
		err := json.Unmarshal([]byte(skillJson.(string)), &skill)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal skill json: %v", err)
		}
		skills = append(skills, &skill)
	}

	return skills, nil
}

func (r *Redis) DeleteSkillFromRedis(profile_code, skillID int64) error {
	keySkillInfo := fmt.Sprintf("%s:%d:%d", PrefixKeySkills, profile_code, skillID)

	_, err := r.Redis.Del(r.Ctx, keySkillInfo).Result()
	if err != nil {
		return fmt.Errorf("failed to delete skill from redis. err := %v", err)
	}

	return nil
}
