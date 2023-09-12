package utils

func EmailSubjectMapping(subject string) (email string) {
	subjectToEmailMap := map[string]string{
		"Partnership":              "benyaminwidjojo@bwintl.com",
		"Be Our Official Retailer": "sales@bwintl.com",
		"Customer Service":         "customerservice@bwintl.com",
	}

	return subjectToEmailMap[subject]
}
