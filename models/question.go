package models

type Question struct {
	Id       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty" validate:"required"`
	Type     string   `json:"type,omitempty" validate:"required"`
	Role     []string `json:"role,omitempty" validate:"required"`
	Password string   `json:"password,omitempty" validate:"required"`
	Mobile   string   `json:"mobile,omitempty" validate:"required"`
	School   []string `json:"school,omitempty"`
	Academy  string   `json:"academy,omitempty" validate:"required"`
	Subject  string   `json:"subject,omitempty"`
	Grade    int16    `json:"grade,omitempty" validate:"required"`
}
