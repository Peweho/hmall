package test

import (
	"bufio"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/api/internal/types"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadLua1(t *testing.T) {
	//读取lua脚本
	file, err := os.Open(types.Luapath)
	if err != nil {
		logx.Errorf("os.Open: %v,error: %v", types.Luapath, err)
	}
	reader := bufio.NewReader(file)
	var luaScript string
	for {
		line, err := reader.ReadString('\n')
		luaScript += line
		if err == io.EOF {
			break
		}
	}
	fmt.Println("entry:", luaScript)
}

func TestReadLua2(t *testing.T) {
	//读取lua脚本
	script, err := ioutil.ReadFile(types.Luapath)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(script))
}
