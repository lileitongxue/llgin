package getlogs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
)

const (
	JSON         = "application/json"
	String       = "application/string"
	HTML         = "application/html"
	IndentedJSON = "application/IndentedJSON"
)

var log lager.Logger
var httpClientG = &http.Client{
	Transport: httpClientB.Transport,
	Timeout:   time.Duration(10) * time.Second,
}
var httpClientB = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout: 0,
}

func init() {
	log = lager.NewLogger("Deployment")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

func GenRequest(method, url, token string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error
	apiHost := "10.1.0.4:6443"
	url = "https://" + apiHost + url
	if len(body) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-patch+json")
	req.Header.Set("Authorization", token)
	return httpClientG.Do(req)
}

//打印返回的数据
// func myPost(w http.ResponseWriter, r *http.Request) {
// 	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
// 	fmt.Fprintf(w, "%s", s)        //在返回页面中显示内容。
// }

func GetInstanceLogs(c *gin.Context) {
	token := "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLWsxeDA4Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIzMzlhNmE1Zi0yMTUwLTExZTgtYjczNi01MjU0MDBhODNmYjEiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06ZGVmYXVsdCJ9.uSHRJwmQfH9MB9Nn6lum2jCc6G2lvL1iKGtKnBIDMWYkP2GAdV-s-LY18h30PNGf0xsl96qDpiZgMoavt1bKycWFfZ0EQ8YPG4FGUmTP_Y3gbANT-B3dNUBURueaa1WzPp4aGGZFoiyn7C9XY-RfdCAEZbQpJIRcVuYAvuZCqSyQ_EcGs616hiTjS8ASWBoB3seqhiMHyNMetz5RZt_UESjeBUV6nVCPNFvptSGxkTxiMUcWVyRRnUVl4diyJN0enpWQf1ywJIMNCryM3ns7EqjEPYkcWzvHFbsdUDn-Gr45z0SybEo6Q1TnKRlLN02fmUs4CKaoi4sWZGNo-IxbvA"
	namespace := c.Query("ns")
	instance := c.Query("instance")
	previous := c.DefaultQuery("previous", "false")
	req, err := GenRequest("GET", "/api/v1/namespaces/"+namespace+"/pods/"+instance+"/log?&previous="+previous+"&container=micro-service&tailLines=500", token, []byte{})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	if req.StatusCode != http.StatusOK {
		var err_messages map[string]interface{}
		project_message, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(project_message, &err_messages)
		err_message, _ := err_messages["message"]
		s := []map[string]interface{}{}
		m1 := map[string]interface{}{"err_info": err_message, "suggest": "Please input right instance name !"}
		s = append(s, m1)
		b, err := json.Marshal(s)
		if err != nil {
			fmt.Println(err)
		}
		c.Data(req.StatusCode, JSON, b)
	} else {

		defer req.Body.Close()
		c.Data(req.StatusCode, JSON, result)
	}
}
