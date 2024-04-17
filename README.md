# chainmaker_study
* 这是一个长安链的基本使用入门仓库
* 实现一键部署并包含有基本的合约，帮助快速入门

## 快速启动
* `35`和`31`服务器已经配置好了环境，可以直接启动
* 其他机器需要进行环境部署,详情见[环境部署.md](./%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE.md) 和 [环境部署章节](#环境部署)
---
* 首先启动链并编译部署合约
```bash
bash fastrun.sh
```
* 然后启动测试代码,能看到很多输出,表示成功
```bash
go run main.go
```
---
* 看代码的话,就从`main.go`和`deploy/main.go`开始看就行了
* 具体的功能实现都在`chainmaker`文件夹下
* `deploy/main.go`是部署链码的代码,在`install_deploy_contract.sh`里面调用的
* `main.go`是与链交互的测试代码,直接执行可以看到返回结果
---
* 重启链
```bash
bash ./restart.sh
```
* 停止链
```bash
bash ./stop.sh
```

## 环境部署
* 需要[申请长安链仓库账号](https://git.chainmaker.org.cn/users/sign_in?redirect_to_referer=yes),才能下载长安链源代码(执行`make`指令会用到)
* [长安链官方环境搭建](https://docs.chainmaker.org.cn/v2.2.1/html/tutorial/%E9%80%9A%E8%BF%87%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7%E5%90%AF%E5%8A%A8%E9%93%BE.html)
* **centos**长安链环境部署详情请看[环境部署.md](./%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE.md)

## 目录结构说明
* `chainmaker`封装的长安链SDK，里面的文件名就是对应的功能
* `config`用到的配置，还有一些常量(`go`语言)
* `config_files`配置文件
* `contracts`全部的合约
* `deploy`部署合约
* `scripts`用到的全部`shell脚本`
* `utils`工具代码
  * 推荐看一下里面的`logger.go`实现了控制台彩色打印

## 合约编写
* 长安链官网的[go语言编写智能合约](https://docs.chainmaker.org.cn/instructions/%E4%BD%BF%E7%94%A8Golang%E8%BF%9B%E8%A1%8C%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6%E5%BC%80%E5%8F%91.html)
* 也可以参考`contracts`目录的合约写法
* 长安链也支持其他语言编写合约,请看[长安链官网智能合约编写](https://docs.chainmaker.org.cn/instructions/%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6%E5%BC%80%E5%8F%91.html)

## 切换公钥模式
* 将`Makefile`的`第6行`注释掉，并打开`第8行`的注释，然后`make`就是启动密钥模式
```bash
make
```
* 然后编译部署合约
```bash
bash ./scripts/build_contract.sh
go run deploy/main.go -m pk
```
* 运行`main.go`查看运行结果
```bash
go run main.go -m pk
```

## 其他
* [长安链cmc命令行工具](https://docs.chainmaker.org.cn/v2.2.1/html/dev/%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7.html)
* [长安链多机部署](https://docs.chainmaker.org.cn/v2.2.1/html/operation/%E5%A4%9A%E6%9C%BA%E9%83%A8%E7%BD%B2.html)
* [龙博同学的毕业寄语](http://10.21.4.21/longbo/worksummary)


## 贡献者
* [冉东川同学](http://10.21.4.21/rangdongchuan):项目发起者
* [龙博同学](http://10.21.4.21/longbo):长安链启动脚本
* [陈兴同学](http://10.21.4.21/chenxing):**centos**下环境配置

## 寄语
* 欢迎大家共同更新、维护仓库，并在**贡献者**一栏加上自己的名字
* 你们[亲爱的川川大宝贝](http://10.21.4.21/rangdongchuan)希望大家互帮互助, 一起进步, 做一个对家庭、国家、社会、世界人民有用的人