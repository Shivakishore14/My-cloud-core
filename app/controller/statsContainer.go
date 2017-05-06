package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shivakishore14/My-cloud-core/app/console"
	"github.com/Shivakishore14/My-cloud-core/app/model"

	lxc "github.com/lxc/go-lxc"
)

//GetStatsContainer is for getting status of container
func GetStatsContainer(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, ok := session.Values["user"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}
	fmt.Print(user)

	name := r.FormValue("name")

	container := model.ContainerStat{}
	c, err := lxc.NewContainer(name, lxcpath)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		//change error part >> information exposture
		webresponse("Please Try again", err, nil, w)
	}
	isRunning := c.Running()
	container.IsRunning = isRunning
	container.Name = name
	fmt.Println(isRunning)
	if isRunning {
		memUsed, err := c.MemoryUsage()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("MemoryUsage: %s\n", memUsed)
			container.MemUsed = memUsed.String()
		}

		memLimit, err := c.MemoryLimit()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("MemoryLimit: %s\n", memLimit)
			container.MemLimit = memLimit.String()
		}

		// kmem
		kmemUsed, err := c.KernelMemoryUsage()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("KernelMemoryUsage: %s\n", kmemUsed)
			container.KernelMemUsed = kmemUsed.String()
		}

		kmemLimit, err := c.KernelMemoryLimit()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("KernelMemoryLimit: %s\n", kmemLimit)
			container.KernelMemLimit = kmemLimit.String()
		}

		// swap
		swapUsed, err := c.MemorySwapUsage()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("MemorySwapUsage: %s\n", swapUsed)
			container.MemSwapUsed = swapUsed.String()
		}

		swapLimit, err := c.MemorySwapLimit()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("MemorySwapLimit: %s\n", swapLimit)
			container.MemSwapLimit = swapLimit.String()
		}

		// blkio
		blkioUsage, err := c.BlkioUsage()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("BlkioUsage: %s\n", blkioUsage)
			container.BulkIOUsage = blkioUsage.String()
		}

		cpuTime, err := c.CPUTime()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("cpuacct.usage: %s\n", cpuTime)
			container.CPUTime = cpuTime.String()
		}
		cpuTimePerCPU, err := c.CPUTimePerCPU()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("cpuacct.usageerrpercpu: %v\n", cpuTimePerCPU)
			container.CPUTimePerCPU = cpuTimePerCPU
		}
		cpuStats, err := c.CPUStats()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("cpuacct.stat: %v\n", cpuStats)
			container.CPUStats = cpuStats
		}

		interfaceStats, err := c.InterfaceStats()
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("InterfaceStats: %v\n", interfaceStats)
			container.InterfaceStats = interfaceStats
		}
		//all ips
		i, _ := c.Interfaces()
		for _, k := range i {
			fmt.Println(k)
			fmt.Println(c.IPAddress(k))
		}
		//imp ip :D
		fmt.Println(c.IPAddresses())
		container.IPAddress, _ = c.IPAddresses()
		fmt.Println(c.ConfigFileName())
		for _, k := range c.ConfigKeys() {
			console.PrintSuccess(k)
		}

	} else {
		fmt.Print("Container is not running")
	}
	webresponse("Success", nil, container, w)
}
