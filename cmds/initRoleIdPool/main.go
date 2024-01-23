package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"zinx-zero/apps/acommon/globalkey"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
)

var start = 10002
var end = 9999999
var per = 0

func main() {
	// 获取可以被整除的
	var total = end - start + 1
	for i := 1000; i < 100000; i++ {
		if total%i == 0 {
			per = i
			fmt.Println(i)
			break
		}
	}
	if per == 0 {
		log.Fatal("per = 0")
	}
	// init_mysql()
	init_redis()
}
func init_redis() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",  // Redis 服务器地址
		Password: "G62m50oigInC30sf", // 密码（如果有的话）
		DB:       0,                  // 使用的数据库索引（默认是0）
	})

	// 连接数据库
	db, err := sql.Open("mysql", "root:PXDN93VRKUm8TeE7@tcp(localhost:33069)/gamex")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 执行查询操作 一次给 redis 插入100万条数据大约30M大小
	rows, err := db.Query("SELECT role_id FROM user_roleid_pool where is_use = 0 Limit 1000000")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var total int64
	// 遍历查询结果
	for rows.Next() {
		total++
		var role_id int64
		err := rows.Scan(&role_id) // 根据你的表结构调整列的顺序
		if err != nil {
			log.Fatal(err)
		}
		err = client.SAdd(context.Background(), globalkey.Cache_GenRoleId_UserIdPool, role_id).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	// 不够100万,池子数据不足
	if total < 1000000 {
		log.Fatal("不够100万,池子数据不足,已插入: ", total)
	}
	fmt.Println("数据插入完成！")
}

func init_mysql() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:PXDN93VRKUm8TeE7@tcp(localhost:33069)/gamex")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var roleIdMapList = make(map[int64]bool)
	// 循环插入数字到数据库中
	for i := start; i <= end; i++ {
		roleIdMapList[cast.ToInt64(i)] = true
	}
	// 准备SQL语句和数据
	// INSERT INTO user_roleid_pool (role_id) VALUES (10000),(10001)
	var execSql = "INSERT INTO user_roleid_pool (role_id) VALUES %v"
	var values []string
	var total int
	for key := range roleIdMapList {
		total++
		if total%per == 0 {
			fmt.Println(total)
			values = append(values, fmt.Sprintf("(%d)", key))
			execSql = fmt.Sprintf(execSql, strings.Join(values, ","))
			_, err = db.Exec(execSql)
			if err != nil {
				log.Fatal(err)
			}
			values = []string{}
			execSql = "INSERT INTO user_roleid_pool (role_id) VALUES %v"
		} else {
			values = append(values, fmt.Sprintf("(%d)", key))
		}
	}

	fmt.Println("数据插入完成！", "total:", end-start+1)
}
