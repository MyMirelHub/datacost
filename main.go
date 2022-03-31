// Get billable usage across your account returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	type Response struct {
		Usage []struct {
			BillingPlan  string `json:"billing_plan"`
			EndDate      string `json:"end_date"`
			NumOrgs      int    `json:"num_orgs"`
			OrgName      string `json:"org_name"`
			PublicID     string `json:"public_id"`
			RatioInMonth int    `json:"ratio_in_month"`
			StartDate    string `json:"start_date"`
			Usage        struct {
				ApmHostSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"apm_host_sum"`
				ApmTraceSearchSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"apm_trace_search_sum"`
				InfraContainerSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"infra_container_sum"`
				InfraHostSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"infra_host_sum"`
				LogsIndexedSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"logs_indexed_sum"`
				ProfContainerSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"prof_container_sum"`
				ProfHostSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"prof_host_sum"`
				RumReplaySum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"rum_replay_sum"`
				SiemSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"siem_sum"`
				SyntheticsAPITestsSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"synthetics_api_tests_sum"`
				TimeseriesSum struct {
					AccountBillableUsage   int       `json:"account_billable_usage"`
					CalcDate               time.Time `json:"calc_date"`
					ElapsedUsageHours      int       `json:"elapsed_usage_hours"`
					FirstBillableUsageHour string    `json:"first_billable_usage_hour"`
					IsBillable             bool      `json:"is_billable"`
					IsTest                 bool      `json:"is_test"`
					LastBillableUsageHour  string    `json:"last_billable_usage_hour"`
					OrgBillableUsage       int       `json:"org_billable_usage"`
					PercentageInAccount    int       `json:"percentage_in_account"`
					UsageUnit              string    `json:"usage_unit"`
				} `json:"timeseries_sum"`
			} `json:"usage"`
		} `json:"usage"`
	}

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageMeteringApi.GetUsageBillableSummary(ctx, *datadog.NewGetUsageBillableSummaryOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageBillableSummary`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Can't access mystery struct so marhsalling it into json for now
	responseContent, _ := json.Marshal(resp)

	var result Response
	if err := json.Unmarshal(responseContent, &result); err != nil {
		fmt.Fprintf(os.Stderr, "Error when parsing response from `UsageMeteringApi.GetUsageBillableSummary`: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Result:\n%+v\n", result.Usage[0].Usage.ApmHostSum.AccountBillableUsage)
}
