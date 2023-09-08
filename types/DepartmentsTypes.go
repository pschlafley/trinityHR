package types

import "time"

type Departments struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"create_at"`
}

type CreateDepartmentRequest struct {
	Name string `json:"name"`
}

func (*Departments) NewDepartment(req *CreateDepartmentRequest, id int) *Departments {
	return &Departments{
		Id:        id,
		Name:      req.Name,
		CreatedAt: time.Now().UTC(),
	}
}
