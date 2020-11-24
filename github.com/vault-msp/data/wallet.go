package data

type IssueCertResponse struct{
	RequestID string `json:"request_id"`
	LeaseID string	`json:"lease_id"`
	Data IssueCertResponseData `json:"data"`
}

type IssueCertResponseData struct{
	Certificate string `json:"certificate`
	IssuingCA string `json:"issuing_ca"`
	PrivateKey string `json:"private_key"`
}