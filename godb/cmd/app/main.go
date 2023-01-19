package main

import (
	"context"
	"dockerPostgres/godb/internal/godb"
	"dockerPostgres/godb/pkg/helpers/pg"
	"fmt"
	"os"
)

func main() {
	//Задаем параметры для подключения к БД (в прошлом задании мы поднимали контейнер с этими credentials)
	cfg := &pg.Config{}
	cfg.Host = "localhost"
	cfg.Username = "postgres"
	cfg.Password = "root"
	cfg.Port = "54320"
	cfg.DbName = "db_test"
	cfg.Timeout = 5
	//Создаем конфиг для пула
	poolConfig, err := pg.NewPoolConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool config error: %v\n", err)
		os.Exit(1)
	}
	//Устанавливаем максимальное количество соединений, которые могут находиться в ожидании
	poolConfig.MaxConns = 5
	//Создаем пул подключений
	c, err := pg.NewConnection(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")
	//Проверяем подключение
	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")
	ins := &godb.Instance{Db: c}
	ins.Start()
}
