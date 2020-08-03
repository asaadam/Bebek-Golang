package user

type UUID [16]byte

type user struct {
	Id       UUID   `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
	Kios     string `json:"kios"`
}
