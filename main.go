package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Account struct {
	Email    string
	password string
	Money    float64
	Sex      bool
}

type Conf struct {
	Redis *Redis `yaml:"redis" json:"redis"`
}

type Redis struct {
	Addrs    []string `yaml:"addrs" json:"addrs"`
	Password string   `yaml:"password" json:"password"`
}

type Channel struct {
	ID            string `json:"id"`
	Account       string `json:"account"`
	Passpwd       string `json:"passpwd"`
	Corpid        string `json:"corpid"`
	TenantId      string `json:"tenantId"`
	LocalExpidLen int    `json:"localExpidLen"`
	ChannelIp     string `json:"channelIp"`
	ChannelPort   int    `json:"channelPort"`
	ServiceId     string `json:"serviceid"`
	SendRate      int    `json:"sendRate"`
	MsgFormat     int    `json:"msgformat"`
	WordsLen      int    `json:"wordsLen"`

	Heartbeat int `json:"heartbeat"`

	Name        string `json:"channelName"`
	WindowStats bool   `json:"windowstats"`
}

func main01() {
	account := Account{
		Email:    "rsj217@gmail.com",
		password: "123456",
		Money:    100.5,
		Sex:      false,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(rs)
	fmt.Println(string(rs))
}

func LoadConf(path string) (*Conf, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Conf
	if err = yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func getRedisKey(key string) string {
	return fmt.Sprintf("channel-info:%s", key)
}
