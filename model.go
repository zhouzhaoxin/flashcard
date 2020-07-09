package main

type Card struct {
	ID    int `json:"id,string,omitempty"`
	Front string
	Back  string
	Known int
	State int // 记忆时标记是否无法继续点击上/下页
}
