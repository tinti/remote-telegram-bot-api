package rbot

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/streadway/amqp"
)

func SimpleServerDefault(url string, bot *tgbotapi.BotAPI) {
	SimpleServer(url, bot, func(error) {})
}

func SimpleServer(url string, bot *tgbotapi.BotAPI, errorHandler func(error)) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		errorHandler(err)
	}

	return NewErrorRemoteBot(FailedConnect, err)
	defer conn.Close()

	ch, err := conn.Channel()
	return NewErrorRemoteBot(FailedOpenChannel, err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RoutingKey, // name
		false,      // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	return NewErrorRemoteBot(FailedDeclareQueue, err)

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	return NewErrorRemoteBot(FailedOptionQoS, err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	return NewErrorRemoteBot(FailedMessageConsume, err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var n RequestMessage
			err = json.Unmarshal(d.Body, &n)
			if err != nil {
				errorHandler(err)
			}

			var r ResponseMessage
			switch n.Operation {
			case OperationMakeRequest:
				apiResponse, err := bot.MakeRequest(n.Endpoint, n.Params)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUploadFile:
				apiResponse, err := bot.UploadFile(n.Endpoint, n.Params2, n.Fieldname, n.File)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFileDirectURL:
				fileID, err := bot.GetFileDirectURL(n.FileID)

				r = ResponseMessage{}
				r.R3 = fileID
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetMe:
				user, err := bot.GetMe()

				r = ResponseMessage{}
				r.R4 = user
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationIsMessageToMe:
				yes := bot.IsMessageToMe(n.Message)

				r = ResponseMessage{}
				r.R5 = yes

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSend:
				switch n.C.Type {
				case reflect.TypeOf(tgbotapi.MessageConfig{}).String():
					m := n.C.ValueMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.ForwardConfig{}).String():
					m := n.C.ValueForwardConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.PhotoConfig{}).String():
					m := n.C.ValuePhotoConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.AudioConfig{}).String():
					m := n.C.ValueAudioConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.DocumentConfig{}).String():
					m := n.C.ValueDocumentConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.StickerConfig{}).String():
					m := n.C.ValueStickerConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.VideoConfig{}).String():
					m := n.C.ValueVideoConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.AnimationConfig{}).String():
					m := n.C.ValueAnimationConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.VideoNoteConfig{}).String():
					m := n.C.ValueVideoNoteConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.VoiceConfig{}).String():
					m := n.C.ValueVoiceConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.MediaGroupConfig{}).String():
					m := n.C.ValueMediaGroupConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.LocationConfig{}).String():
					m := n.C.ValueLocationConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.VenueConfig{}).String():
					m := n.C.ValueVenueConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.ContactConfig{}).String():
					m := n.C.ValueContactConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.GameConfig{}).String():
					m := n.C.ValueGameConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.SetGameScoreConfig{}).String():
					m := n.C.ValueSetGameScoreConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.GetGameHighScoresConfig{}).String():
					m := n.C.ValueGetGameHighScoresConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.ChatActionConfig{}).String():
					m := n.C.ValueChatActionConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.EditMessageTextConfig{}).String():
					m := n.C.ValueEditMessageTextConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.EditMessageCaptionConfig{}).String():
					m := n.C.ValueEditMessageCaptionConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.EditMessageReplyMarkupConfig{}).String():
					m := n.C.ValueEditMessageReplyMarkupConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.InvoiceConfig{}).String():
					m := n.C.ValueInvoiceConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.DeleteMessageConfig{}).String():
					m := n.C.ValueDeleteMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.PinChatMessageConfig{}).String():
					m := n.C.ValuePinChatMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.UnpinChatMessageConfig{}).String():
					m := n.C.ValueUnpinChatMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.SetChatTitleConfig{}).String():
					m := n.C.ValueSetChatTitleConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.SetChatDescriptionConfig{}).String():
					m := n.C.ValueSetChatDescriptionConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				case reflect.TypeOf(tgbotapi.DeleteChatPhotoConfig{}).String():
					m := n.C.ValueDeleteChatPhotoConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = NewConcreteError(err)
				//case reflect.TypeOf(tgbotapi.SetChatPhotoConfig{}).String():
				//	m := n.C.ValueSetChatPhotoConfig
				//	message, err := bot.Send(m)
				//
				//	r = ResponseMessage{}
				//	r.R6 = message
				//	r.R2 = NewConcreteError(err)
				default:
					r = ResponseMessage{}
					r.R2 = NewConcreteError(fmt.Errorf("not implemented"))
				}

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUserProfilePhotos:
				userProfilePhotos, err := bot.GetUserProfilePhotos(n.Config)

				r = ResponseMessage{}
				r.R7 = userProfilePhotos
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFile:
				file, err := bot.GetFile(n.Config2)

				r = ResponseMessage{}
				r.R8 = file
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdates:
				updates, err := bot.GetUpdates(n.Config3)

				r = ResponseMessage{}
				r.R9 = updates
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRemoveWebhook:
				apiResponse, err := bot.RemoveWebhook()

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetWebhook:
				apiResponse, err := bot.SetWebhook(n.Config4)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetWebhookInfo:
				webhookInfo, err := bot.GetWebhookInfo()

				r = ResponseMessage{}
				r.R10 = webhookInfo
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdatesChan:
				r = ResponseMessage{}
				r.R2 = NewConcreteError(fmt.Errorf("not implemented"))

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationListenForWebhook:
				r = ResponseMessage{}
				r.R2 = NewConcreteError(fmt.Errorf("not implemented"))

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerInlineQuery:
				apiResponse, err := bot.AnswerInlineQuery(n.Config5)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerCallbackQuery:
				apiResponse, err := bot.AnswerCallbackQuery(n.Config6)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationKickChatMember:
				apiResponse, err := bot.KickChatMember(n.Config7)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationLeaveChat:
				apiResponse, err := bot.LeaveChat(n.Config8)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChat:
				chat, err := bot.GetChat(n.Config8)

				r = ResponseMessage{}
				r.R12 = chat
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatAdministrators:
				chatMembers, err := bot.GetChatAdministrators(n.Config8)

				r = ResponseMessage{}
				r.R13 = chatMembers
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMembersCount:
				count, err := bot.GetChatMembersCount(n.Config8)

				r = ResponseMessage{}
				r.R14 = count
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMember:
				chatMember, err := bot.GetChatMember(n.Config9)

				r = ResponseMessage{}
				r.R15 = chatMember
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnbanChatMember:
				apiResponse, err := bot.UnbanChatMember(n.Config10)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRestrictChatMember:
				apiResponse, err := bot.RestrictChatMember(n.Config11)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPromoteChatMember:
				apiResponse, err := bot.PromoteChatMember(n.Config12)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetGameHighScores:
				gameHighScores, err := bot.GetGameHighScores(n.Config13)

				r = ResponseMessage{}
				r.R16 = gameHighScores
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerShippingQuery:
				apiResponse, err := bot.AnswerShippingQuery(n.Config14)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerPreCheckoutQuery:
				apiResponse, err := bot.AnswerPreCheckoutQuery(n.Config15)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteMessage:
				apiResponse, err := bot.DeleteMessage(n.Config16)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetInviteLink:
				link, err := bot.GetInviteLink(n.Config8)

				r = ResponseMessage{}
				r.R3 = link
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPinChatMessage:
				apiResponse, err := bot.PinChatMessage(n.Config17)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnpinChatMessage:
				apiResponse, err := bot.UnpinChatMessage(n.Config18)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatTitle:
				apiResponse, err := bot.SetChatTitle(n.Config19)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatDescription:
				apiResponse, err := bot.SetChatDescription(n.Config20)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatPhoto:
				apiResponse, err := bot.SetChatPhoto(n.Config21)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteChatPhoto:
				apiResponse, err := bot.DeleteChatPhoto(n.Config22)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = NewConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			default:
				r = ResponseMessage{}
				r.R2 = NewConcreteError(fmt.Errorf("not implemented"))

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			}

			response, err := json.Marshal(r)
			if err != nil {
				errorHandler(err)
			}

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			if err != nil {
				errorHandler(NewErrorRemoteBot(FailedMessagePublish, err))
			}

			d.Ack(false)
		}
	}()

	<-forever

	return nil
}
