package rbot

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/streadway/amqp"
)

const (
	RandomStringLength = 32
	RoutingKey         = "tgbotapi"
	DefaultTimeout     = 15 * time.Second
)

const (
	FailedConnect             = "failed to connect to RabbitMQ"
	FailedConvertBodyRequest  = "failed to convert body to request"
	FailedConvertBodyResponse = "failed to convert body to response"
	FailedDeclareQueue        = "failed to declare a queue"
	FailedOpenChannel         = "failed to open a channel"
	FailedMessagePublish      = "failed to publish a message"
	FailedMessageConsume      = "failed to register a consumer"
	FailedOptionQoS           = "failed to set QoS"
)

const (
	OperationMakeRequest            = "MakeRequest"
	OperationUploadFile             = "UploadFile"
	OperationGetFileDirectURL       = "GetFileDirectURL"
	OperationGetMe                  = "GetMe"
	OperationIsMessageToMe          = "IsMessageToMe"
	OperationSend                   = "Send"
	OperationGetUserProfilePhotos   = "GetUserProfilePhotos"
	OperationGetFile                = "GetFile"
	OperationGetUpdates             = "GetUpdates"
	OperationRemoveWebhook          = "RemoveWebhook"
	OperationSetWebhook             = "SetWebhook"
	OperationGetWebhookInfo         = "GetWebhookInfo"
	OperationGetUpdatesChan         = "GetUpdatesChan"
	OperationListenForWebhook       = "ListenForWebhook"
	OperationAnswerInlineQuery      = "AnswerInlineQuery"
	OperationAnswerCallbackQuery    = "AnswerCallbackQuery"
	OperationKickChatMember         = "KickChatMember"
	OperationLeaveChat              = "LeaveChat"
	OperationGetChat                = "GetChat"
	OperationGetChatAdministrators  = "GetChatAdministrators"
	OperationGetChatMembersCount    = "GetChatMembersCount"
	OperationGetChatMember          = "GetChatMember"
	OperationUnbanChatMember        = "UnbanChatMember"
	OperationRestrictChatMember     = "RestrictChatMember"
	OperationPromoteChatMember      = "PromoteChatMember"
	OperationGetGameHighScores      = "GetGameHighScores"
	OperationAnswerShippingQuery    = "AnswerShippingQuery"
	OperationAnswerPreCheckoutQuery = "AnswerPreCheckoutQuery"
	OperationDeleteMessage          = "DeleteMessage"
	OperationGetInviteLink          = "GetInviteLink"
	OperationPinChatMessage         = "PinChatMessage"
	OperationUnpinChatMessage       = "UnpinChatMessage"
	OperationSetChatTitle           = "SetChatTitle"
	OperationSetChatDescription     = "SetChatDescription"
	OperationSetChatPhoto           = "SetChatPhoto"
	OperationDeleteChatPhoto        = "DeleteChatPhoto"
)

type BotAPIIface interface {
	MakeRequest(endpoint string, params url.Values) (tgbotapi.APIResponse, error)
	UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (tgbotapi.APIResponse, error)
	GetFileDirectURL(fileID string) (string, error)
	GetMe() (tgbotapi.User, error)
	IsMessageToMe(message tgbotapi.Message) bool
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUserProfilePhotos(config tgbotapi.UserProfilePhotosConfig) (tgbotapi.UserProfilePhotos, error)
	GetFile(config tgbotapi.FileConfig) (tgbotapi.File, error)
	GetUpdates(config tgbotapi.UpdateConfig) ([]tgbotapi.Update, error)
	RemoveWebhook() (tgbotapi.APIResponse, error)
	SetWebhook(config tgbotapi.WebhookConfig) (tgbotapi.APIResponse, error)
	GetWebhookInfo() (tgbotapi.WebhookInfo, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	ListenForWebhook(pattern string) tgbotapi.UpdatesChannel
	AnswerInlineQuery(config tgbotapi.InlineConfig) (tgbotapi.APIResponse, error)
	AnswerCallbackQuery(config tgbotapi.CallbackConfig) (tgbotapi.APIResponse, error)
	KickChatMember(config tgbotapi.KickChatMemberConfig) (tgbotapi.APIResponse, error)
	LeaveChat(config tgbotapi.ChatConfig) (tgbotapi.APIResponse, error)
	GetChat(config tgbotapi.ChatConfig) (tgbotapi.Chat, error)
	GetChatAdministrators(config tgbotapi.ChatConfig) ([]tgbotapi.ChatMember, error)
	GetChatMembersCount(config tgbotapi.ChatConfig) (int, error)
	GetChatMember(config tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error)
	UnbanChatMember(config tgbotapi.ChatMemberConfig) (tgbotapi.APIResponse, error)
	RestrictChatMember(config tgbotapi.RestrictChatMemberConfig) (tgbotapi.APIResponse, error)
	PromoteChatMember(config tgbotapi.PromoteChatMemberConfig) (tgbotapi.APIResponse, error)
	GetGameHighScores(config tgbotapi.GetGameHighScoresConfig) ([]tgbotapi.GameHighScore, error)
	AnswerShippingQuery(config tgbotapi.ShippingConfig) (tgbotapi.APIResponse, error)
	AnswerPreCheckoutQuery(config tgbotapi.PreCheckoutConfig) (tgbotapi.APIResponse, error)
	DeleteMessage(config tgbotapi.DeleteMessageConfig) (tgbotapi.APIResponse, error)
	GetInviteLink(config tgbotapi.ChatConfig) (string, error)
	PinChatMessage(config tgbotapi.PinChatMessageConfig) (tgbotapi.APIResponse, error)
	UnpinChatMessage(config tgbotapi.UnpinChatMessageConfig) (tgbotapi.APIResponse, error)
	SetChatTitle(config tgbotapi.SetChatTitleConfig) (tgbotapi.APIResponse, error)
	SetChatDescription(config tgbotapi.SetChatDescriptionConfig) (tgbotapi.APIResponse, error)
	SetChatPhoto(config tgbotapi.SetChatPhotoConfig) (tgbotapi.APIResponse, error)
	DeleteChatPhoto(config tgbotapi.DeleteChatPhotoConfig) (tgbotapi.APIResponse, error)
}

var _ BotAPIIface = (*tgbotapi.BotAPI)(nil)
var _ BotAPIIface = (*RemoteBotAPI)(nil)

func RemoteBotDial(url string) (*RemoteBotAPI, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	rbot := new(RemoteBotAPI)
	rbot.Connection = conn
	rbot.Timeout = DefaultTimeout

	return rbot, nil
}

func RemoteBotClose(rbot *RemoteBotAPI) {
	rbot.Connection.Close()
}

type RemoteBotAPI struct {
	Connection *amqp.Connection
	Timeout    time.Duration
}

func (rbot *RemoteBotAPI) MakeRequest(endpoint string, params url.Values) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationMakeRequest,
		CorrelationId: randomString(RandomStringLength),
		Endpoint:      endpoint,
		Params:        params,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) UploadFile(endpoint string, params2 map[string]string, fieldname string, file interface{}) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationUploadFile,
		CorrelationId: randomString(RandomStringLength),
		Endpoint:      endpoint,
		Params2:       params2,
		Fieldname:     fieldname,
		File:          file,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetFileDirectURL(fileID string) (string, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result string

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetFileDirectURL,
		CorrelationId: randomString(RandomStringLength),
		FileID:        fileID,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R3, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetMe() (tgbotapi.User, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.User

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetMe,
		CorrelationId: randomString(RandomStringLength),
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R4, response.R2.ToError()
}

func (rbot *RemoteBotAPI) IsMessageToMe(message tgbotapi.Message) bool {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	ch, err := rbot.Connection.Channel()
	if err != nil {
		//return result, NewErrorRemoteBot(FailedOpenChannel, err)
		return false
	}
	defer ch.Close()

	q, err := CreateQueue(ch)
	if err != nil {
		//return result, NewErrorRemoteBot(FailedDeclareQueue, err)
		return false
	}

	msgs, err := CreateConsumeChannel(ch, q.Name)
	if err != nil {
		//return result, NewErrorRemoteBot(FailedMessageConsume, err)
		return false
	}

	requestMessage := RequestMessage{
		Operation:     OperationIsMessageToMe,
		CorrelationId: randomString(RandomStringLength),
		Message:       message,
	}

	request, err := json.Marshal(requestMessage)
	if err != nil {
		panic(err)
	}

	err = Publish(ch, &requestMessage, q.Name, request)
	if err != nil {
		//return result, NewErrorRemoteBot(FailedMessagePublish, err)
		return false
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				//return result, NewErrorRemoteBot(FailedConvertBodyResponse, err)
				return false
			}
			break
		}
	}

	return response.R5
}

func (rbot *RemoteBotAPI) Send(c tgbotapi.Chattable) (result tgbotapi.Message, err error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	cC := NewConcreteChattable(c)

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSend,
		CorrelationId: randomString(RandomStringLength),
		C:             cC,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R6, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetUserProfilePhotos(config tgbotapi.UserProfilePhotosConfig) (tgbotapi.UserProfilePhotos, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.UserProfilePhotos

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetUserProfilePhotos,
		CorrelationId: randomString(RandomStringLength),
		Config:        config,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R7, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetFile(config2 tgbotapi.FileConfig) (tgbotapi.File, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.File

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetFile,
		CorrelationId: randomString(RandomStringLength),
		Config2:       config2,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R8, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetUpdates(config3 tgbotapi.UpdateConfig) ([]tgbotapi.Update, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result []tgbotapi.Update

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdates,
		CorrelationId: randomString(RandomStringLength),
		Config3:       config3,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R9, response.R2.ToError()
}

func (rbot *RemoteBotAPI) RemoveWebhook() (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationRemoveWebhook,
		CorrelationId: randomString(RandomStringLength),
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) SetWebhook(config4 tgbotapi.WebhookConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetWebhook,
		CorrelationId: randomString(RandomStringLength),
		Config4:       config4,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetWebhookInfo() (tgbotapi.WebhookInfo, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.WebhookInfo

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetWebhookInfo,
		CorrelationId: randomString(RandomStringLength),
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R10, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetUpdatesChan(config3 tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error) {
	// TODO(tinti) not implemented
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.UpdatesChannel

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdatesChan,
		CorrelationId: randomString(RandomStringLength),
		Config3:       config3,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	result = make(tgbotapi.UpdatesChannel)
	return result, response.R2.ToError()
}

func (rbot *RemoteBotAPI) ListenForWebhook(pattern string) tgbotapi.UpdatesChannel {
	// TODO(tinti) not implemented
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.UpdatesChannel

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationListenForWebhook,
		CorrelationId: randomString(RandomStringLength),
		Pattern:       pattern,
	}
	request, err := json.Marshal(requestMessage)
	if err != nil {
		panic(err)
	}

	remoteBotErr = PublishWithTimeout(ch, &requestMessage, q.Name, request, ticker)
	if remoteBotErr != nil {
		return result
	}

	response, remoteBotErr := ConsumeWithTimeout(requestMessage.CorrelationId, msgs, ticker)
	if remoteBotErr != nil {
		return result
	}

	// TODO(tinti) remove this
	// This is just for refactoring
	_ = response
	return result
}

func (rbot *RemoteBotAPI) AnswerInlineQuery(config5 tgbotapi.InlineConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerInlineQuery,
		CorrelationId: randomString(RandomStringLength),
		Config5:       config5,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) AnswerCallbackQuery(config6 tgbotapi.CallbackConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerCallbackQuery,
		CorrelationId: randomString(RandomStringLength),
		Config6:       config6,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) KickChatMember(config7 tgbotapi.KickChatMemberConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationKickChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config7:       config7,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) LeaveChat(config8 tgbotapi.ChatConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationLeaveChat,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetChat(config8 tgbotapi.ChatConfig) (tgbotapi.Chat, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.Chat

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChat,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R12, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetChatAdministrators(config8 tgbotapi.ChatConfig) ([]tgbotapi.ChatMember, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result []tgbotapi.ChatMember

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChatAdministrators,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R13, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetChatMembersCount(config8 tgbotapi.ChatConfig) (int, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result int

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMembersCount,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R14, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetChatMember(config9 tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.ChatMember

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config9:       config9,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R15, response.R2.ToError()
}

func (rbot *RemoteBotAPI) UnbanChatMember(config10 tgbotapi.ChatMemberConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationUnbanChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config10:      config10,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) RestrictChatMember(config11 tgbotapi.RestrictChatMemberConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationRestrictChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config11:      config11,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) PromoteChatMember(config12 tgbotapi.PromoteChatMemberConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationPromoteChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config12:      config12,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetGameHighScores(config13 tgbotapi.GetGameHighScoresConfig) ([]tgbotapi.GameHighScore, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result []tgbotapi.GameHighScore

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetGameHighScores,
		CorrelationId: randomString(RandomStringLength),
		Config13:      config13,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R16, response.R2.ToError()
}

func (rbot *RemoteBotAPI) AnswerShippingQuery(config14 tgbotapi.ShippingConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerShippingQuery,
		CorrelationId: randomString(RandomStringLength),
		Config14:      config14,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) AnswerPreCheckoutQuery(config15 tgbotapi.PreCheckoutConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerPreCheckoutQuery,
		CorrelationId: randomString(RandomStringLength),
		Config15:      config15,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) DeleteMessage(config16 tgbotapi.DeleteMessageConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationDeleteMessage,
		CorrelationId: randomString(RandomStringLength),
		Config16:      config16,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) GetInviteLink(config8 tgbotapi.ChatConfig) (string, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result string

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetInviteLink,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R3, response.R2.ToError()
}

func (rbot *RemoteBotAPI) PinChatMessage(config17 tgbotapi.PinChatMessageConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationPinChatMessage,
		CorrelationId: randomString(RandomStringLength),
		Config17:      config17,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) UnpinChatMessage(config18 tgbotapi.UnpinChatMessageConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationUnpinChatMessage,
		CorrelationId: randomString(RandomStringLength),
		Config18:      config18,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) SetChatTitle(config19 tgbotapi.SetChatTitleConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetChatTitle,
		CorrelationId: randomString(RandomStringLength),
		Config19:      config19,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) SetChatDescription(config20 tgbotapi.SetChatDescriptionConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetChatDescription,
		CorrelationId: randomString(RandomStringLength),
		Config20:      config20,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) SetChatPhoto(config21 tgbotapi.SetChatPhotoConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetChatPhoto,
		CorrelationId: randomString(RandomStringLength),
		Config21:      config21,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}

func (rbot *RemoteBotAPI) DeleteChatPhoto(config22 tgbotapi.DeleteChatPhotoConfig) (tgbotapi.APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result tgbotapi.APIResponse

	ch, q, msgs, remoteBotErr := CreateRpcBase(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationDeleteChatPhoto,
		CorrelationId: randomString(RandomStringLength),
		Config22:      config22,
	}

	response, remoteBotErr := RpcWithTimeout(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.ToError()
}
