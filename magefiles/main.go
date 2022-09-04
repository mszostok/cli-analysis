package main

import (
	"github.com/magefile/mage/mg"

	"tools/target"
)

type Gen mg.Namespace

func (Gen) All() {
	target.Template("./README.tpl.md", "./README.md")
}
