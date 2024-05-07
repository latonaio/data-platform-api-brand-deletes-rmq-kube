package requests

type Header struct {
	Brand				int     `json:"Brand"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
