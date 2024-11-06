package request

// CommonIDRequest 通用ID
type CommonIDRequest struct {
	ID int32 `json:"id" form:"id" uri:"id"`
}

type Paginate struct {
	Page  int `json:"page,omitempty" form:"page" uri:"page"`
	Limit int `json:"limit,omitempty" form:"limit" uri:"limit"`
}

func (p *Paginate) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Paginate) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}
