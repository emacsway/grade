package diff

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

func MakePatch(source, destination string) string {
	dmp := diffmatchpatch.New()
	return dmp.PatchToText(dmp.PatchMake(source, destination))
}

func ApplyPatch(source, patchRepr string) string {
	dmp := diffmatchpatch.New()
	patches := Must[[]diffmatchpatch.Patch](dmp.PatchFromText(patchRepr))
	result, _ := dmp.PatchApply(patches, source)
	return result
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
