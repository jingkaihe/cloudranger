/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token 
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type SearchQuery struct {

	ServersCount float32 `json:"serversCount,omitempty"`

	BackupCount float32 `json:"backupCount,omitempty"`

	ScheduleCount float32 `json:"scheduleCount,omitempty"`

	PolicyCount float32 `json:"policyCount,omitempty"`
}
