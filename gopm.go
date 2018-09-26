// 项目管理工具
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var cmd string
var projectname string

/*
* 项目名不能用clean，create，work
 */
func main() {
	help()
	for i := 0; i < 1; i++ {
		// 读取用户输入
		fmt.Scanln(&cmd, &projectname)
		if cmd != "" {
			result := handle(cmd, projectname)
			if result == false {
				i--
			}
		} else {
			// 没输入命令 继续循环
			fmt.Println("您的输入为空，请输入命令")
			i--
		}
		// cmd初始化
		cmd = ""
	}
}

// handle 处理
func handle(cmd string, projectname string) bool {
	switch cmd {
	// 退出
	case "exit":
		{
			fmt.Println("Goodbye")
			return true
		}
	// 显示帮助
	case "help":
		{
			help()
			return false
		}
	// 清空work目录
	case "clean":
		{
			clean()
			return false
		}
	// 创建项目
	case "create":
		{
			create(projectname)
			return false
		}
	// 使用项目
	case "use":
		{
			if clean() == true {
				use(projectname)
			} else {
				fmt.Println("work目录清理失败，请手工修改work目录文件名")
			}
			return false
		}
	// 没有此命令
	default:
		{
			fmt.Println("没有此命令")
			return false
		}
	}
}

// create 创建项目
func create(projectname string) bool {
	if projectname == "" {
		fmt.Println("新建项目名称不能为空")
		return false
	}
	// 获取当前目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 不能用work
	if projectname == "work" {
		fmt.Println("创建失败，不能使用此项目名", projectname)
		return false
	}
	// 读取work项目名称
	f, fileerr := os.Open(dir + "/work/project.txt")
	if fileerr == nil {
		workprojectname, _ := ioutil.ReadAll(f)
		if projectname == string(workprojectname) {
			fmt.Println("创建失败，此项目名正在使用", projectname)
			return false
		}
	}
	defer f.Close()
	// 判断是否已经存在projectname
	_, haveerr := os.Stat(dir + "/" + projectname)
	if haveerr == nil {
		fmt.Println("创建失败，项目已存在", haveerr)
		return false
	}
	// 创建目录 子目录 bin pkg src project.txt
	rooterr := os.Mkdir(dir+"/"+projectname, os.ModePerm)
	if rooterr != nil {
		fmt.Println(rooterr)
		return false
	}
	binerr := os.Mkdir(dir+"/"+projectname+"/bin", os.ModePerm)
	if binerr != nil {
		fmt.Println(binerr)
		return false
	}
	pkgerr := os.Mkdir(dir+"/"+projectname+"/pkg", os.ModePerm)
	if pkgerr != nil {
		fmt.Println(pkgerr)
		return false
	}
	srcerr := os.Mkdir(dir+"/"+projectname+"/src", os.ModePerm)
	if srcerr != nil {
		fmt.Println(srcerr)
		return false
	}
	// 新建文件
	fw, fwerr := os.Create(dir + "/" + projectname + "/project.txt")
	if fwerr != nil {
		fmt.Println(fwerr)
		return false
	}
	fwstrerr := ioutil.WriteFile(dir+"/"+projectname+"/project.txt", []byte(projectname), 0666)
	if fwstrerr != nil {
		fmt.Println(fwstrerr)
		return false
	}
	defer fw.Close()
	fmt.Println("成功创建项目", projectname)
	return true
}

// use 使用项目：将项目的名字改为work，操作之前先将work名字改为其他
func use(projectname string) bool {
	if projectname == "" {
		fmt.Println("要操作的项目名称不能为空")
		return false
	}
	// 获取当前目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 判断是否存在work，存在直接返回失败
	_, workerr := os.Stat(dir + "/work")
	if workerr == nil {
		fmt.Println("work目录已经存在，请先清理work", workerr)
		return false
	}
	// 判断是否存在目标文件夹
	objectdir := dir + "/" + projectname
	_, haveerr := os.Stat(objectdir)
	if haveerr != nil {
		fmt.Println("项目"+projectname+"不存在，请输入正确的项目名称", haveerr)
		return false
	}
	// 重命名文件夹
	actionerr := os.Rename(objectdir, dir+"/work")
	if actionerr != nil {
		fmt.Println(actionerr)
		return false
	}
	fmt.Println("切换工作区成功", projectname)
	return true
}

// clean 清空work目录：读取work里的项目名称 将work文件夹改名为项目名称
func clean() bool {
	// 获取当前目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 判断是否存在work，不存在直接返回成功
	_, workerr := os.Stat(dir + "/work")
	if workerr != nil {
		fmt.Println("work已清理")
		return true
	}
	// 读取项目名称
	f, fileerr := os.Open(dir + "/work/project.txt")
	if fileerr != nil {
		fmt.Println(fileerr)
		return false
	}
	projectname, _ := ioutil.ReadAll(f)
	f.Close()
	// 判断是否存在目标文件夹
	objectdir := dir + "/" + string(projectname)
	_, haveerr := os.Stat(objectdir)
	if haveerr == nil {
		fmt.Println(haveerr)
		return false
	}
	// 重命名文件夹
	actionerr := os.Rename(dir+"/work", objectdir)
	if actionerr != nil {
		fmt.Println(actionerr)
		return false
	}
	fmt.Println("清理work成功")
	return true
}

// help 显示帮助文档
func help() {
	help := `
----------
命令列表
将xxx项目目录设置为work: use xxx项目名
清空work目录: clean
创建新项目: create xxx
退出命令行: exit
帮助: help
PS：如果windows下出现 rename xxx xxx Access is denied. 现象，请关闭所有打开的文件夹/资源管理器窗口再进行操作
----------
请输入命令:`
	fmt.Println(help)
}
