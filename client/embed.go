package client

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var dist embed.FS

func Dist() fs.FS {
	sfs, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	return sfs
}
