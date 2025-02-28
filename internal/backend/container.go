package backend

type Container struct {
	ID              int    `json:"ID"`
	IPv4            string `json:"IPv4"`
	PingDuration    string `json:"PingDuration"`
	SuccessPingTime string `json:"SuccessPingTime"`
	IsSuccess       bool   `json:"-"`
}

type ContainerStat struct {
	SuccessPingTime string   `json:"SuccessPingTime"`
	PingDurations   []string `json:"PingDurations"`
}
