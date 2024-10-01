package model

type Certificate struct {
	Status string
}

type CertificateChallenge struct {
	DNSName      string
	DNSText      string
	ChallegeType string
}
