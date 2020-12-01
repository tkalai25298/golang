package data

type IssueCertResponse struct{
	RequestID string `json:"request_id"`
	LeaseID string	`json:"lease_id"`
	Data IssueCertResponseData `json:"data"`
}

type IssueCertResponseData struct{
	Certificate string `json:"certificate"`
	PrivateKey string `json:"private_key"`
	Organization string `json:"organization"`

}