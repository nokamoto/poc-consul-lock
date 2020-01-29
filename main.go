package main

import (
	consul "github.com/hashicorp/consul/api"
	"log"
)

func main() {
	log.Println("ready")

	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	key := "foo"
	value := []byte("value")
	pair := &consul.KVPair{
		Key:         key,
		Value:       value,
		ModifyIndex: 0,
	}

	kv := c.KV()

	log.Println("acquire")
	success, meta, err := kv.CAS(pair, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("success=%v meta=%v", success, *meta)
		log.Println("should be success")
	}

	log.Println("re-acquire (failed)")
	success, meta, err = kv.CAS(pair, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("success=%v meta=%v", success, *meta)
		log.Println("should not be success")
	}

	log.Println("get")
	got, meta2, err := kv.Get(key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("got=%v meta=%v", *got, *meta2)
	}

	log.Println("release (cas)")
	pair.ModifyIndex = meta2.LastIndex
	success, meta, err = kv.DeleteCAS(pair, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("success=%v meta=%v", success, *meta)
		log.Println("should be success")
	}

	log.Println("re-acquire")
	pair.ModifyIndex = 0
	success, meta, err = kv.CAS(pair, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("success=%v meta=%v", success, *meta)
		log.Println("should be success")
	}

	log.Println("release")
	meta, err = kv.Delete(key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("meta=%v", *meta)
	}

	log.Println("re-acquire")
	pair.ModifyIndex = 0
	success, meta, err = kv.CAS(pair, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("success=%v meta=%v", success, *meta)
		log.Println("should be success")
	}
}
