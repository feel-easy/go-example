package ui

import (
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	CurMark  = "(bg-red)"
	PageSize = 45
)

type Layout struct {
	chatBox         *widgets.Paragraph //聊天窗口
	msgInBox        *widgets.Paragraph //消息窗口
	editBox         *widgets.Paragraph // 输入框
	userNickListBox *widgets.List
	userNickList    []string
	userIDList      []string
	curUserIndex    int
	masterName      string // 主人的名字
	masterID        string //主人的id
	currentMsgCount int
	maxMsgCount     int
	userIn          chan []string   // 用户的刷新
	msgIn           chan Message    // 消息刷新
	msgOut          chan MessageOut //  消息输出
	closeChan       chan int
	autoReply       chan int
	showUserList    []string
	userCount       int //用户总数，这里有重复,后面会修改
	pageCount       int // page总数。
	userCur         int // 当前page中所选中的用户
	curPage         int // 当前所在页
	pageSize        int // page的size默认是50
	curUserId       string
	userMap         map[string]string
	logger          *log.Logger
}

func NewLayout(myName, myID string, msgIn chan Message, msgOut chan MessageOut, closeChan, autoReply chan int, logger *log.Logger) *Layout {
	//用户列表框
	userMap := make(map[string]string)
	userNickListBox := widgets.NewList()
	userNickListBox.Title = "用户列表"
	userNickListBox.TextStyle = ui.NewStyle(ui.ColorMagenta)
	userNickListBox.Block.BorderStyle.Fg = ui.ColorGreen

	chatBox := widgets.NewParagraph()
	chatBox.Title = "hole"
	chatBox.TextStyle.Fg = ui.ColorRed
	chatBox.BorderStyle.Fg = ui.ColorMagenta

	msgInBox := widgets.NewParagraph()
	msgInBox.TextStyle.Fg = ui.ColorWhite
	msgInBox.Title = "消息窗"
	msgInBox.BorderStyle.Fg = ui.ColorCyan
	msgInBox.TextStyle.Fg = ui.ColorYellow

	editBox := widgets.NewParagraph()
	editBox.TextStyle.Fg = ui.ColorWhite
	editBox.Title = "输入框"
	editBox.BorderStyle.Fg = ui.ColorCyan

	return &Layout{
		userCur:         0,
		curPage:         0,
		msgInBox:        msgInBox,
		userNickListBox: userNickListBox,
		chatBox:         chatBox,
		editBox:         editBox,
		msgIn:           msgIn,
		msgOut:          msgOut,
		closeChan:       closeChan,
		currentMsgCount: 0,
		maxMsgCount:     18,
		pageSize:        PageSize,
		curUserIndex:    0,
		userMap:         userMap,
		masterID:        myID,
		masterName:      myName,
		logger:          logger,
	}
}

func (l *Layout) Init() {
	//	chinese := false
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	width, height := ui.TerminalDimensions()
	l.chatBox.SetRect(0, 0, width*2/10, height)

	l.userNickListBox.SetRect(width*2/10, 0, width*4/10, height*8/10)
	l.msgInBox.SetRect(width*4/10, 0, width, height*8/10)

	l.editBox.SetRect(width*2/10, height*8/10, width, height)

	go l.displayMsgIn()

	// 注册各个组件
	ui.Render(l.msgInBox, l.chatBox, l.editBox, l.userNickListBox)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Enter>":
			appendToPar(l.editBox, l.masterName+"->"+DelBgColor(l.chatBox.Text)+":"+l.editBox.Text+"\n")
			if l.editBox.Text != "" {
				// l.SendText(l.editBox.Text)
			}
			resetPar(l.editBox)
		case "<Space>":
			appendToPar(l.editBox, " ")
		case "<Backspace>":
			backToPar(l.editBox)
		default:
			if e.Type == ui.KeyboardEvent {
				appendToPar(l.editBox, e.ID)
			}
		}
	}
}

func (l *Layout) displayMsgIn() {
	var (
		msg Message
	)
	for {
		select {
		case msg = <-l.msgIn:
			text := msg.String()
			appendToPar(l.msgInBox, text)
			if msg.FromUserName == l.userIDList[l.curPage*PageSize+l.userCur] {
				appendToPar(l.chatBox, text)
			}
		case <-l.closeChan:
			break
		}
	}
}

func (l *Layout) PrevUser() {
	if l.userCur-1 < 0 { //如果是第一行
		if l.curPage > 0 { //如果不是第一页
			l.userCur = PageSize - 1
			l.curPage-- //到上一页
			//刷新一下显示的内容
			l.showUserList = l.userNickList[l.curPage*l.pageSize : l.curPage*l.pageSize+l.pageSize]
		} else {
			//如果是第一页
			//跳转到最后一页
			l.userCur = (l.userCount % l.pageSize) - 1
			if l.userCur < 0 {
				l.userCur = l.pageSize - 1
			}
			l.curPage = l.pageCount - 1
			l.showUserList = l.userNickList[l.curPage*l.pageSize : l.userCount]
		}
		l.showUserList[l.userCur] = AddBgColor(l.showUserList[l.userCur])
		l.userNickListBox.Rows = l.showUserList
	} else { //不是第一行，则删掉前面一行的信息，更新上一个的信息。
		l.userNickListBox.Rows[l.userCur] = DelBgColor(l.userNickListBox.Rows[l.userCur])
		l.userCur--
		l.userNickListBox.Rows[l.userCur] = AddBgColor(l.userNickListBox.Rows[l.userCur])
	}
	l.chatBox.Text = DelBgColor(l.showUserList[l.userCur])
	ui.Render(l.userNickListBox, l.chatBox)
}

func (l *Layout) NextUser() {
	if l.userCur+1 >= l.pageSize || l.userCur+1 >= len(l.showUserList) { //跳出了对应的下标
		l.userNickListBox.Rows[l.userCur] = DelBgColor(l.userNickListBox.Rows[l.userCur])
		l.userCur = 0
		l.userNickListBox.Rows[l.userCur] = AddBgColor(l.userNickListBox.Rows[l.userCur])
		if l.curPage+1 >= l.pageCount { //当前页是最后一页了
			l.curPage = 0
		} else {
			l.curPage++
		}
		if l.curPage == l.pageCount-1 { //最后一页，判断情况
			l.showUserList = l.userNickList[l.curPage*l.pageSize : l.userCount]
		} else {
			l.showUserList = l.userNickList[l.curPage*l.pageSize : l.curPage*l.pageSize+l.pageSize]
		}
		//设定第一行是背景色
		l.showUserList[0] = AddBgColor(l.showUserList[0])
		l.userNickListBox.Rows = l.showUserList
	} else {
		l.userNickListBox.Rows[l.userCur] = DelBgColor(l.userNickListBox.Rows[l.userCur])
		l.userCur++
		l.userNickListBox.Rows[l.userCur] = AddBgColor(l.userNickListBox.Rows[l.userCur])
	}
	l.chatBox.Text = DelBgColor(l.userNickListBox.Rows[l.userCur])
	ui.Render(l.userNickListBox, l.chatBox)
}

func (l *Layout) SendText(text string) {
	msg := MessageOut{}
	msg.Content = text
	msg.ToUserName = l.userIDList[l.curPage*PageSize+l.userCur]
	//appendToPar(l.msgInBox, fmt.Sprintf(text))

	l.msgOut <- msg
}

func AddBgColor(msg string) string {
	if strings.HasPrefix(msg, "[") {
		return msg
	}
	return "[" + msg + "]" + CurMark
}

func DelBgColor(msg string) string {
	if !strings.HasPrefix(msg, "[") {
		return msg
	}
	return msg[1 : len(msg)-9]
}

func appendToPar(p *widgets.Paragraph, k string) {
	if strings.Count(p.Text, "\n") >= 20 {
		p.Text = ""
	}
	p.Text += k
	ui.Render(p)
}

func backToPar(p *widgets.Paragraph) {
	if len(p.Text) == 0 {
		return
	}
	tmp := []rune(p.Text)
	p.Text = string(tmp[:len(tmp)-1])
	ui.Render(p)
}

func resetPar(p *widgets.Paragraph) {
	p.Text = ""
	ui.Render(p)
}

func setPar(p *widgets.Paragraph) {
	ui.Render(p)
}
