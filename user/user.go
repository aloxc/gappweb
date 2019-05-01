package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aloxc/gappweb/io"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
)

type User_Level int

const (
	User_Level_OK         User_Level = iota //正常用户
	User_Level_DENY                         //被封杀用户
	User_LEVEL_BLACK_LIST                   //黑名单用户
)

type User struct {
	Id       int32
	UserName string
	Password string
	Level    User_Level
}

func (this *User) String() string {
	return fmt.Sprintf("User[id:%d,name:%s,password:%s,level:%d]", this.Id, this.UserName, this.Password, this.Level)
}

func GetUser(user *User) {
	d := client.NewPeer2PeerDiscovery("tcp@localhost:13331", "")
	xclient := client.NewXClient("gappuser", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	request := &io.Request{
		Method: "getUser",
		Params: map[string]interface{}{
			"userId": user.Id,
		},
		Head: io.RequestHead{
			UserPort:   33,
			UserIp:     "1.2.3.4",
			UserCookie: "",
		},
	}
	if user.Password != "" {
		request.Params["password"] = user.Password
	}
	response := &io.Response{}
	err := xclient.Call(context.Background(), "Execute", request, response)
	bytes, err := json.Marshal(response.Data)
	json.Unmarshal(bytes, user)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Infof("[%d].info = %s", user.Id, user.String())
}
func Register(userName string, password string, level User_Level) *User {
	d := client.NewPeer2PeerDiscovery("tcp@localhost:13331", "")
	xclient := client.NewXClient("gappuser", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	request := &io.Request{
		Method: "register",
		Params: map[string]interface{}{
			"userName": userName,
			"password": password,
			"level":    level,
		},
		Head: io.RequestHead{
			UserPort:   33,
			UserIp:     "1.2.3.4",
			UserCookie: "",
		},
	}
	response := &io.Response{}
	err := xclient.Call(context.Background(), "Execute", request, response)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	bytes, err := json.Marshal(response.Data)

	var user User
	json.Unmarshal(bytes, &user)
	log.Infof("[%d].info = %s", user.Id, user.String())
	return &user
}
