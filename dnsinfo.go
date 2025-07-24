package main

import (
        "bufio"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "strings"
)

func main() {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter domain (e.g., google.com): ")
        domain, _ := reader.ReadString('\n')
        domain = strings.TrimSpace(domain)

        url := fmt.Sprintf("https://XXXX.com/%s", domain)

        resp, err := http.Get(url)
        if err != nil {
                fmt.Println("Error fetching domain info:", err)
                return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Println("Error reading response:", err)
                return
        }

        // Try pretty-printing
        var prettyJSON map[string]interface{}
        err = json.Unmarshal(body, &prettyJSON)
        if err != nil {
                fmt.Println("Raw response:")
                fmt.Println(string(body)) // fallback
        } else {
                formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
                fmt.Println("Formatted JSON:")
                fmt.Println(string(formatted))
        }

        // Save raw JSON to file
        filename := fmt.Sprintf("domaininfo_%s.json", strings.ReplaceAll(domain, ".", "_"))
        err = os.WriteFile(filename, body, 0644)
        if err != nil {
                fmt.Println("Error writing file:", err)
        } else {
                fmt.Println("Saved to:", filename)
        }
}

