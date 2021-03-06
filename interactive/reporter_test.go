package interactive

import (
	"reflect"
	"testing"

	"github.com/viktomas/godu/core"
)

func TestReportTree(t *testing.T) {
	testTree := &core.File{"b", 180, true, []*core.File{
		&core.File{"c", 100, false, []*core.File{}},
		&core.File{"d", 80, true, []*core.File{
			&core.File{"e", 50, false, []*core.File{}},
			&core.File{"f", 30, false, []*core.File{}},
		}},
	}}
	expected := []string{"100B\tc", "80B\td/"}
	testTreeAgainstOutput(testTree, expected, t)
}

func TestPrintsEmptyDir(t *testing.T) {
	testTree := &core.File{"", 50, true, []*core.File{
		&core.File{"a", 50, true, []*core.File{}},
	}}
	expected := []string{"50B\ta/"}
	testTreeAgainstOutput(testTree, expected, t)
}

func TestReportingUnits(t *testing.T) {
	testTree := &core.File{"X", 0, true, []*core.File{
		&core.File{"B", 1 << 0, false, []*core.File{}},
		&core.File{"K", 1 << 10, false, []*core.File{}},
		&core.File{"M", 1048576, false, []*core.File{}},
		&core.File{"G", 1073741824, false, []*core.File{}},
		&core.File{"T", 1099511627776, false, []*core.File{}},
		&core.File{"P", 1125899906842624, false, []*core.File{}},
	}}
	ex := []string{
		"1B\tB",
		"1K\tK",
		"1M\tM",
		"1G\tG",
		"1T\tT",
		"1P\tP",
	}
	testTreeAgainstOutput(testTree, ex, t)
}

func testTreeAgainstOutput(testTree *core.File, expected []string, t *testing.T) {
	result := ReportTree(testTree)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected:\n%sbut got:\n%s", expected, result)
	}
}
