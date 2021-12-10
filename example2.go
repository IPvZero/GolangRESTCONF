package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "time"
)

type Response struct {
        CiscoIOSXENativeMemory struct {
                Free struct {
                        LowWatermark struct {
                                Processor int `json:"processor"`
                        } `json:"low-watermark"`
                } `json:"free"`
        } `json:"Cisco-IOS-XE-native:memory"`
}

func main() {

        client := http.Client{Timeout: 3 * time.Second}

        req, err := http.NewRequest(http.MethodGet,
                "https://sandbox-iosxe-latest-1.cisco.com:443/restconf/data/native/memory",
                http.NoBody)
        if err != nil {
                log.Fatal(err)
        }

        req.SetBasicAuth("developer", "C1sco12345")
        req.Header.Set("Accept", "application/yang-data+json")

        res, err := client.Do(req)
        if err != nil {
                log.Fatal(err)
        }

        defer res.Body.Close()

        rBody, err := ioutil.ReadAll(res.Body)
        if err != nil {
                log.Fatal(err)
        }

        empBytes := []byte(rBody)
        var yang Response
        json.Unmarshal(empBytes, &yang)
        //This prints the status code of the response
        fmt.Printf("Status Code: %d\n", res.StatusCode)
        //This prints the string output of the response
        fmt.Printf("Body: %s\n", string(rBody))
        //This parses the name field of the struct and appends to the print statement
        fmt.Println("The memory processor value on this device is", yang.CiscoIOSXENativeMemory.Free.LowWatermark.Processor)

}
