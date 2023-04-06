package arfblocks

type IRequestModel struct {
}

type IResponseModel struct {
}

type IEndoint interface {
	Validate() error
	Verificate() error
	Handle() (any, error)
}
