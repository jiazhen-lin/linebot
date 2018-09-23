package command

import (
	"context"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/line/line-bot-sdk-go/linebot"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type messageCMD struct {
	db *sqlx.DB
}

type costParam struct {
	price          float64
	category       string
	accountingTime string
	detail         string
}

var costUsage = "Usage:\n\tcost PRICE CATEGORY -t --time ACCOUTING-TIME -d --detail DETAIL"

func parseCostCommand(cmd []string) (*costParam, error) {
	a := kingpin.New("cost", "Accounting command line.")
	price := a.Arg("price", "Price").Required().Float64()
	category := a.Arg("category", "Category").Required().String()
	accountingTime := a.Flag("time", "Accounting time.").Default(time.Now().UTC().Format("2006-01-02")).Short('t').String()
	detail := a.Flag("detail", "Detail.").Default("").Short('d').String()
	_, err := a.Parse(cmd)
	if err != nil {
		return nil, err
	}
	c := &costParam{price: *price, category: *category, accountingTime: *accountingTime, detail: *detail}
	return c, nil
}

func (c *messageCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		splitedCommand := strings.Split(message.Text, " ")
		if splitedCommand[0] == "cost" || splitedCommand[0] == "Cost" {
			// parse cost command
			param, err := parseCostCommand(splitedCommand[1:])
			if err != nil {
				m := linebot.NewTextMessage(err.Error())
				return []linebot.Message{m}, nil
			}
			userID := event.Source.UserID

			// insert new record to db
			insert := "INSERT INTO Accounting (userID, category, accountingTime, price, detail) VALUES (?, ?, ?, ?)"
			_, err = c.db.Exec(insert, userID, param.category, param.accountingTime, param.price, param.detail)
			if err != nil {
				m := linebot.NewTextMessage(err.Error())
				return []linebot.Message{m}, nil
			}
			m := linebot.NewTextMessage("Accounting recorded!")
			return []linebot.Message{m}, nil
		}
		m := linebot.NewTextMessage("Unknown command!\n" + costUsage)
		return []linebot.Message{m}, nil
	default:
		m := linebot.NewTextMessage("Not supported event yet, sorry!")
		return []linebot.Message{m}, nil
	}
}

func (c *messageCMD) Register(cmd string, handler func() error) error {
	return nil
}

// NewMessageCommand returns command concrete instance
func NewMessageCommand(db *sqlx.DB) Interface {
	return &messageCMD{db}
}
