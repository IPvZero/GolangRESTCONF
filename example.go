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
        CiscoIOSXEBgpBgp []struct {
                ID  int `json:"id"`
                Bgp struct {
                        LogNeighborChanges bool `json:"log-neighbor-changes"`
                } `json:"bgp"`
                Neighbor []struct {
                        ID       string `json:"id"`
                        RemoteAs int    `json:"remote-as"`
                } `json:"neighbor"`
        } `json:"Cisco-IOS-XE-bgp:bgp"`
}

func main() {

        client := http.Client{Timeout: 3 * time.Second}

        req, err := http.NewRequest(http.MethodGet,
                "https://sandbox-iosxe-latest-1.cisco.com:443/restconf/data/native/router/bgp",
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
        //Prints the status code of the response
        fmt.Printf("Status Code: %d\n", res.StatusCode)
        //Prints the string output of the response
        fmt.Printf("Body: %s\n", string(rBody))
        //Assigns var, then loops through and parses data
        neig := yang.CiscoIOSXEBgpBgp[0].Neighbor
        for i := 0; i < len(neig); i++ {
                fmt.Println("This device has a BGP Neighbor:", neig[i].ID, "with an ASN:", neig[i].RemoteAs)
        }

}
