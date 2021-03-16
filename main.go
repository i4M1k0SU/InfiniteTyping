package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		// error
		fmt.Println("Invalid .env")
		return
	}

	SLACK_TOKEN := os.Getenv("SLACK_TOKEN")
	TARGET_CHANNEL_ID := os.Getenv("TARGET_CHANNEL_ID")

	api := slack.New(SLACK_TOKEN)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		info := rtm.GetInfo()
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			if ev.Channel != TARGET_CHANNEL_ID || ev.User == info.User.ID {
				break
			}
			fmt.Println("MessageEvent")
			fmt.Println("Typing")
			rtm.SendMessage(rtm.NewTypingMessage(TARGET_CHANNEL_ID))

		case *slack.ReactionAddedEvent:
			if ev.Item.Channel != TARGET_CHANNEL_ID || ev.User == info.User.ID {
				break
			}
			fmt.Println("ReactionAddedEvent")
			fmt.Println("Typing")
			rtm.SendMessage(rtm.NewTypingMessage(TARGET_CHANNEL_ID))

		case *slack.ReactionRemovedEvent:
			if ev.Item.Channel != TARGET_CHANNEL_ID || ev.User == info.User.ID {
				break
			}
			fmt.Println("ReactionRemovedEvent")
			fmt.Println("Typing")
			rtm.SendMessage(rtm.NewTypingMessage(TARGET_CHANNEL_ID))

		case *slack.UserTypingEvent:
			if ev.Channel != TARGET_CHANNEL_ID || ev.User == info.User.ID {
				break
			}
			fmt.Println("TypingEvent")
			fmt.Println("Typing")
			rtm.SendMessage(rtm.NewTypingMessage(TARGET_CHANNEL_ID))

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
			// Ignore other events
		}
	}
}
