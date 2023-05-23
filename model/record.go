package model

import "encoding/json"

type Record struct {
	Username  string          `json:"username,omitempty"`    // username
	Id        int             `json:"id" xorm:"pk autoincr"` // record id
	Time      string          `json:"time"`                  // record time
	OriginURL string          `json:"origin_url,omitempty"`
	PointURL  string          `json:"point_url,omitempty"`
	ArrowURL  string          `json:"arrow_url,omitempty"`
	PT        float64         `json:"pt,omitempty"`
	MT        float64         `json:"mt,omitempty"`
	TL        float64         `json:"tl,omitempty"`
	Points    json.RawMessage `json:"points,omitempty"`
}
