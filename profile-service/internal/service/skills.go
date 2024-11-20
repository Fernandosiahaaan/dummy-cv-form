package service

import (
	"profiles-service/internal/model"
)

func (s *Service) GetSkills(profile_code int64) ([]*model.Skill, error) {
	skillInfo, err := s.redis.GetSkillsFromRedis(profile_code)
	if err == nil && skillInfo != nil {
		return skillInfo, nil
	}

	existSkills, err := s.repo.GetSkillsByProfileCode(profile_code)
	if err != nil {
		return nil, err
	}

	return existSkills, nil
}

func (s *Service) CreateSkill(skill *model.Skill) (*model.Skill, error) {
	var err error
	skill.DeletedAt = nil

	skill.ID, err = s.repo.CreateNewSkill(skill)
	if err != nil {
		return nil, err
	}

	if err = s.redis.SetSkillToRedis(skill); err != nil {
		return nil, err
	}

	return skill, nil
}

func (s *Service) DeleteSkill(skillID, profile_code int64) error {
	err := s.repo.SoftDeleteSkill(skillID)
	if err != nil {
		return err
	}

	return s.redis.DeleteSkillFromRedis(profile_code, skillID)
}
