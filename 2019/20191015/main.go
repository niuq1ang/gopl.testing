package main

import (
	"fmt"
	// "path"
	"strings"
)

var sql = `mysql: [Warning] Using a password on the command line interface can be insecure.
+------------------------------+
| Tables_in_project            |
+------------------------------+
| batch_task                   |
| batch_task_record            |
| blog_subscribe               |
| blog_subscribe_count         |
| board                        |
| build_result                 |
| card                         |
| commit                       |
| component                    |
| component_template           |
| component_view               |
| component_view_user          |
| component_view_user_config   |
| corp_sync_department         |
| create_team_code             |
| crm_permission               |
| csm                          |
| dashboard                    |
| dashboard_member             |
| dashboard_preference         |
| department                   |
| department_member            |
| dingding_sync_relation       |
| domain_info                  |
| email_code                   |
| exchange_code                |
| field                        |
| field_access_rule            |
| field_config                 |
| field_multi_option_value     |
| field_option                 |
| field_option_config          |
| field_value                  |
| file                         |
| file_resource_view           |
| filter                       |
| gantt_chart                  |
| gantt_chart_field_value      |
| gantt_chart_personal_config  |
| gantt_data                   |
| important_field              |
| invitation                   |
| issue_type                   |
| kanban_status_sort           |
| kanban_task_sort             |
| ldap_config                  |
| ldap_config_status           |
| lint_issue                   |
| manhour                      |
| manhour_report               |
| message                      |
| notice                       |
| notice_config                |
| notice_config_sending_method |
| notice_config_user_domain    |
| notification_rule            |
| notification_scheme          |
| ones_app                     |
| ones_app_base_url            |
| organization                 |
| organization_update_stamp    |
| permission_rule              |
| phone_code                   |
| pipeline                     |
| pipeline_run                 |
| pipeline_run_artifact        |
| pipeline_run_commit          |
| pipeline_run_lint_issue      |
| pipeline_run_test_case       |
| pipeline_stage               |
| pipeline_test_case           |
| plugin                       |
| product_authorization        |
| product_trial_info           |
| program                      |
| program_role                 |
| program_role_member          |
| project                      |
| project_browse               |
| project_daily_stats          |
| project_field                |
| project_field_value          |
| project_filter               |
| project_issue_type           |
| project_member               |
| project_notice_config        |
| project_panel                |
| project_pin                  |
| project_pipeline             |
| project_plugin               |
| project_report               |
| project_sprint_field         |
| project_sprint_field_option  |
| project_sprint_status        |
| project_task_sort            |
| report_project_attribute     |
| report_project_progress      |
| report_project_target        |
| request_code                 |
| reset_code                   |
| reset_password_code          |
| resource                     |
| role                         |
| role_config                  |
| role_member                  |
| section                      |
| session_token                |
| setting                      |
| sprint                       |
| sprint_field_value           |
| sprint_pin                   |
| sprint_pipeline              |
| sprint_status_value          |
| sprint_timeline              |
| sso_verify_info              |
| support_request              |
| sync_department_config       |
| sync_user                    |
| sync_user_config             |
| tag                          |
| task                         |
| task_attendee                |
| task_attribute               |
| task_case                    |
| task_commit                  |
| task_plan_case               |
| task_preference              |
| task_related                 |
| task_richtext_message        |
| task_sort                    |
| task_status                  |
| task_status_config           |
| task_testcase_plan           |
| task_watcher                 |
| team                         |
| team_member                  |
| team_notice_config           |
| team_plugin                  |
| team_report                  |
| team_scm                     |
| template                     |
| template_attribute           |
| template_attribute_value     |
| template_category            |
| template_tag_relational      |
| template_target              |
| testcase_case                |
| testcase_case_step           |
| testcase_default_member      |
| testcase_field               |
| testcase_field_config        |
| testcase_field_value         |
| testcase_library             |
| testcase_library_pin         |
| testcase_module              |
| testcase_plan                |
| testcase_plan_assigns        |
| testcase_plan_case           |
| testcase_plan_case_step      |
| testcase_project_config      |
| testcase_report              |
| testcase_report_template     |
| third_party_config           |
| third_party_organization     |
| transition                   |
| update_stamp                 |
| user                         |
| user_access                  |
| user_filter_view             |
| user_filter_view_config      |
| user_filter_view_member      |
| user_group                   |
| user_group_member            |
| user_role                    |
| user_task_sort               |
| user_template_relational     |
| version                      |
| version_sprint               |
+------------------------------+`

func main() {
	dataSclice := getSQLResut(sql)
	fmt.Printf(dataSclice)
	// for _, data := range dataSclice {
	// 	fmt.Printf("%+v \n", data)
	// }

	// testMap := map[string]int{
	// 	"xx": 2,
	// }
	// value, ok := testMap["yy"]
	// fmt.Println(value, ok)
	// fmt.Println(strings.Replace("119.137.55.169_enterprise-reg.ones.ai/test/ones-release_20190912", "/", "_", 5))
	// fmt.Println(path.Dir("logs/119.137.55.169_enterprise-reg.ones.ai/test/ones-release_20190912.log"))
}

// func getSQLResut(data string) []string {
// 	dataSclice := strings.Split(data, "\n")
// 	dataSclice = dataSclice[5 : len(dataSclice)-1]
// 	var dataTmp string
// 	for _, data := range dataSclice {
// 		dataTmp += data
// 	}
// 	dataSclice = strings.Split(dataTmp, "|")
// 	var resut []string
// 	for _, data := range dataSclice {
// 		if data != "" {
// 			resut = append(resut, data)
// 		}
// 	}
// 	return resut
// }

func getSQLResut(data string) string {
	dataSclice := strings.Split(data, "\n")
	dataSclice = dataSclice[4 : len(dataSclice)-1]
	var dataTmp string
	for _, data := range dataSclice {
		dataTmp += data
	}
	dataSclice = strings.Split(dataTmp, "|")
	dataTmp = ""
	for _, data := range dataSclice {
		if data != "" {
			if dataTmp == "" {
				dataTmp += "drop table " + strings.TrimSpace(data)
			}
			dataTmp += ", " + strings.TrimSpace(data)
		}
	}
	return dataTmp
}
