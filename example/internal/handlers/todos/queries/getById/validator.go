package todo_queries_getById

import "errors"

func (e Endpoint) Validate() error {
	if e.RequestPayload.Id < 0 || e.RequestPayload.Id > 3 {
		return errors.New("todo Id must be in 0-3 range")
	}

	return nil
}
