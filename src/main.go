package main

import (
	"./object"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)




/**
	class文件解析
 */
func parserClass(){
	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\TestDemo1.class"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666);
	if err != nil {
		fmt.Println("Open file error!", err);
		return
	}

	stat, _ := file.Stat();

	fmt.Print(stat.Size());

	buf := make([]byte, stat.Size());

	//var data []byte;



	for {
		length, err := file.Read(buf);
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		//data = append(data,buf);
		fmt.Println(length, string(buf));
	}
	for _, val := range buf {
		H:=fmt.Sprintf("%016x", val);
		fmt.Printf("index slice= %s\n",  H)

	}
		x :=1

	fmt.Print(x)

}





func main() {



	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\TestDemo1.class"
	classN  := new(object.ClassByteFile);
	classN.Load(fileName);

	data := classN.Read(4)
	//能正常使用
	fmt.Println( 	binary.BigEndian.Uint32(data)==0xcafebabe )


	version 		:= classN.ReadShort();
	masterVersion 	:= classN.ReadShort();
	constanCount 	:= classN.ReadShort();
	fmt.Print( fmt.Sprintf("主版本%d,副版本%d,常量池%d\n",masterVersion,version,constanCount) );



	x :=1;fmt.Print(x)

	for{}

}


//func main() {
//	str := "月色真美，风也温柔，233333333，~！@#"  //go字符串编码为utf-8
//	fmt.Println("before convert:", str)   //打印转换前的字符串
//	fmt.Println("isUtf8:", isUtf8([]byte(str)))   //判断是否是utf-8
//	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))  //使用官方库将utf-8转换为gbk
//	fmt.Println("gbk直接打印会出现乱码:", string(gbkData))   //乱码字符串
//	fmt.Println("isGBK:", isGBK(gbkData)) //判断是否是gbk
//	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbkData) //将gbk再转换为utf-8
//	fmt.Println("isUtf8:", isUtf8(utf8Data) )  //判断是否是utf-8
//	fmt.Println("after convert:", string(utf8Data))   //打印转换前的字符串
//}