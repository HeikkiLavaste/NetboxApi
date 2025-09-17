package netboxapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type NetBoxConnection struct {
	Url   string
	Token string
}

func NewConnection(Url string, Token string) NetBoxConnection {
	return NetBoxConnection{Url: Url, Token: Token}
}

func (n NetBoxConnection) GetL2VPN() {
	url := "vpn/l2vpns/"
	fmt.Println(string(n.Query(url, 0)))

}
func (n NetBoxConnection) Query(endpoint string, id int) []byte {
	client := &http.Client{}
	fUrl := n.Url + "/" + endpoint
	if id != 0 {
		fUrl = n.Url + "/" + endpoint + "/" + strconv.Itoa(id)
	}
	req, err := http.NewRequest("GET", fUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", "Token "+n.Token)
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
