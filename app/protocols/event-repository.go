package protocols

import "github.com/iamseki/rua-api/domain"

type SaveEventRepository interface {
	Save(domain.EventModel) error
}
