package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// 全局常量，避免每次都读取环境变量
const (
	geminiAPIKeyEnv   = "GEMINI_API_KEY"
	telegramBotTokenEnv = "MONSTER_K_BOT"
)

// HTTP 客户端实例，复用连接池
var client = &http.Client{
	Timeout: 10 * time.Second, // 设置请求超时
}

// 创意内容池
var creativeReplies = []string{
	"生活就像一盒巧克力，你永远不知道下一颗是什么味道。",
	"梦想是我们前进的动力，就像汽车需要油。",
	"世界是一本书，不旅行的人只读了其中的一页。",
	"失败并不可怕，放弃才是真正的失败。",
	"生活不止眼前的苟且，还有诗和远方。",
	"不管多么坎坷，前方的路总会迎来曙光。",
	"每个现在，都包含着未来的种子。",
	"记住你为什么开始，这样就不会轻易放弃。",
	"真正的朋友是那个能让你成为更好的自己的人。",
	"年轻的心，永远不怕挑战。",
	"人的一生不过是两次见面，一次来，一次走。",
	"只要你敢做，世界就会给你让路。",
	"世界上最远的距离是你站在我面前，却看不见我的心。",
	"不要停下，你的努力正在改变世界。",
	"失败只是成功的一部分，它从未是终点。",
	"人生没有重来，只有从头开始。",
	"每一朵云都能遮住阳光，但阳光始终会破云而出。",
	"行动是治愈恐惧的良药，而犹豫将继续滋养恐惧。",
	"痛苦是暂时的，放弃是永远的。",
	"无论发生什么，永远不要忘记自己的初心。",
	"追求梦想的路上，最大的敌人是你自己。",
	"任何时候都不晚，最怕的是没有行动。",
	"看远一点，生命的意义就在于不断挑战自己。",
	"那些看似不起眼的努力，最终会改变你的人生。",
	"想得多不如做得多，行动比空想更能带来改变。",
	"有时候，人生就是一次不断修正的旅程。",
	"不要急于给自己下定义，人生充满了无限可能。",
	"你越害怕什么，它就越逼近你。",
	"你不能改变别人，但你可以改变自己。",
	"最大的风险就是不冒任何风险。",
	"不敢尝试的人，永远也不会知道成功的滋味。",
	"当你足够努力，幸运也会来找你。",
	"心态决定命运，努力改变未来。",
	"人活在世上，不是为了获得更多，而是为了能给予更多。",
	"不在乎做过什么，而在乎能做到什么。",
	"生活的美好，不在于得到，而在于发现。",
	"生活不会因为你的等待而停下脚步。",
	"成功不是终点，失败也不是终结，最重要的是继续前进。",
	"行走在世界的每一角，才会发现无限的可能。",
	"每一天都是一个全新的开始，勇敢追求梦想。",
	"只有不断改变自己，才能遇见更好的未来。",
	"不要轻易放弃，未来的你会感谢现在努力的自己。",
	"昨天的伤痛，今天的勇气，明天的希望。",
	"时间会带走一切，但真正属于你的会留下。",
	"有时候，路不一定要走很远，但要走得坚定。",
	"痛苦是短暂的，改变是永恒的。",
	"你能做的事情，从来都不止你认为的那么少。",
	"无论多困难，生活都能从容面对。",
	"只要心中有光，黑暗再深也不怕。",
	"成功是一种习惯，而不是偶然。",
	"每一次经历都是一次成长。",
	"希望的火种永远不能熄灭，心中的信念是最强的力量。",
	"这世上没有什么不可能，除非你自己认定它。",
	"只要心存希望，世界总会给你回报。",
	"这个世界最美的事，就是你的微笑。",
	"未来属于那些相信自己梦想之美的人。",
	"努力不一定成功，但放弃一定失败。",
	"对未来的期待，就是今天的动力。",
	"生活永远不会简单，幸福是自己创造的。",
	"时间是最公正的，它不偏向任何人。",
	"不经历风雨，怎能见彩虹。",
	"人的潜力是无限的，只有你敢于挑战自己。",
	"走得最远的人，往往是那些坚持到最后的人。",
	"脚步再小，也比停下要强。",
	"不怕千万人阻挡，只怕自己投降。",
	"不要为难自己，活得轻松一些。",
	"成功的背后是无数的坚持和努力。",
	"把每一天当作新的开始，迎接未知的挑战。",
	"再小的努力，也能积累成惊人的成就。",
	"从不放弃是成功的秘诀。",
	"永远不要低估自己，未来属于勇敢的人。",
	"生活不一定是你想的那样，但你可以选择如何面对它。",
	"今天的你比昨天的你更接近梦想。",
	"心中有阳光，世界就是明亮的。",
	"永远不要停止追求美好生活的脚步。",
	"困难是上帝赐予你的一次机会，能让你变得更强大。",
	"所有的梦想都值得去追求，前提是你要勇敢。",
	"不论目标多远，迈出的每一步都会让你离它更近。",
	"勇敢并不是没有恐惧，而是在恐惧中前行。",
	"生命中最重要的不是你做了什么，而是你做了多少。",
	"行动才是最好的答案，想得再多都不如付出。",
	"未来属于那些努力追求自己梦想的人。",
	"生活充满着不确定性，最重要的是保持乐观。",
	"让自己的每一天都充满意义。",
	"如果你想做大事，必须从小事做起。",
	"每一个明天，都是你改变自己的一天。",
	"相信自己，世界就会为你让路。",
	"梦想没有边界，成功没有限制。",
	"不要放弃任何可能的机会，改变就在下一刻。",
	"人生不会总是完美，但每一天都是新的开始。",
	"跨越每一道坎，最终会遇到更宽广的天地。",
	"只有不断前行，人生的目标才能渐渐清晰。",
	"每一次的尝试，都会带来一次成长。",
	"人生的精彩，不在于结果，而在于过程。",
	"每个人的潜力都没有极限，只有不敢去尝试的人。",
	"生活的意义，不是等待，而是不断去追寻。",
	"没有做不到的，只有不敢做的。",
	"时光不待人，只有不断奋斗才能不负韶华。",
	"无论多难的路，只要一步步走下去，总会见到希望。",
	"从不言败的人，才会有最终的成功。",
	"生活可以苦，但永远不能绝望。",
	"所有的伤痛终究会过去，而我们依然是那个坚强的自己。",
	"做自己想做的事，成就自己想成就的生活。",
	"生命中最重要的事就是忠于自己的梦想。",
}

// 随机返回一条创意内容
func getRandomCreativeReply() string {
	rand.Seed(time.Now().UnixNano())
	return creativeReplies[rand.Intn(len(creativeReplies))]
}

// 发送请求到 Gemini API
func callGeminiAPI(prompt string) (string, error) {
	// 获取 GEMINI_API_KEY
	geminiAPIKey := os.Getenv(geminiAPIKeyEnv)
	if geminiAPIKey == "" {
		log.Fatal("请设置环境变量 GEMINI_API_KEY")
	}

	// API 地址
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s", geminiAPIKey)

	// 构造请求体
	requestBody, err := json.Marshal(map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{{"text": prompt}}},
		},
	})
	if err != nil {
		return "", err
	}

	// 使用 NewRequest 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送 HTTP 请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败, 状态码: %d", resp.StatusCode)
	}

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析 JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// 提取返回的文本内容
	if contents, ok := result["candidates"].([]interface{}); ok {
		if len(contents) > 0 {
			if candidate, ok := contents[0].(map[string]interface{}); ok {
				if content, ok := candidate["content"].(map[string]interface{}); ok {
					if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
						if part, ok := parts[0].(map[string]interface{}); ok {
							if text, ok := part["text"].(string); ok {
								return text, nil
							}
						}
					}
				}
			}
		}
	}

	return "Gemini 没有返回有效的回答", nil
}

func main() {
	// 获取 Telegram 机器人 API Token
	botToken := os.Getenv(telegramBotTokenEnv)
	if botToken == "" {
		log.Fatal("请设置环境变量 MONSTER_K_BOT")
	}

	// 创建机器人实例
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("无法创建机器人: ", err)
	}

	// 关闭 Debug 模式提高性能
	bot.Debug = false
	log.Printf("已授权机器人账号: %s", bot.Self.UserName)

	// 注册机器人命令
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "开始"},
		{Command: "help", Description: "帮助"},
		{Command: "ai", Description: "人工智能"},
	}

	// 设置命令
	_, err = bot.Request(tgbotapi.NewSetMyCommands(commands...))
	if err != nil {
		log.Fatalf("设置机器人命令失败: %v", err)
	}

	// 创建更新通道
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	var wg sync.WaitGroup

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// 记录收到的消息
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var replyText string
		wg.Add(1)
		go func(update tgbotapi.Update) {
			defer wg.Done()
			// 处理不同命令
			switch update.Message.Command() {
			case "start":
				replyText = "人工智能机器人"
			case "help":
				replyText = "可用命令: \n/start - 开始\n/help - 帮助\n/ai [问题] - 人工智能"
			case "ai":
				question := strings.TrimSpace(update.Message.CommandArguments())
				if question == "" {
					replyText = "例如: /ai 你是谁"
				} else {
					replyText, err = callGeminiAPI(question)
					if err != nil {
						replyText = fmt.Sprintf("调用人工智能失败: %v", err)
					}
				}
			default:
				replyText = getRandomCreativeReply()
			}

			// 发送消息并检查错误
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				log.Println("消息发送失败: ", err)
			}
		}(update)
	}

	// 等待所有 Goroutines 完成
	wg.Wait()
}