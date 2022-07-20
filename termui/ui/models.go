package ui

type MessageOut struct {
	ToUserName string
	Content    string
	Type       int
}

type Message struct {
	FromUserName         string
	PlayLength           int
	RecommendInfo        []string
	Content              string
	StatusNotifyUserName string
	StatusNotifyCode     int
	Status               int
	VoiceLength          int
	ToUserName           string
	ForwardFlag          int
	AppMsgType           int
	AppInfo              AppInfo
	Url                  string
	ImgStatus            int
	MsgType              int
	ImgHeight            int
	MediaId              string
	FileName             string
	FileSize             string
	FromUserNickName     string
	ToUserNickName       string
}

func (m Message) String() string {
	from := m.FromUserNickName
	to := m.ToUserNickName
	if from == "" {
		from = m.FromUserName
	}
	if to == "" {
		to = m.ToUserName
	}
	return from + "->" + to + ":" + m.Content + "\n"
}

type AppInfo struct {
	Type  int
	AppID string
}
