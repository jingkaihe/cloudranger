/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token 
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type JobTaskStep struct {

	Id string `json:"id,omitempty"`

	JobTaskId string `json:"job_task_id,omitempty"`

	JobId string `json:"job_id,omitempty"`

	DependsOnId string `json:"depends_on_id,omitempty"`

	DependsOnResult string `json:"depends_on_result,omitempty"`

	Seq float32 `json:"seq,omitempty"`

	Step string `json:"step,omitempty"`

	Region string `json:"region,omitempty"`

	TargetAccountId string `json:"target_account_id,omitempty"`

	ResourceId string `json:"resource_id,omitempty"`

	ResourceName string `json:"resource_name,omitempty"`

	ResourceType string `json:"resource_type,omitempty"`

	ResourceSource string `json:"resource_source,omitempty"`

	Data string `json:"data,omitempty"`

	IsSuccess string `json:"is_success,omitempty"`

	StatusMessage string `json:"status_message,omitempty"`

	StatusCode string `json:"status_code,omitempty"`

	StartedAt string `json:"started_at,omitempty"`

	LastCheckedAt string `json:"last_checked_at,omitempty"`

	FinishedAt string `json:"finished_at,omitempty"`

	Attempts float32 `json:"attempts,omitempty"`

	CreatedAt float32 `json:"created_at,omitempty"`

	UpdatedAt float32 `json:"updated_at,omitempty"`
}
