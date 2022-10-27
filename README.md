# gpm

##Golang多项目配置工具(原是一个老哥写的我再原基础上进行文档再次整理一备使用)

  初用golang，发现在同时进行好几个不相关的项目时候，很难管理，我希望我公司的项目和个人的项目严格分离。然而GOPATH里多个项目路径的设置导致引用的第三方包会在一个GOPATH路径里，管理很麻烦，所以写了这样一个小工具。初衷就是只设置一个GOPATH作为当前工作目录，通过重命名文件夹的方式将需要设置为当前工作目录的正开发项目路径改为GOPATH路径。不是当前正开发项目的改为其他路径。这样做到了每个项目都严格分离。

### win
* 安装golang到C:\Go目录
* project文件夹用于存放全部的项目，work存放当前工作的项目，每个项目包含与GOPATH一致的3个文件夹bin,pkg,src，额外包含一个project.txt文件，内容为项目名称
* 设置系统环境变量GOPATH 为 D:\xxx\project\work
* 设置系统环境变量GOBIN 为 D:\xxx\project\work\bin
### mac
* 安装golang到/usr/local/go
* project文件夹用于存放全部的项目，work存放当前工作的项目，每个项目包含与GOPATH一致的3个文件夹bin,pkg,src，额外包含一个project.txt文件，内容为项目名称
* 设置系统环境变量GOPATH 为 /Users/xxx/gowork/work
* 设置系统环境变量GOBIN 为 GOBIN=$GOPATH/bin  PATH=$PATH:$GOBIN

```javascript
  gpm命令列表
  将xxx项目目录设置为work: use xxx项目名
  清空work目录: clean
  创建新项目: create xxx
  退出命令行: exit
  帮助: help
```
### 操作步骤
	1、安装GO
	2、配置环境变量
		export GOROOT=/usr/local/go
		export GOPATH=/Users/xxx/gowork/work
		export GOBIN=$GOPATH/bin
		export PATH=$PATH:$GOBIN 
	3、创建目录
		首先gowork我们视为go根目录
		work为GOPATH目录
		bin pkg src
	4、创建文件
		xxx.go 此处 gopm.go
		/Users/xxx/gowork/work/src/gopm.go
	5、生成可执行文件
		cd src or cd bin
		go install gopm.go
	6、移动gopm可执行文件
		与work 并列
	7、在work目录下创建project.txt文件
		内容为目录名，首次不要与work相同，比如写mywork companywork	
	至此已基本完成接下来跑次流程
	8、cd .. 
		./gopm
		clean
		create loc
		use loc
ok 基本实现		
