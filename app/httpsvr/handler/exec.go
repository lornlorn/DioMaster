package handler

import (
    "app/dispatcher"
    "app/utils"
    "net/http"
    "strings"

    "github.com/cihub/seelog"
)

// ExecuteHandler func(res http.ResponseWriter, req *http.Request)
func ExecuteHandler(res http.ResponseWriter, req *http.Request) {

    seelog.Infof("RouteRequest Params Execute : %v", req.URL)
    // key := mux.Vars(req)["key"]

    reqBody := utils.ReadRequestBody2JSON(req.Body)
    seelog.Debugf("Request Body : %v", string(reqBody))

    reqURL := req.URL.Query()
    seelog.Debugf("Request Params : %v", reqURL)

    ips := utils.GetJSONResultFromRequestBody(reqBody, "data.ips")

    var ret []byte
    var retlist []string

    for _, ip := range strings.Split(ips.String(), ",") {
        output, err := dispatcher.Dispatcher(ip, reqBody)

        if err != nil {
            seelog.Errorf("Command Run Error : %v", err.Error())
            retlist = append(retlist, err.Error())
        } else {
            seelog.Debugf("执行结果 : %v", string(output))
            retlist = append(retlist, string(output))
        }
    }

    ret = utils.GetAjaxRetWithDataJSON("0000", nil, retlist)
    res.Write(ret)
    return
}
