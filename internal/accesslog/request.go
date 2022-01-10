package accesslog

type CreateRequest struct {
	ID     string `json:"id" validate:"required"`
	UserID string `json:"user_id"`
	Method string `json:"method" validate:"required,oneof=DELETE GET OPTIONS PATCH POST PUT PROPFIND DELETE HEAD COPY MKCOL"`
	URL    string `json:"url" validate:"required"`
	Data   []byte `json:"data" validate:"required"`
}
