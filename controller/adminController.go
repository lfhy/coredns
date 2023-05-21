package controller

import (
	"dns/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Oninit() {
	go model.WatchDBUpdate()
}

func DnsEditGet() func(*gin.Context) {
	return func(c *gin.Context) {
		key := c.DefaultQuery("key", "")
		if key == "" {
			c.HTML(http.StatusOK, "dnsedit.html", gin.H{})
		} else {
			obj, err := model.DnsGet(key)
			if err != nil {
				c.HTML(http.StatusOK, "dnsedit.html", gin.H{"Message": model.Message{Error: err.Error()}})
			} else {
				c.HTML(http.StatusOK, "dnsedit.html", gin.H{"data": obj})
			}
		}
	}
}
func DnsEditPost() func(*gin.Context) {
	return func(c *gin.Context) {
		key := c.DefaultQuery("key", "")
		name := c.PostForm("name")
		data := c.PostForm("data")
		ttl := c.PostForm("ttl")
		intTTL, _ := strconv.Atoi(ttl)
		if name == "" || data == "" {
			c.HTML(http.StatusOK, "dnsedit.html", gin.H{"obj": &model.Dns{Origin: name, NameServer: data, TTL: intTTL, Key: key}, "error": "名称和数据不能为空！"})
			return
		}
		var value []byte
		if intTTL == 0 {
			value, _ = json.Marshal(model.A{Host: data})
		} else {
			value, _ = json.Marshal(model.A{Host: data, TTL: intTTL})
		}
		if key == "" {
			model.DnsAdd(name, string(value))
		} else {
			model.DnsEdit(key, string(value))
		}
		c.Redirect(http.StatusMovedPermanently, "/admin/dns")
	}
}

func Dnslist() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "dns2.html", gin.H{})
	}
}
func Dnslist2() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "dns2.html", gin.H{})
	}
}
func WsHandler() func(context *gin.Context) {
	return func(c *gin.Context) {
		var conn *websocket.Conn
		var err error
		Wsupgrader := websocket.Upgrader{
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			HandshakeTimeout: 5 * time.Second,
			// 取消ws跨域校验
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err = Wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		go func(conn *websocket.Conn) {
			defer conn.Close()
			for {
				select {
				case message, ok := <-model.NewMessage:
					if !ok {
						conn.WriteMessage(websocket.CloseMessage, []byte{})
					}
					conn.PingHandler()
					err := conn.WriteJSON(message)
					if err != nil {
						fmt.Println(err)
						model.NewMessage <- message
						break
					}
				}
			}
		}(conn)
	}
}

func DelDns() func(*gin.Context) {
	return func(c *gin.Context) {
		key := c.Query("key")
		err := model.DnsDel(key)
		if err != nil {
			c.JSON(http.StatusOK, model.Message{Error: err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, model.Message{Error: ""})
			return
		}

	}
}

// api 接口
func DnsApiList() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, model.DnsData{Data: model.DnsList()})
	}
}
