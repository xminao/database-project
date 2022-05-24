/*
 * 定义包名，每一个源文件第一行都要指明文件属于哪个包
 * package main 表示一个可独立执行的程序
 * 每一个Go应用程序都包含一个名为main的包
 *
 * 文件名和包名没有直接关系，不可以将文件名和包名同名
 * 文件夹名与包名没有直接关系，不需要一直
 * 同一个文件夹下的文件只能有一个包名，否则编译会出错
 */
package main

// fmt包包含格式化I/O函数
import "fmt"

/*
 * 每程序开始执行的函数
 * main函数实每一个可执行程序必须包含的
 * 一般是在启动后第一个执行的函数（如果有init()则会先执行该函数）
 */
func main() {
	fmt.Println("第一个go程序")
}

/*
 * 标识符（常量、变量、类型、函数名、结构字段等）以一个大写字母开头，则这种形式的标识符的对象可以被外部包代码使用（如同面向对象的public），称为导出
 * 标识符以小写字母开头，则对包外是不可见的，但他们在整个包的内部是可见可用的（如同面向对象的protected）
 */
