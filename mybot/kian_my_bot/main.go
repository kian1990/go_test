package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const telegramBotTokenEnv = "KIAN_MY_BOT"

var pendingVerifications = make(map[int64]string) // 存储用户的验证问题和答案

func generateMathQuestion() (string, string) {
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(10) + 1 // 1~10
	num2 := rand.Intn(10) + 1
	answer := strconv.Itoa(num1 + num2)
	question := fmt.Sprintf("%d + %d =", num1, num2)
	return question, answer
}

func main() {
	botToken := os.Getenv(telegramBotTokenEnv)
	if botToken == "" {
		log.Fatal("请设置环境变量 KIAN_MY_BOT")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("无法创建机器人: ", err)
	}

	bot.Debug = false
	log.Printf("已授权机器人账号: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	var wg sync.WaitGroup

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// 处理新成员加入群组
		if len(update.Message.NewChatMembers) > 0 {
			for _, newMember := range update.Message.NewChatMembers {
				question, answer := generateMathQuestion()
				pendingVerifications[newMember.ID] = answer // 存储正确答案

				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					fmt.Sprintf("@%s 请回答以下数学问题以验证身份: \n\n%s", newMember.UserName, question))
				bot.Send(msg)
			}
			continue
		}

		// 处理用户发送的消息
		userID := update.Message.From.ID
		if expectedAnswer, exists := pendingVerifications[userID]; exists {
			userAnswer := strings.TrimSpace(update.Message.Text)
			if userAnswer == expectedAnswer {
				// 验证成功
				delete(pendingVerifications, userID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("✅ 验证通过, 欢迎 @%s", update.Message.From.UserName))
				bot.Send(msg)
			} else {
				// 验证失败, 踢出用户
				kick := tgbotapi.BanChatMemberConfig{
					ChatMemberConfig: tgbotapi.ChatMemberConfig{
						ChatID: update.Message.Chat.ID,
						UserID: userID,
					},
				}
				bot.Request(kick)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("❌ 验证失败, @%s 已被移除", update.Message.From.UserName))
				bot.Send(msg)
			}
		}
	}

	wg.Wait()
}