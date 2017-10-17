# service-computing-selpg

## 实验介绍
使用 golang 开发 开发 Linux 命令行实用程序 中的 selpg（GO语言）

提示：
1. 请按文档 使用 selpg 章节要求测试你的程序
2. golang 文件读写、读环境变量，请自己查 os 包
3. golang stdin 读，参考这里
4. 应用程序只能修改自己的环境变量，不能修改父程序的环境变量

## 设计说明
[c语言链接](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c)

## 使用与测试结果
 - ./selpg -s1 -e1 file1
该命令将把“input_file”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。由于没有指定行号，默认一页72行。
截图1
截图2

-  ./selpg -s1 -e1 < file1
截图4
截图5

-  ./selpg -s-3 -e-2 -l3
> panic: ./selpg: invalid start page -3

-  ./selpg -s-2 -e-1 -l3
> panic: ./selpg: invalid start page -2

-  ./selpg -s1 -e1 -l3 "no"
> USAGE: ./selpg -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]
> panic: ./selpg: input file "no" does not exist


-  ./selpg -s1 -e2 -l3 < file1
截图3
> the first file-1
> the first file-2
> the first file-3
> the first file-4
> the first file-5
> the first file-6

-  ./selpg -s1 -e1 file1 > outputfile
截图6
> the first file-1
> the first file-2
> the first file-3
> ...
> the first file-100

- ./selpg -s500 -e200 file1 > error_file
截图6

- $ selpg -s10 -e20 input_file >output_file 2>/dev/null 
错误信息被丢弃



