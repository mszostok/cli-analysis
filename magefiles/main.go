package main

import (
	"github.com/magefile/mage/mg"
	"github.com/samber/lo"
	"go.szostok.io/magex/deps"
	"go.szostok.io/magex/shx"

	"tools/target"
)

const (
	MdoxVersion = "" // latest
	bin         = "bin"
)

type Gen mg.Namespace

func (Gen) All() {
	mg.Deps(mg.F(deps.EnsureMdox, bin, MdoxVersion))
	target.Template("./README.tpl.md", "./README.md")

	lo.Must0(FmtDocs())
}

func FmtDocs() error {
	return shx.MustCmdf(`./bin/mdox fmt --soft-wraps README.md`).Run()
}
