package parser

import (
	"dumper/model"
)

type Parser interface {
	Parse(params string) (chunks []model.Chunk, ok bool)
}