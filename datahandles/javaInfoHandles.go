package javaInfoHandle

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"text/tabwriter"
)

type JavaData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

/* 添加Java版本数据 */
func AddJava2Json(data JavaData) {
	// 读取本地java版本数据
	list, err := GetLocalJava()
	if err != nil {
		fmt.Println("读取本地java版本数据失败:", err)
		return
	}
	list = append(list, data)
	newData, err := json.Marshal(list)
	if err != nil {
		fmt.Println("格式化java版本数据失败:", err)
		return
	}
	err = os.WriteFile("config/data.json", newData, 0644)
	if err != nil {
		fmt.Println("写入java版本数据失败:", err)
		return
	}
	fmt.Println("添加Java版本数据成功!")
	// 格式化打印版本数据
	PrintJavaList(list)
}

/*从yojava管理列表中删除java*/
func RemoveJavaFromJson(nameOrVersion string) (isOk bool, err error) {
	// 读取本地java版本数据
	list, err := GetLocalJava()
	if err != nil {
		return false, errors.New("读取Java版本信息失败: " + err.Error())
	}
	// 遍历数据
	found := false
	for i, data := range list {
		// 完全匹配直接删除结束
		if data.Name == nameOrVersion || data.Version == nameOrVersion {
			list = append(list[:i], list[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		// 模糊匹配进行删除
		state := listFuzzyMatchDel(&list, nameOrVersion)
		if state {
			return true, nil
		}

		// 无法匹配到则报错提示
		return false, errors.New("没有找到对应" + nameOrVersion + "的Java版本")
	}
	newData, err := json.Marshal(list)
	if err != nil {
		return false, err
	}
	err = os.WriteFile("config/data.json", newData, 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}

/* 打印Java版本数据 */
func PrintJavaList(list []JavaData) {
	fmt.Print("本地Java版本列表:\n\n")
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "Name\t\tVersion\t\tPath\n")
	for _, data := range list {
		fmt.Fprintf(w, "%s\t\t%s\t\t%s\n", data.Name, data.Version, data.Path)
	}
	w.Flush()
	data, err := os.ReadFile("descTemplate/usage.txt")
	if err != nil {
		fmt.Println("读取使用说明失败:", err)
		return
	}
	fmt.Println("\n", string(data))
}

/*
获取本机中收录的java版本
*/
func GetLocalJava() ([]JavaData, error) {
	data, err := os.ReadFile("config/data.json")
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

func listFuzzyMatchDel(list *[]JavaData, nameOrVersion string) bool {
	var found bool = false
	reg := regexp.MustCompile(`^[` + nameOrVersion + `]+`)
	// 查询匹配项列表
	matchList := []JavaData{}
	for _, data := range *list {
		if reg.MatchString(data.Version) {
			matchList = append(matchList, data)
		}
	}
	if len(matchList) > 0 {
		found = true
		// 排序，版本号更大的在前面
		sort.Slice(matchList, func(i, j int) bool {
			return matchList[i].Version >= matchList[j].Version
		})
		// 从list中删除匹配matchList的第一项
		for i, data := range *list {
			if data.Version == matchList[0].Version {
				*list = append((*list)[:i], (*list)[i+1:]...)
				break
			}
		}
		return found
	}
	return found
}
