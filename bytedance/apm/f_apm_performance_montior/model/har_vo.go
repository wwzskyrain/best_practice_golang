package model

// HAR 文件结构定义（简化版，仅包含需要的字段）
type HAR struct {
	Log struct {
		Entries []struct {
			Request  `json:"request"`
			Response struct {
				Status int `json:"status"`
			} `json:"response"`
			Time float64 `json:"time"` // 响应时间（毫秒）
		} `json:"entries"`
	} `json:"log"`
}

type Request struct {
	URL      string `json:"url"`
	PostData struct {
		Text string `json:"text"`
	} `json:"postData"`
}

type Text struct {
	Id     string `json:"id"`
	Config Config `json:"config"`
}

type Config struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Body   Body   `json:"body"`
}

type Body struct {
	Aid         string        `json:"aid"`
	Region      string        `json:"region"`
	StartTime   string        `json:"start_time"`
	EndTime     string        `json:"end_time"`
	Granularity string        `json:"granularity"`
	IssueTags   []interface{} `json:"issue_tags"`
	DeviceId    string        `json:"device_id"`
	PageSize    int           `json:"page_size"`
	PageNo      int           `json:"page_no"`
}
