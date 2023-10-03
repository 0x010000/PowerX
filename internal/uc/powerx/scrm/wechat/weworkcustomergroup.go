package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/response"
	creq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	crsp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/response"
	"time"
)

//
// PullListWeWorkCustomerGroupRequest
//  @Description:
//  @receiver this
//  @param opt
//  @param sync
//  @return list
//  @return err
//
func (this *wechatUseCase) PullListWeWorkCustomerGroupRequest(opt *request.RequestGroupChatList, sync bool) (list []*response.GroupChat, err error) {

	if !sync {
		list = transferMapToCustomerGroups(this.getCustomerGroupFromKVByChatId())
		if list != nil {
			return list, nil
		}
	}

	reply, err := this.wework.ExternalContactGroupChat.List(this.ctx, opt)
	if err != nil {
		panic(err)
	} else {
		err = this.help.error(`scrm.wework.list.customer.group.error`, reply.ResponseWork)
	}

	if reply != nil {
		this.gLock.Add(len(reply.GroupChatList))
		for _, chat := range reply.GroupChatList {
			go func(chatID string) {
				get, _ := this.wework.ExternalContactGroupChat.Get(this.ctx, chatID, 1)
				if get.ErrCode == 0 && get.GroupChat != nil {
					this.pushCustomerGroupToKV(get.GroupChat)
					list = append(list, get.GroupChat)
				}
				this.gLock.Done()
			}(chat.ChatID)
		}
		this.gLock.Wait()
	}

	return list, err

}

// PushWoWorkCustomerTemplateRequest
//
//	@Description:
//	@receiver this
//	@param opt
//	@return *crsp.ResponseAddMessageTemplate
//	@return error
func (this *wechatUseCase) PushWoWorkCustomerTemplateRequest(opt *creq.RequestAddMsgTemplate, sendTime int64) (*crsp.ResponseAddMessageTemplate, error) {

	if sendTime > time.Now().Unix() {

		this.pushTimerMessageToKV(AppGroupCustomerMessageTimerTypeByte, sendTime, opt)

	}
	reply, err := this.wework.ExternalContactMessageTemplate.AddMsgTemplate(this.ctx, opt)
	if err != nil {
		panic(err)
	} else {
		err = this.help.error(`scrm.push.wework.customer.message.error.`, reply.ResponseWork)
	}
	return reply, err

}

//
// pushCustomerGroupToKV
//  @Description:
//  @receiver this
//  @param chatGroup
//
func (this *wechatUseCase) pushCustomerGroupToKV(chatGroup *response.GroupChat) {

	mars, _ := json.Marshal(chatGroup)
	_ = this.kv.Hset(HRedisScrmCustomerGroupKey, chatGroup.ChatID, string(mars))
	for _, member := range chatGroup.MemberList {
		meber, _ := json.Marshal(member)
		_ = this.kv.Hset(fmt.Sprintf(HRedisScrmCustomerGroupChatIDKey, chatGroup.ChatID), member.UserID, string(meber))
	}

}

//
// GetCustomerGroupFromKVByUserId
//  @Description:
//  @receiver this
//  @param chatId
//  @param userId
//  @return interface{}
//
func (this *wechatUseCase) GetCustomerGroupFromKVByUserId(chatId, userId string) interface{} {

	key := fmt.Sprintf(HRedisScrmCustomerGroupChatIDKey, chatId)
	val, _ := this.kv.Hget(key, userId)
	return val
}

//
// getCustomerGroupFromKVByChatId
//  @Description:
//  @receiver this
//  @param chatId
//  @return chat
//
func (this *wechatUseCase) getCustomerGroupFromKVByChatId(chatId ...string) (chat map[string]string) {

	chat = make(map[string]string, len(chatId))
	if len(chatId) == 0 {
		chat, _ = this.kv.Hgetall(HRedisScrmCustomerGroupKey)
	} else {
		for _, id := range chatId {
			val, _ := this.kv.Hget(HRedisScrmCustomerGroupKey, id)
			if val != `` {
				chat[id] = val
			}
		}
	}
	return chat
}

//
// transferMapToCustomerGroups
//  @Description:
//  @param chat
//  @return chats
//
func transferMapToCustomerGroups(chat map[string]string) (chats []*response.GroupChat) {

	if chat != nil {
		for _, val := range chat {
			var group *response.GroupChat
			_ = json.Unmarshal([]byte(val), &group)
			chats = append(chats, group)
		}
	}
	return chats

}

//
// GetCustomerGroupFromKVByChatId
//  @Description:
//  @receiver this
//  @param chatID
//  @param sync
//  @return chat
//
func (this *wechatUseCase) GetCustomerGroupFromKVByChatId(chatID string, sync bool) (chat *response.GroupChat) {

	chats := transferMapToCustomerGroups(this.getCustomerGroupFromKVByChatId(chatID))
	if sync || chats == nil {
		get, _ := this.wework.ExternalContactGroupChat.Get(this.ctx, chatID, 1)
		if get.GroupChat != nil {
			this.pushCustomerGroupToKV(get.GroupChat)
		}
		return get.GroupChat
	}
	return chats[0]

}
