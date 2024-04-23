package javaInfoHandle

import (
	"encoding/json"
	"fmt"
	"os"
)

type JavaData struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Path     string `json:"path"`
}

/* 添加Java版本数据 */
func AddJava2Json(data JavaData){
	// 读取本地java版本数据
	list, err := GetLocalJava()
	if err!= nil {
		fmt.Println("读取本地java版本数据失败:", err)
	  return
	}
	list = append(list, data)
	newData, err := json.Marshal(list)
	if err!= nil {
		fmt.Println("格式化java版本数据失败:", err)
	  return
	}
	err = os.WriteFile("data.json", newData, 0644)
	if err!= nil {
		fmt.Println("写入java版本数据失败:", err)
	  return
	}
	fmt.Println("添加Java版本数据成功!")
	// 格式化打印版本数据
	PrintJavaList(list)
}

func RemoveJavaFromJson(name string){
  // 读取本地java版本数据
  list, err := GetLocalJava()
  if err!= nil {
    fmt.Println("读取Java版本信息失败", err)
    return
  }
  // 遍历数据
	found := false
	for i, data := range list {
		if data.Name == name {
			list = append(list[:i], list[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("%s 没有找到对应的Java版本\n", name)
		return
	}
	newData, err := json.Marshal(list)
	if err!= nil {
		fmt.Println(err)
	  return
	}
	err = os.WriteFile("data.json", newData, 0644)
	if err!= nil {
		fmt.Println(err)
	  return
	}
	fmt.Println(name + "删除成功!")
}

/* 打印Java版本数据 */
func PrintJavaList(list []JavaData){
	fmt.Println("Name", "\t", "Version", "\t", "Path")
	for _, data := range list {
		fmt.Println(data.Name, "\t", data.Version, "\t", data.Path)
	}
}

/*
获取本机中收录的java版本
*/
func GetLocalJava() ([]JavaData, error) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var list []JavaData
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}