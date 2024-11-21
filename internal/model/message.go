package model

var (
	// Profile Message
	ProfileCodeErr01 string = "Profile does not exist with code "
	ProfileCodeErr02 string = "profile already created with code "

	// Profile Message
	PhotoErr01 string = "Photo URL does not exist with profile_code "

	// Employment
	EmploymentErr01 string = "Employment does not exist with id "

	// Education
	EducationErr01 string = "Education does not exist with id "

	// Skill
	SkillErr01 string = "Skill does not exist with id "

	// Others
	ErrParseJson        string = "failed parse body request"
	ErrParseProfileCode string = "failed to convert profile code"
)
