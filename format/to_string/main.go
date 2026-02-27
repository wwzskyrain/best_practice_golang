package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	list := []PenaltyForNotify{
		{UserId: 11},
		{ID: 12},
	}
	for i, notify := range list {
		fmt.Printf("No.%d: p=%s\n", i, notify)
	}
}

type PenaltyForNotify struct {
	ID          int64 `json:"ID,omitempty"`
	PenaltyType int   `json:"penaltyType,omitempty"` // 1.反作弊 2.风控

	DimensionType int8  `json:"dimensionType,omitempty"`
	DimensionId   int64 `json:"dimensionId,omitempty"`

	Deduct bool `json:"deduct,omitempty"` // yes, no

	UserId      int64  `json:"userId,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	SiteID      int64  `json:"siteID,omitempty"`
	SiteName    string `json:"siteName,omitempty"`
	PackageName string `json:"packageName,omitempty"`
	CodeID      int64  `json:"codeID,omitempty"`

	Method int `json:"method,omitempty"` // 填充方式

	HitCountry         string `json:"hitCountry,omitempty"`
	RiskTwoLevelTagIds string `json:"riskTwoLevelTagIds,omitempty"` // 映射后风控2级标签
	OriginTagIds       string `json:"originTagIds,omitempty"`       // 原判罚命中标签
}

func (p PenaltyForNotify) String() string {
	jsonStr, _ := json.Marshal(p)
	return string(jsonStr)
}
