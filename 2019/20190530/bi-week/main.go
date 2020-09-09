package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bangwork/mygo/20190530/bi-week/config"
	"github.com/bangwork/mygo/20190530/bi-week/db"
	gorp "gopkg.in/gorp.v1"
	"log"
	"net/url"
	"os"
	"regexp"
	"time"
)

const (
	sqlDefaultSep  = " ,"
	splitPartNginx = 13
	splitPartTeams = 2
	maxInsertCount = 100
	// team_detail_data
	teamDetailDataV1 = "team_uuid, data_time, " +
		"team_name, team_scale, team_type, team_deploy_type, team_create_time, team_member_count, " +
		"team_user_access, project_user_access, wiki_user_access," +
		"team_admin_count, project_admin_count, wiki_admin_count, " +
		"normal_project_count, archive_project_count, task_count, create_task_count, completed_task_count, space_count, " +
		"page_count, create_page_count, create_draft_page_count, " +
		"all_task_count, all_page_count, " +
		"testcase_create_count, testcase_in_count, testcase_count, testcase_unfinishedplan, testcase_library_count, testcase_plan_count, testcaseall_count, " +
		"pipeline_ci_count, pipeline_lint_count, pipeline_test_count, pipeline_artifact_count, allpipeline_count, pipeline_scm_count, pipeline_scm_updatetime, pipeline_coreupdatecount, pipeline_corecommitcount, version, permitted_products, expire_time "
)

func main() {
	if err := config.LoadConfigs(""); err != nil {
		log.Fatalf("load config err, err=%v", err)
	}
	if err := db.InitPrivateDB(); err != nil {
		log.Fatalf("init db err, err=%v", err)
	}

	infos, err := loadPrivateData()
	if err != nil {
		log.Fatalf("parse log err, err=%v", err)
	}
	fmt.Println("len", len(infos))
	if err := SaveTeamData(db.PrivateDBMap, infos); err != nil {
		log.Fatalf("save private data err, err=%v", err)
	}
	fmt.Println("[done]")
}

func loadPrivateData() ([]TeamItemInfo, error) {
	templates := make([]TeamItemInfo, 0)

	logPath := config.StringOrPanic("nginx_log_path")
	file, err := os.Open(logPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r1 := regexp.MustCompile(`([^ ]*) ([^ ]*) ([^ ]*) (\[.*\]) (\".*?\") (-|[0-9]*) (-|[0-9]*) (\".*?\") (\".*?\") ([^ ]*) ([^ ]*) (\".*?\")`)
	r2 := regexp.MustCompile(`^\"month_teams=(.*)\"$`)
	//是否有下一行
	for scanner.Scan() {
		readLine := scanner.Text()
		ret := r1.FindStringSubmatch(string(readLine))
		if len(ret) >= splitPartNginx {
			data, err := url.QueryUnescape(ret[splitPartNginx-1])
			if err != nil {
				log.Println(err)
				continue
			}

			rawstr := r2.FindStringSubmatch(data)
			if len(rawstr) >= splitPartTeams {
				var teamIteamInfos []TeamItemInfo
				err = json.Unmarshal([]byte(rawstr[splitPartTeams-1]), &teamIteamInfos)
				if err != nil {
					continue
				}
				for _, item := range teamIteamInfos {
					if item.TeamName == "" || item.UUID == "" {
						continue
					}
					templates = append(templates, item)
				}
			}
		}
	}

	return templates, nil
}

func InsertTeamData(src gorp.SqlExecutor, infos ...TeamItemInfo) error {
	length := len(infos)
	if length == 0 {
		return nil
	}

	fieldsNumber := 44
	placeholds := sqlInMultiInsertValues(fieldsNumber, length)
	sql := fmt.Sprintf("INSERT IGNORE INTO %s (%s)", "monthly_team_data", teamDetailDataV1)
	sql += "VALUES %s;"

	sql = fmt.Sprintf(sql, placeholds)
	args := make([]interface{}, 0)
	for _, info := range infos {
		args = append(args, info.UUID)
		args = append(args, info.DataTime)
		args = append(args, info.TeamName)
		args = append(args, info.TeamScale)
		args = append(args, info.Type)
		args = append(args, info.DeployType)
		args = append(args, info.TeamCreateTime)
		args = append(args, info.MemberCount)
		args = append(args, info.TeamUserAccessCount)
		args = append(args, info.ProjectUserAccessCount)
		//args = append(args, info.TestcaseUserAccessCount)
		args = append(args, info.WikiUserAccessCount)
		args = append(args, info.TeamAdminCount)
		args = append(args, info.ProjectAdminCount)
		args = append(args, info.WikiAdminCount)
		args = append(args, info.NormalProjectCount)
		args = append(args, info.ArchiveProjectCount)
		args = append(args, info.TaskCount)
		args = append(args, info.CreateTaskCount)
		args = append(args, info.CompletedTaskCount)
		args = append(args, info.SpaceCount)
		args = append(args, info.PageCount)
		args = append(args, info.CreatePageCount)
		args = append(args, info.CreateDraftPageCount)
		args = append(args, info.AllTaskCount)
		args = append(args, info.AllPageCount)
		args = append(args, info.TestCaseCreateCount)
		args = append(args, info.TestCaseInCount)
		args = append(args, info.TestCaseCount)
		args = append(args, info.TestCaseUnfinishedPlanCount)
		args = append(args, info.TestCaseLibraryCount)
		args = append(args, info.TestCasePlanCount)
		args = append(args, info.AllTestCaseCount)
		args = append(args, info.PipelineCiCount)
		args = append(args, info.PipelineLintCount)
		args = append(args, info.PipelineTestCount)
		args = append(args, info.PipelineArtifactCount)
		args = append(args, info.AllPipelineCount)
		args = append(args, info.PipelineScmCount)
		args = append(args, info.PipelineScmUpdateTime)
		args = append(args, info.CoreDayTaskCount)
		args = append(args, info.CoreDayCommitCount)
		args = append(args, info.Version)
		args = append(args, info.PermittedProducts)
		args = append(args, info.ExpireTime)
	}

	_, err := src.Exec(sql, args...)
	return err
}

func SaveTeamData(src gorp.SqlExecutor, templates []TeamItemInfo) error {
	data := templates
	var err error
	for {
		if len(data) == 0 {
			return nil
		}

		size := maxInsertCount
		if len(data) < size {
			size = len(data)
		}
		group := data[:size]
		data = data[size:]
		err = InsertTeamData(src, group...)
		if err != nil {
			return err
		}
	}
}

type TeamItemInfo struct {
	UUID                        string    `db:"team_uuid" json:"team_uuid"`
	DataTime                    time.Time `db:"data_time" json:"data_time"`
	Type                        int       `db:"team_type" json:"team_type"`
	DeployType                  int       `db:"team_deploy_type" json:"team_deploy_type"`
	TeamName                    string    `db:"team_name" json:"team_name"`
	IsNewCreate                 bool      `db:"-" json:"-"`
	TeamScale                   int       `db:"team_scale" json:"team_scale"`
	TeamCreateTime              time.Time `db:"team_create_time" json:"team_create_time"`
	MemberCount                 int       `db:"team_member_count" json:"team_member_count"`
	TeamUserAccessCount         int       `db:"team_user_access" json:"team_user_access"`
	ProjectUserAccessCount      int       `db:"project_user_access" json:"project_user_access"`
	WikiUserAccessCount         int       `db:"wiki_user_access" json:"wiki_user_access"`
	PipelineUserAccessCount     int       `db:"pipeline_user_access" json:"pipeline_user_access"`
	TeamAdminCount              int       `db:"team_admin_count" json:"team_admin_count"`
	ProjectAdminCount           int       `db:"project_admin_count" json:"project_admin_count"`
	WikiAdminCount              int       `db:"wiki_admin_count" json:"wiki_admin_count"`
	NormalProjectCount          int       `db:"normal_project_count" json:"normal_project_count"`
	ArchiveProjectCount         int       `db:"archive_project_count" json:"archive_project_count"`
	TaskCount                   int       `db:"task_count" json:"task_count"`
	CreateTaskCount             int       `db:"create_task_count" json:"create_task_count"`
	CompletedTaskCount          int       `db:"completed_task_count" json:"completed_task_count"`
	SpaceCount                  int       `db:"space_count" json:"space_count"`
	PageCount                   int       `db:"page_count" json:"page_count"`
	CreatePageCount             int       `db:"create_page_count" json:"create_page_count"`
	CreateDraftPageCount        int       `db:"create_draft_page_count" json:"create_draft_page_count"`
	AllTaskCount                int       `db:"all_task_count" json:"all_task_count"`
	AllPageCount                int       `db:"all_page_count" json:"all_page_count"`
	TestCaseCreateCount         int       `db:"testcase_create_count" json:"testcase_create_count"`
	TestCaseInCount             int       `db:"testcase_in_count" json:"testcase_in_count"`
	TestCaseCount               int       `db:"testcase_count" json:"testcase_count"`
	TestCaseUnfinishedPlanCount int       `db:"testcase_unfinishedplan" json:"testcase_unfinishedplan"`
	TestCaseLibraryCount        int       `db:"testcase_library_count" json:"testcase_library_count"`
	TestCasePlanCount           int       `db:"testcase_plan_count" json:"testcase_plan_count"`
	AllTestCaseCount            int       `db:"testcaseall_count" json:"testcaseall_count"`
	PipelineCiCount             int       `db:"pipeline_ci_count" json:"pipeline_ci_count"`
	PipelineLintCount           int       `db:"pipeline_lint_count" json:"pipeline_lint_count"`
	PipelineTestCount           int       `db:"pipeline_test_count" json:"pipeline_test_count"`
	PipelineArtifactCount       int       `db:"pipeline_artifact_count" json:"pipeline_artifact_count"`
	AllPipelineCount            int       `db:"allpipeline_count" json:"allpipeline_count"`
	PipelineScmCount            int       `db:"pipeline_scm_count" json:"pipeline_scm_count"`
	PipelineScmUpdateTime       int64     `db:"pipeline_scm_updatetime" json:"pipeline_scm_updatetime"`
	CoreDayTaskCount            int       `db:"pipeline_coreupdatecount" json:"pipeline_coreupdatecount"`
	CoreDayCommitCount          int       `db:"pipeline_corecommitcount" json:"pipeline_corecommitcount"`
	Version                     string    `db:"version" json:"version"`
	PermittedProducts           int       `db:"permitted_products" json:"permitted_products"`
	ExpireTime                  time.Time `db:"expire_time" json:"expire_time"`
}

func sqlMultiColumnPlaceholders(columnCount int, count int) string {
	if columnCount <= 0 || count <= 0 {
		return ""
	}
	one := fmt.Sprintf("(%s)", appendDuplicateString("?", sqlDefaultSep, columnCount))
	return appendDuplicateString(one, sqlDefaultSep, count)
}

func sqlInMultiInsertValues(columnCount int, count int) string {
	return sqlMultiColumnPlaceholders(columnCount, count)
}

func appendDuplicateString(character, separator string, count int) string {
	if count <= 0 {
		return ""
	}
	var b bytes.Buffer
	for i := 0; i < count; i++ {
		if i > 0 {
			b.WriteString(separator)
		}
		b.WriteString(character)
	}
	return b.String()
}
