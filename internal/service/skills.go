package service

import (
	"dummy-cv-form/internal/model"
	"fmt"
)

func (s *Service) GetSkills(profileCode int64) ([]*model.Skill, error) {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	skillInfo, err := s.redis.GetSkillsFromRedis(profileCode)
	if err == nil && skillInfo != nil {
		return skillInfo, nil
	}

	existSkills, err := s.repo.GetSkillsByProfileCode(profileCode)
	if err != nil {
		return nil, err
	}

	for _, skill := range existSkills {
		if err = s.redis.SetSkillToRedis(skill); err != nil {
			fmt.Printf("failed set to redis data skill with id %d\n", skill.ID)
		}
	}

	return existSkills, nil
}

func (s *Service) CreateSkill(skill *model.Skill) (*model.Skill, error) {
	profile, err := s.GetProfile(skill.ProfileCode)
	if err != nil {
		return nil, err

	} else if profile == nil {
		return nil, fmt.Errorf("%s%d", model.ProfileCodeErr01, skill.ProfileCode)
	}

	skill.ID, err = s.repo.CreateNewSkill(skill)
	if err != nil {
		return nil, err
	}

	if err = s.redis.SetSkillToRedis(skill); err != nil {
		fmt.Printf("failed set to redis data skill with id %d\n", skill.ID)
	}

	return skill, nil
}

func (s *Service) DeleteSkill(skillID, profileCode int64) error {
	profile, err := s.GetProfile(profileCode)
	if err != nil {
		return err

	} else if profile == nil {
		return fmt.Errorf("%s%d", model.ProfileCodeErr01, profileCode)
	}

	_, err = s.repo.GetSkillByID(skillID)
	if err != nil {
		return fmt.Errorf("%s%d", model.SkillErr01, skillID)
	}

	err = s.repo.SoftDeleteSkill(skillID)
	if err != nil {
		return err
	}

	return s.redis.DeleteSkillFromRedis(profileCode, skillID)
}
