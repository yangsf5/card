// Author: sheppard(ysf1026@gmail.com) 2014-01-08

package hall

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/yangsf5/card/util"
)

type RoomConfig struct {
	XMLName xml.Name `xml:"room"`
	Games []Game `xml:"game"`
}

type Game struct {
	Name string `xml:"name,attr"`
	Items []Item `xml:"item"`
}

type Item struct {
	Name string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}


var (
	configXml RoomConfig
	configs map[string] map[string] string
)

func init() {
	configs = make(map[string] map[string] string)
}


func ReadConfig(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	util.CheckFatal(err)

	err = xml.Unmarshal(content, &configXml)
	util.CheckFatal(err)

	fmt.Println(configXml)

	configTurn()
	fmt.Println(configs)
}

func configTurn() {
	for _, game := range configXml.Games {
		for _, item := range game.Items {
			config, ok := configs[game.Name]
			if !ok {
				config = make(map[string] string)
				configs[game.Name] = config
			}
			_, ok = config[item.Name]
			if ok {
				panic("Room config field repeated, field=" + item.Name)
			}
			config[item.Name] = item.Value
		}
	}
}
