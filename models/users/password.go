package users

type Password struct {
	Old	string		`json:"current"`
	New	string		`json:"new"`
}
