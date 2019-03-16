package main

import (
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/armon/go-socks5"
	"github.com/caarlos0/env"
)

type params struct {
	User     string `env:"SOCKS5_USER" envDefault:""`
	Password string `env:"SOCKS5_PASS" envDefault:""`
	Port     string `env:"SOCKS5_PORT" envDefault:"1080"`
	Up       string `env:"SOCKS5_UP"   envDefault:""`
}

func main() {
	// Working with app params
	cfg := params{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	//Initialize socks5 config
	socsk5conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	if cfg.User+cfg.Password != "" {
		creds := socks5.StaticCredentials{
			cfg.User: cfg.Password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		socsk5conf.AuthMethods = []socks5.Authenticator{cator}
	}

	server, err := socks5.New(socsk5conf)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start listening proxy service on port %s\n", cfg.Port)

	if cfg.Up != "" {
		err = exec.Command(cfg.Up).Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
