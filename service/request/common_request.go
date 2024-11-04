package request

// CommonIDRequest 通用ID
type CommonIDRequest struct {
	ID int32 `json:"id" form:"id" uri:"id"`
}
