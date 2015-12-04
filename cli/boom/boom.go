package main

import (
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
  "fmt"
)

func main() {
  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Action = func(c *cli.Context) {
    output, err := exec.Command("git", "status").Output()

    if err != nil {
        fmt.Println("Error getting branch")
    } else {
        fmt.Println(string(output))
    }
  }

  app.Run(os.Args)
}
