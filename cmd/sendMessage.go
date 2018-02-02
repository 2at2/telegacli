package cmd

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mono83/xray"
	xargs "github.com/mono83/xray/args"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// TgToken telegram api token <botId>:<token>
const TgToken = "TGTOKEN"

var sendMessageCmdToken string

// SendMessageCmd sends message to chat
var SendMessageCmd = &cobra.Command{
	Use:   "send <chatId> <message>",
	Short: "Sends message to telegram chat",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ray := xray.ROOT.Fork().WithLogger("sender")

		chatID, err := strconv.ParseInt(args[0], 10, 64)

		if err != nil {
			ray.Error("Unable to parse chat id - :err", xargs.Error{Err: err})
			return err
		}

		ray.Debug("Parsed chat id :id", xargs.ID64(chatID))

		message := args[1]

		if len(message) == 0 {
			ray.Error("Empty message given")
			return errors.New("empty message")
		}

		var token string
		if len(sendMessageCmdToken) != 0 {
			token = sendMessageCmdToken
		} else if len(os.Getenv(TgToken)) != 0 {
			token = os.Getenv(TgToken)
		} else {
			ray.Error("Undefined telegram api token")
			return errors.New("undefined api token. Please provide it through the flag --token or env variable " + TgToken)
		}

		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			ray.Error("Unable to create bot api - :err", xargs.Error{Err: err})
			return err
		}

		ray.Debug("Bot api is created")

		if _, err := bot.Send(tgbotapi.NewMessage(chatID, message)); err != nil {
			ray.Error("Unable to send message to chat :id - :err", xargs.ID64(chatID), xargs.Error{Err: err})
			return err
		}

		ray.Debug("Message sent to chat :id successfully", xargs.ID64(chatID))

		return nil
	},
}

func init() {
	SendMessageCmd.Flags().StringVarP(&sendMessageCmdToken, "token", "t", "", "")
}
