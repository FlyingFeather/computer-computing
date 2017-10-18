# service-computing-selpg

## 实验介绍
使用 golang 开发 开发 Linux 命令行实用程序 中的 selpg（GO语言）

提示：
1. 请按文档 使用 selpg 章节要求测试你的程序
2. golang 文件读写、读环境变量，请自己查 os 包
3. golang stdin 读，参考这里
4. 应用程序只能修改自己的环境变量，不能修改父程序的环境变量

main.go为程序代码。

## 参数介绍
1. -s : 指定打印的起始页码
2. -e ：指定打印的结束页码
3. -d ：指定打印的位置——此实验利用其他命令（cat）代替
4. -l ：打印每页固定的行数
5. -f ：打印的每页行数不固定

没有-d则默认将输出打印到终端，默认长度为72


## 说明
参考下面的c语言程序，学习go语言。

[c语言程序链接](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c)

## 使用与测试结果
 1. ./selpg -s1 -e1 file1

该命令将把“input_file”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。由于没有指定行号，默认一页72行。

![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/1.png)
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/2.png)

2.  ./selpg -s1 -e1 < file1

![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/4.png)
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/5.png)


3.  cat main.go | ./selpg -s1 -e1
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/8.png)


4. ./selpg -s1 -e1 file1 > outputfile
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/6.png)
outputfile文件内容如对应文件中所示。

5. ./selpg -s500 -e200 file1 > error_file
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/7.png)
error_file文件内容如对应文件中所示。

6. ./selpg -s1 -e2 f file1 >outputfile2>error_file2

selpg 将第1页到第2页写至标准输出，标准输出被重定向至“outputfile2”；selpg 写至标准错误的所有内容都被重定向至“error_file2”。因为有错误所以outputfile2为空。

![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/9.png)

输出文件内容展示：
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/10.png)
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/11.png)
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/12.png)



7. ./selpg -s10 -e20 file1 >output_file3 2>/dev/null 
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/13.png)


8. ./selpg -s1 -e2 f file1 >/dev/null 
没有输出usage的内容
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/14.png)


9. ./selpg -s1 -e2 file1 | cat 
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/15.png)


10. ./selpg -s1 -e2 file1>error_file3 | cat 
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/16.png)


11. ./selpg -s1 -e2 -l3 file1
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/17.png)

12. ./selpg -s1 -e2 -f file1
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/18.png)
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/19.png)

13. ./selpg -s1 -e2 -dlp1 file2
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/20.png)
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/21.png)


### 其他错误命令展示
- ./selpg -s-3 -e-2 -l3

> panic: ./selpg: invalid start page -3


- ./selpg -s1 -e1 -l3 "no"

USAGE: ./selpg -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]
panic: ./selpg: input file "no" does not exist


- ./selpg -s1 -e2 -l3 < file1
![](https://github.com/FlyingFeather/service-computing/blob/master/selpg/screenshot/3.png)



- $ selpg -s10 -e20 input_file >output_file 2>/dev/null 

错误信息被丢弃


## end


