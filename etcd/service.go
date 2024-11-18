package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Service interface {
	Export() (map[string]string, error)
}

type impl struct {
	Endpoints   []string
	DialTimeout time.Duration
}

func Define(endpoints []string, dialTimeout time.Duration) Service {
	return impl{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	}
}

func (i impl) Export() (map[string]string, error) {
	// 创建etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   i.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Error creating etcd client: %v", err)
	}
	defer cli.Close()
	// 获取所有键值对
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := cli.Get(ctx, "", clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("Error getting keys: %v", err)
	}
	// 创建一个map来存储键值对
	kvMap := make(map[string]string)
	for _, kv := range resp.Kvs {
		kvMap[string(kv.Key)] = string(kv.Value)
	}
	return kvMap, nil
}
