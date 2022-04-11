package sqlmaker

import (
	"testing"
)

func TestBasic(t *testing.T) {
	where := Where{}
	where["username|nickname|intro|remark#like"] = "panco"
	where["id|user_id|group_id#="] = "1"
	where["username#like"] = "panco"
	where["username|nickname"] = "panco"
	where["id|user_id#>="] = "1"
	where["id|user_id#<"] = "2"
	arr := WhereOrArray(where)
	t.Log(arr)
	str := WhereOrString(where)
	t.Log(str)
}
