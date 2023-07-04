package storage

import "github.com/akhil-is-watching/goshort/types"

type Storage interface {
	CreateShort(*types.GoShortInput) (types.Goshort, error)
}
