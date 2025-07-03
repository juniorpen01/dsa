package other

func LastWorkExperience(experiences []string) *string {
	if len(experiences) == 0 {
		return nil
	}
	return &experiences[len(experiences)-1]
}
