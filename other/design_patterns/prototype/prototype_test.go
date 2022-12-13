package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	copier := NewCopier("云打印机")
	oneNewspaper := NewNewspaper("Go是最好的编程语言", "Go语言十大优势")
	oneResume := NewResume("小明", 29, "5年码农")

	otherNewspaper := copier.copy(oneNewspaper)
	copyNewspaperMsg := make([]byte, 100)
	byteSize, _ := otherNewspaper.Read(copyNewspaperMsg)
	fmt.Println("copyNewspaperMsg:" + string(copyNewspaperMsg[:byteSize]))

	otherResume := copier.copy(oneResume)
	copyResumeMsg := make([]byte, 100)
	byteSize, _ = otherResume.Read(copyResumeMsg)
	fmt.Println("copyResumeMsg:" + string(copyResumeMsg[:byteSize]))
}
