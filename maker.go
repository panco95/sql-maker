package sqlmaker

import (
	"fmt"
	"strings"

	"github.com/panco95/sql-maker/utils"
)

type Where map[string]string

func ScanWhereQuery(fields Where) []string {
	where := make([]string, 0)
	for field, value := range fields {
		fieldName := utils.ToSnakeCase(field)
		sql := ""
		method := "="
		//scan custom method
		customMethod := strings.Split(fieldName, "#")
		if len(customMethod) > 1 {
			fieldName = customMethod[0]
			method = customMethod[1]
		}
		if method == "like" {
			value = "%" + value + "%"
		}
		//scan whereOr
		whereOrFields := strings.Split(fieldName, "|")
		for index, thisField := range whereOrFields {
			tempSql := fmt.Sprintf("`%s` %s '%s'", thisField, method, value)
			if index > 0 {
				tempSql = " OR " + tempSql
			}
			sql += tempSql
		}
		where = append(where, sql)
	}
	return where
}
