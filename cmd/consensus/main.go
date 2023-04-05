package main

import (
	"fmt"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/app"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("started to running subscriber lottery command")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	container := app.GetContainer()
	cfg := config.GetConfig()

	//TODO: create a state machine to control the flow
	go startSubscriber(container, cfg)
	go startRegister(container, cfg)

	select {
	case q := <-quit:
		log.Printf("finished process to subscribe peer on signal: %s", q)
	}
}

func checkPermission(c *app.Container) bool {
	log.Println("started plateaus validation")

	bal, err := c.PlateausValidationService.GetBalance()

	if err != nil {
		return false
	}

	return bal > 0
}

func startSubscriber(c *app.Container, cfg *config.Config) {
	for {
		log.Println("started lottery subscriber")

		if ok := checkPermission(c); !ok {
			log.Printf("was not able to join on lottery")
			<-time.After(1 * time.Minute)
			continue
		}

		networks := c.LotteryManager.CheckNetworkBalances(cfg.Networks)

		if len(networks) <= 0 {
			log.Printf("none networks to participate from lottery")
			<-time.After(1 * time.Minute)
			continue
		}

		if err := c.LotteryManager.SubscribePeer(cfg.Peer, networks); err != nil {
			log.Println(fmt.Sprintf("some error while c.LotteryManager.SubscribePeer: %s", err))
		}

		log.Printf("finished lottery subscriber [waiting for %d minutes]", 1)
		<-time.After(1 * time.Minute)
	}
}

func startRegister(c *app.Container, cfg *config.Config) {
	for {
		log.Println("started lottery processor")

		if ok := checkPermission(c); !ok {
			log.Printf("was not able to join on lottery")
			<-time.After(1 * time.Minute)
			continue
		}

		networks := c.LotteryManager.CheckNetworkBalances(cfg.Networks)

		if len(networks) <= 0 {
			log.Printf("none networks to participate from lottery")
			<-time.After(1 * time.Minute)
			continue
		}

		if err := c.LotteryManager.RegisterTx(cfg.Peer, networks); err != nil {
			log.Println(fmt.Sprintf("some error while app.m.Process: %s", err))
		}

		log.Printf("finished lottery register [waiting for %d minutes]", 1)
		<-time.After(1 * time.Minute)
	}
}
