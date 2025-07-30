package domain

import "time"

type IdempotencyKey struct {
	Key          string    `json:"key" db:"key"`
	RequestHash  string    `json:"requestHash" db:"request_hash"`
	ResponseBody string    `json:"responseBody" db:"response_body"`
	StatusCode   int       `json:"statusCode" db:"status_code"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}
