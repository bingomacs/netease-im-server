package nimserversdk

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Chatroom struct {
	APPKEY    string
	APPSECRET string
}

type ChatroomResult struct {
	BaseResp
	ChatroomInfo ChatroomInfo `json:"chatroom"`
}

type ChatroomInfo struct {
	Roomid       int    `json:"roomid"`
	Valid        bool   `json:"valid"`
	Announcement string `json:"announcement"`
	Name         string `json:"name"`
	Broadcasturl string `json:"broadcasturl"`
	Ext          string `json:"ext"`
	Creator      string `json:"creator"`
}

type ChatroomDetailInfoResult struct {
	BaseResp
	ChatroomDetailInfo ChatroomDetailInfo `json:"chatroom"`
}

type ChatroomDetailInfo struct {
	Roomid          int    `json:"roomid"`
	Valid           bool   `json:"valid"`
	Muted           bool   `json:"muted"`
	Announcement    string `json:"announcement"`
	Name            string `json:"name"`
	Broadcasturl    string `json:"broadcasturl"`
	Onlineusercount int    `json:"onlineusercount"`
	Ext             string `json:"ext"`
	Creator         string `json:"creator"`
	Queuelevel      int    `json:"queuelevel"`
}

type BatchChatroomResult struct {
	BaseResp
	NoExistRooms []int64              `json:"noExistRooms"`
	FailRooms    []int64              `json:"failRooms"`
	SuccRooms    []ChatroomDetailInfo `json:"succRooms"`
}

type ToggleCloseStatResutl struct {
	BaseResp
	ChatroomInfo ChatroomInfo `json:"desc"`
}

type SetMember struct {
	Roomid int64  `json:"roomid"`
	Level  int    `json:"level"`
	Accid  string `json:"accid"`
	Type   string `json:"type"`
}

type SetMemberResult struct {
	SetMember SetMember `json:"desc"`
	BaseResp
}

// Create ...
func (chatroom *Chatroom) Create(creator string, name string, announcement string, broadcasturl string, ext string, queuelevel int) (*ChatroomResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_CREATE, url.Values{"creator": {creator}, "name": {name}, "announcement": {announcement}, "broadcasturl": {broadcasturl}, "ext": {ext}, "queuelevel": {strconv.Itoa(queuelevel)}})
	if err != nil {
		return nil, err
	}
	chatroomResult := &ChatroomResult{}
	err = json.Unmarshal(res, chatroomResult)
	if err != nil {
		return nil, err
	}
	return chatroomResult, nil

}

// Get ...
func (chatroom *Chatroom) Get(romid int64, needOnlineUserCount bool) (*ChatroomDetailInfoResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_GET, url.Values{"roomid": {strconv.FormatInt(romid, 10)}, "needOnlineUserCount": {strconv.FormatBool(needOnlineUserCount)}})
	if err != nil {
		return nil, err
	}
	chatroomDetailInfoResult := &ChatroomDetailInfoResult{}
	err = json.Unmarshal(res, chatroomDetailInfoResult)
	if err != nil {
		return nil, err
	}
	return chatroomDetailInfoResult, nil

}

// GetBatch ...
func (chatroom *Chatroom) GetBatch(roomids string, needOnlineUserCount bool) (*BatchChatroomResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_GET_BATCH, url.Values{"roomids": {roomids}, "needOnlineUserCount": {strconv.FormatBool(needOnlineUserCount)}})
	if err != nil {
		return nil, err
	}
	batchChatroomResult := &BatchChatroomResult{}
	err = json.Unmarshal(res, batchChatroomResult)
	if err != nil {
		return nil, err
	}
	return batchChatroomResult, nil

}

// Update ...
func (chatroom *Chatroom) Update(roomid int64, name string, announcement string, broadcasturl string, ext string, needNotify bool, notifyExt string, queuelevel int) (*ChatroomResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_UPDATE, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "name": {name}, "announcement": {announcement}, "broadcasturl": {broadcasturl}, "ext": {ext}, "needNotify": {strconv.FormatBool(needNotify)}, "notifyExt": {notifyExt}, "queuelevel": {strconv.Itoa(queuelevel)}})
	if err != nil {
		return nil, err
	}
	chatroomResult := &ChatroomResult{}
	err = json.Unmarshal(res, chatroomResult)
	if err != nil {
		return nil, err
	}
	return chatroomResult, nil

}

// name ...
func (chatroom *Chatroom) ToggleCloseStat(roomid int64, operator string, valid bool) (*ToggleCloseStatResutl, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_TOGGLE_CLOSE_STAT, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "operator": {operator}, "valid": {strconv.FormatBool(valid)}})
	if err != nil {
		return nil, err
	}
	toggleCloseStat := &ToggleCloseStatResutl{}
	err = json.Unmarshal(res, toggleCloseStat)
	if err != nil {
		return nil, err
	}
	return toggleCloseStat, nil

}

// SetMemberRole ...
func (chatroom *Chatroom) SetMemberRole(roomid int64, operator string, target string, opt int, optvalue bool, notifyExt string) (*SetMemberResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_SET_MEMBER_ROLE, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "operator": {operator}, "target": {target}, "opt": {strconv.Itoa(opt)}, "optvalue": {strconv.FormatBool(optvalue)}, "notifyExt": {notifyExt}})
	if err != nil {
		return nil, err
	}
	setMemberResult := &SetMemberResult{}
	err = json.Unmarshal(res, setMemberResult)
	if err != nil {
		return nil, err
	}
	return setMemberResult, nil

}
