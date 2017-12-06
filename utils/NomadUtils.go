package utils

import (
	"strings"
)

func ToNomadJobId(kind string, id string) string {
	return kind + "_" + id
}

func FromNomadJobId(nomadJobId string) (string, string) {
	kind := nomadJobId[:strings.Index(nomadJobId, "_")]
	id := nomadJobId[strings.Index(nomadJobId, "_") + 1:]

	return kind, id
}