package offer

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestReplaceSpace(t *testing.T) {
    Convey("replaceSpace测试", t, func() {
        So(replaceSpace("hello  world! XXXX xxXX."), ShouldEqual, "hello%20%20world!%20XXXX%20xxXX.")
        So(replaceSpace(" "), ShouldEqual, "%20")
        So(replaceSpace(""), ShouldEqual, "")
    })
}
