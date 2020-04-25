package network

import (
    "bytes"
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
)
const (
    baseUrl = `https://httpbin.org`
)

func Get() {
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }
    resp, err := client.Get(baseUrl+"/get")
    if err != nil {
        fmt.Println("get httpbin website error: ", err)
        return
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read data from httpbin error: ", err)
        return
    }
    fmt.Println(string(data))
}

func Post() {
    reqBody := map[string]string{
        "username": "cooc",
        "password": "123456",
    }
    body, _ := json.Marshal(reqBody)
    req, err := http.NewRequest("POST", baseUrl+"/post", bytes.NewBuffer(body))
    if err != nil {
        fmt.Println("build post rquest error: ", err)
        return 
    }
    req.Header.Add("Content-Type", "application/json")
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("post httpbin website error: ", err)
        return
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read data from post request error: ", err)
        return
    }
    fmt.Println(string(data))
}
