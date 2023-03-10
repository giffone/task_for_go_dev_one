package config

import (
	"errors"
	"flag"
	"log"
)

// "mongodb://localhost:27017"

type Cfg struct {
	Addr, StorageAddr, UrlSalt string
}

func (c *Cfg) Read() error {
	flag.StringVar(&c.Addr, "addr", "", "port address")
	flag.StringVar(&c.StorageAddr, "storageAddr", "", "storageAddr")
	flag.StringVar(&c.UrlSalt, "urlSalt", "", "urlSalt")
	flag.Parse()

	return c.validate()
}

func (c *Cfg) validate() error {
	if c.Addr == "" {
		return errors.New("cfg: empty addr")
	}
	if c.StorageAddr == "" {
		return errors.New("cfg: empty storage addr")
	}
	if c.UrlSalt == "" {
		return errors.New("cfg: empty urlSalt")
	}
	log.Printf("\n[addr] %s\n[storageAddr] %s\n[urlSalt] %s\n", c.Addr, c.StorageAddr, c.UrlSalt)
	return nil
}
