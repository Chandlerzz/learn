package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"time"
)
const (
	USERNAME = "root"
	PASSWORD = "root"
    NETWORK  = "tcp"
	SERVER   = "0.0.0.0"
	PORT     = 3306
	DATABASE = "dbooks"
			)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就console.log(ose
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数)
	resp, err := http.Get("https://www.dbooks.org/hashtags/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	if err !=nil {
		fmt.Println(err)
	}

	doc.Find("*[href^='/hashtags/']").Each(func(i int, s *goquery.Selection){
		title := s.Text()
		fmt.Println(i,title)
	})
}
