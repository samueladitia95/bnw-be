package utils

func EmailSubjectMapping(subject string) (email string) {
	subjectToEmailMap := map[string]string{
		// "Partnership":              "benyaminwidjojo@bwintl.com",
		// "Be Our Official Retailer": "sales@bwintl.com",
		// "Customer Service":         "customerservice@bwintl.com",
		"Partnership":              "monicapangestu23@gmail.com",
		"Be Our Official Retailer": "monicapangestudsgn@gmail.com",
		"Customer Service":         "monicapangestu23@gmail.com",
	}

	return subjectToEmailMap[subject]
}
