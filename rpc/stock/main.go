package main

import (
	"log"
	"net"
	stock "shoprpc/kitex_gen/example/shop/stock/stockservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
    addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
    svr := stock.NewServer(new(StockServiceImpl), server.WithServiceAddr(addr))

    err := svr.Run()

    if err != nil {
       log.Println(err.Error())
    }
}
