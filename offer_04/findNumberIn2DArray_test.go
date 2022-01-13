package offer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindNumberIn2DArray(t *testing.T) {
	Convey("findNumberIn2DArray测试", t, func() {
		matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
		So(findNumberIn2DArray(matrix, 5), ShouldBeTrue)
		So(findNumberIn2DArray(matrix, 55), ShouldBeFalse)
		So(findNumberIn2DArray([][]int{}, 55), ShouldBeFalse)
		So(findNumberIn2DArray([][]int{{}}, 55), ShouldBeFalse)
	})
}
