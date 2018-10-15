package command

import (
	"context"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/line/line-bot-sdk-go/linebot"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func init() {
	categoryStoI = make(map[string]int)
	categoryItoS = make(map[int]string)
	categoryStoI["unknown"] = 0
	categoryStoI["breakfast"] = 1
	categoryStoI["lunch"] = 2
	categoryStoI["dinner"] = 3
	categoryStoI["drink"] = 4
	categoryStoI["snack"] = 5
	categoryStoI["closthing"] = 6
	categoryStoI["fun"] = 7
	categoryStoI["traffic"] = 8
	categoryStoI["sport"] = 9

	for k, v := range categoryStoI {
		categoryItoS[v] = k
	}
}

type messageCMD struct {
	db *sqlx.DB
}

type accountingModel struct {
	price          float64
	category       int
	accountingTime string
	detail         string
}

type RemainCmd []string

func (r *RemainCmd) Set(value string) error {
	*r = append(*r, value)
	return nil
}

func (r *RemainCmd) String() string {
	return ""
}

func (r *RemainCmd) IsCumulative() bool {
	return true
}

func RemainCmdList(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*RemainCmd)(target))
	return
}

const costUsage = "Usage:\n\tcost PRICE CATEGORY -t --time ACCOUTING-TIME(2018/01/01) -d --detail DETAIL"

var categoryStoI map[string]int
var categoryItoS map[int]string

func parseCostCommand(cmd []string) (*accountingModel, error) {
	// example: cost 999 breakfast -t 2018-09-09 -d kappa breakfast
	accounting := kingpin.New("cost", "Accounting command line.")
	price := accounting.Arg("price", "Price").Required().Float64()
	category := accounting.Arg("category", "Category").Default("unknown").String()
	accountingTime := accounting.Flag("time", "Accounting time.").Default(time.Now().UTC().Format("2006-01-02")).Short('t').String()
	detail := accounting.Flag("detail", "Detail.").Default("").Short('d').String()
	remain := RemainCmdList(accounting.Arg("remain", "Remain part of command."))
	_, err := accounting.Parse(cmd)
	if err != nil {
		return nil, err
	}

	// combine data and validate content
	*category = strings.ToLower(*category)
	categoryEnum, exist := categoryStoI[*category]
	if !exist {
		categoryEnum = 0
	}
	var combinedDetail string
	if len(*remain) != 0 {
		combinedDetail = *detail + " " + strings.Join(*remain, " ")
	} else {
		combinedDetail = *detail
	}
	c := &accountingModel{price: *price, category: categoryEnum, accountingTime: *accountingTime, detail: combinedDetail}
	return c, nil
}

func (c *messageCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		splitedCommand := strings.Split(message.Text, " ")
		if strings.ToLower(splitedCommand[0]) == "cost" {
			// parse cost command
			param, err := parseCostCommand(splitedCommand[1:])
			if err != nil {
				replyMessage := linebot.NewTextMessage(err.Error())
				return []linebot.Message{replyMessage}, nil
			}
			userID := event.Source.UserID

			// insert new record to db
			insert := "INSERT INTO Accounting (userID, category, accountingTime, price, detail) VALUES (?, ?, ?, ?, ?)"
			_, err = c.db.Exec(insert, userID, param.category, param.accountingTime, param.price, param.detail)
			if err != nil {
				replyMessage := linebot.NewTextMessage(err.Error())
				return []linebot.Message{replyMessage}, nil
			}
			replyMessage := linebot.NewTextMessage("Accounting recorded!")
			return []linebot.Message{replyMessage}, nil
		}
		replyMessage := linebot.NewTextMessage("Unknown command!\n" + costUsage)
		return []linebot.Message{replyMessage}, nil
	default:
		replyMessage := linebot.NewTextMessage("Not supported event yet, sorry!")
		return []linebot.Message{replyMessage}, nil
	}
}

func (c *messageCMD) Register(cmd string, handler func() error) error {
	return nil
}

// NewMessageCommand returns command concrete instance
func NewMessageCommand(db *sqlx.DB) Interface {
	return &messageCMD{db}
}
