package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	_ "reptile_book/model"
	"reptile_book/mysql"
	_ "reptile_book/mysql"
	"strconv"
)

var ctx = context.Background()

func QueryBookTitleSort(autherId int, bookId int, bookName string) int {
	redisClient := Rd
	bookTitleSortKey := strconv.Itoa(autherId) + ":" + strconv.Itoa(bookId) + ":" + bookName
	val, err := redisClient.HGet(ctx, bookTitleSortKey, "sort").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bookTitleSortKey, val)

	sort, err := strconv.Atoi(val)

	if err == redis.Nil { // key不存在
		mysql.QueryBookTitleSort()
	} else if err != nil {
		fmt.Println("sort不是数字")
		panic(err)
	}

	return sort
}
