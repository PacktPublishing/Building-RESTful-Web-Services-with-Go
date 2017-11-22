package main

import (
        "bytes"
        "fmt"
        "github.com/julienschmidt/httprouter"
        "log"
        "net/http"
        "os/exec"
)


func getCommandOutput(command string, arguments ...string) string{
        // args... unpacks arguments array into elements
        cmd := exec.Command(command, arguments...)
        var out bytes.Buffer
        var stderr bytes.Buffer
        cmd.Stdout = &out
        cmd.Stderr = &stderr
        err := cmd.Start()
        if err != nil {
                log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
        }
        err = cmd.Wait()
        if err != nil {
                log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
        }
        return out.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
        fmt.Fprintf(w, getCommandOutput("/usr/local/bin/go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
     fmt.Fprintf(w, getCommandOutput("/bin/cat",  params.ByName("name")))
}

func main() {
        router := httprouter.New()
        // Mapping to methods is possible with HttpRouter
        router.GET("/api/v1/go-version", goVersion)
        // Path variable called name used here
        router.GET("/api/v1/show-file/:name", getFileContent)
        log.Fatal(http.ListenAndServe(":8000", router))
}
