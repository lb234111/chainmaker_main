// coding:utf-8
// 日志
package utils

import (
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 在控制台打印出不同颜色的信息
// 0 - 黑色
// 1 - 红色
// 2 - 绿色
// 3 - 黄色
// 4 - 蓝色
// 5 - 紫红色
// 6 - 青蓝色;
func ColorPrint(color int, message string) {
	fmt.Printf("\033[0;%dm%s\033[0m\n", color+30, message)
}

// 打印输出一条错误消息
// @param funcName 哪一个函数名报错, 自己指定
// @param err 错误信息
func InfoError(funcName string, err error) string {
	info := fmt.Sprintln("funcName->" + funcName + " error->" + err.Error())
	ColorPrint(1, info)
	return info
}

// 打印输出消息
// @param info 任意消息
// @return 合并后的消息
func Info(info ...any) string {
	result := fmt.Sprint("info: ", info)
	ColorPrint(2, result)
	return result
}

// 打印输出消息, 不加前缀
// @param info 任意消息
// @return 合并后的消息
func Print(info ...any) string {
	result := fmt.Sprint(info...)
	ColorPrint(0, result)
	return result
}

// 打印警告消息
// @param info 任意消息
// @return 合并后的消息
func InfoWarning(info ...any) string {
	result := fmt.Sprint("warning: ", info)
	ColorPrint(3, result)
	return result
}

// 打印提示消息
// @param info 任意消息
// @return 合并后的消息
func InfoTips(info ...any) string {
	result := fmt.Sprint("tips: ", info)
	ColorPrint(4, result)
	return result
}

// 自定义的log函数
// @param funcName 是哪一个函数出来的log
// @param resp 合约返回的response
// @param err 其他错误
// @param return 错误信息
func MyLogRespErr(funcName string, resp *common.TxResponse, err error) string {
	info := "=== " + funcName + " ==="
	// fmt.Println(resp)
	if resp != nil {
		info += "\nresp.Message->" +
			resp.Message + "\nresp.ContractsResult->" +
			resp.ContractResult.String()
	}
	if err != nil {
		info += ("\nerror->" + err.Error())
	}
	info += "\n"
	fmt.Println(info)
	return info
}

// 自定义的log函数
// @param funcName 是哪一个函数出来的log
// @param resp 合约返回的response
// @param err 其他错误
// @param return 错误信息
func MyLogPayloadErr(funcName string, payload *common.Payload, err error) string {
	info := "=== " + funcName + " ==="
	// fmt.Println(resp)
	if payload != nil {
		info += "\npayload.string->" +
			payload.String()
	}
	if err != nil {
		info += ("\nerror->" + err.Error())
	}
	info += "\n"
	fmt.Println(info)
	return info
}
