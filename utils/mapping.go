package utils

func EmailSubjectMapping(subject string) (email string) {
	subjectToEmailMap := map[string]string{
		// "Partnership":              "benyaminwidjojo@bwintl.com",
		// "Be Our Official Retailer": "sales@bwintl.com",
		// "Customer Service":         "customerservice@bwintl.com",
		"Partnership":              "reza@polaroisme.com",
		"Be Our Official Retailer": "rezafiansyah.putra@gmail.com",
		"Customer Service":         "samueladitia95@gmail.com",
	}

	return subjectToEmailMap[subject]
}
