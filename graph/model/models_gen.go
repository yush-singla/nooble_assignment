// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Audio struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	CreatorID   string   `json:"creatorId"`
	Creator     *Creator `json:"creator"`
}

type Creator struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewAudio struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	CreatorID   string `json:"CreatorId"`
}

type NewCreator struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}