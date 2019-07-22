package dispatcher

import (
    "fmt"
    "github.com/cihub/seelog"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "strings"
)

func Dispatcher(ip string, data []byte) ([]byte, error) {
    // Client http.Client
    var Client *http.Client
    seelog.Info("InitClient begin ...")
    Client = &http.Client{}
    jar, _ := cookiejar.New(nil)
    Client.Jar = jar

    //data := map[string][]string{}

    req, err := http.NewRequest("POST", fmt.Sprintf("http://%v/exec", ip), strings.NewReader(string(data)))
    if err != nil {
        seelog.Errorf("ERROR : %v", err.Error())
        return nil, err
    }

    header := map[string][]string{}
    req.Header = header

    res, err := Client.Do(req)
    if err != nil {
        seelog.Errorf("ERROR : %v", err.Error())
        return nil, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        seelog.Errorf("ERROR : %v", err.Error())
        return nil, err
    }

    seelog.Trace(string(body))

    return body, nil
}
