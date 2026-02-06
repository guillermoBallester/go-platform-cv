package data

import _ "embed"

//go:embed skills.json
var SkillsJSON []byte

//go:embed experiences.json
var ExperiencesJSON []byte

//go:embed achievements.json
var AchievementsJSON []byte

//go:embed projects.json
var ProjectsJSON []byte
