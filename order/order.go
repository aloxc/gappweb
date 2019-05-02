package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aloxc/gappweb/config"
	"github.com/aloxc/gappweb/io"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	"os"
	"time"
)

var orderSrvHost string = ""

func init() {
	orderSrvHost = os.Getenv(config.ORDER_SERVER_HOST)
	if orderSrvHost == "" {
		orderSrvHost = config.ORDER_SERVER_HOST_DEFAULT
	}
}

type Order struct {
	Id          int32
	GoodsId     int32
	UserId      int32
	UserName    string
	Address     string
	Create_time time.Time
	Version     int32
}

func (this *Order) String() string {
	return fmt.Sprintf("Order[id:%d,goodsId:%d,userId:%d,userName:%s,address:%s]",
		this.Id, this.GoodsId, this.UserId, this.UserName, this.Address)
}

func GetOrder(order *Order) {
	d := client.NewPeer2PeerDiscovery("tcp@"+orderSrvHost, "")
	xclient := client.NewXClient(config.ORDER_SERVER_PATH, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	request := &io.Request{
		Method: "getOrder",
		Params: map[string]interface{}{
			"id": order.Id,
		},
		Head: io.RequestHead{
			UserPort:   33,
			UserIp:     "1.2.3.4",
			UserCookie: "",
		},
	}
	response := &io.Response{}
	err := xclient.Call(context.Background(), "Execute", request, response)
	bytes, err := json.Marshal(response.Data)
	json.Unmarshal(bytes, order)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
}
