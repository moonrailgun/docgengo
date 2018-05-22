package main

import (
  "fmt"
  "flag"
  "os"
  "io/ioutil"
  "log"
  "gopkg.in/yaml.v2"
  "html/template"
  "bufio"
  "strings"
)

type Param struct {
  Type string `yaml:"type,omitempty"`
  Default string `yaml:"default,omitempty"`
  Required bool `yaml:"required,omitempty"`
  Desc bool `yaml:"desc,omitempty"`
}
type Operation struct {
  Title string `yaml:"title"`
  Desc string `yaml:"desc"`
  Params []map[string]Param `yaml:"params"`
}
type Route map[string]Operation
type Group map[string]Route
type T struct {
  Host string `yaml:"host"`
  Group map[string]Group `yaml:"route"`
}


func main() {
  configPath := flag.String("c", ".docgengo.yaml", "doc config yaml")
  flag.Parse()
  log.Println("正在读取配置 :", *configPath)
  buf, err := ioutil.ReadFile(*configPath)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  conf := string(buf)
  fmt.Printf("配置文件:\n↓=========\n%s\n↑=========\n\n", conf)

  t := T{}
  err = yaml.Unmarshal(buf, &t)
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  fmt.Printf("路由列表:\n↓=========\n")
  for groupName, group := range t.Group {
    for path, info := range group {
      fmt.Printf("[GET]\t%s --- %v\n", groupName + ":" + path, info)
    }
  }
  fmt.Printf("↑=========\n\n")

  fmt.Printf("输出网页:\n↓=========\n")
  tplbuf, err := ioutil.ReadFile("template/template.html")
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  tplFunc := template.FuncMap{"toUpper": strings.ToUpper}
  tpl, err := template.New("dochtml").Funcs(tplFunc).Parse(string(tplbuf))
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  tplData := struct { Group map[string]Group } {
    Group: t.Group,
  }

  // 创建html文件
  f, err := os.Create("dist/api.html")
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  writer := bufio.NewWriter(f)
  // tpl.Execute(os.Stdout, tplData)
  err = tpl.Execute(writer, tplData)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  writer.Flush()
  fmt.Printf("↑=========\n\n")

  fmt.Printf("completed!")
}
