package netboxapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type NetBoxConnection struct {
	Url   string
	Token string
}

func NewConnection(Url string, Token string) NetBoxConnection {
	return NetBoxConnection{Url: Url, Token: Token}
}

func (n NetBoxConnection) GetL2VPN() {
	var reply Result
	url := "vpn/l2vpns/"
	err := json.Unmarshal(n.Query("get", url, 0, nil), &reply)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reply.Results)

}
func (n NetBoxConnection) Query(mode string, endpoint string, id int, data []byte) []byte {
	client := &http.Client{}
	fUrl := n.Url + "/" + endpoint
	if id != 0 {
		fUrl = n.Url + "/" + endpoint + "/" + strconv.Itoa(id)
	}
	req, err := http.NewRequest(strings.ToUpper(mode), fUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", "Token "+n.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
