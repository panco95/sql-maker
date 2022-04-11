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
	res := ScanWhereQuery(where)
	t.Log(res[0])
	t.Log(res[1])
	t.Log(res[2])
	t.Log(res[3])
	t.Log(res[4])
	t.Log(res[5])
}
