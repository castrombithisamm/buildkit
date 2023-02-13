/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type SystemTime struct {
	Year int32 `json:"Year,omitempty"`

	Month int32 `json:"Month,omitempty"`

	DayOfWeek int32 `json:"DayOfWeek,omitempty"`

	Day int32 `json:"Day,omitempty"`

	Hour int32 `json:"Hour,omitempty"`

	Minute int32 `json:"Minute,omitempty"`

	Second int32 `json:"Second,omitempty"`

	Milliseconds int32 `json:"Milliseconds,omitempty"`
}