package organization

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"gopkg.in/gorp.v1"
)

func MapTeamNameInOrganizationByTeamUUID(src gorp.SqlExecutor, teamUUID string) (map[string]string, error) {
	teamNameMap := make(map[string]string)
	sql := "SELECT uuid AS _1, name AS _2 FROM team WHERE org_uuid=(SELECT org_uuid FROM team WHERE uuid=? and status=?);"
	args, _ := utils.BuildSqlArgs(teamUUID, 1)
	result := make([]*utils.Tuple2_String_String, 0)
	if _, err := src.Select(&result, sql, args...); err != nil {
		return nil, errors.Sql(err)
	}
	for _, data := range result {
		teamNameMap[data.Ele_1] = data.Ele_2
	}
	return teamNameMap, nil
}
