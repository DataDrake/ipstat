package main
import (
	"os"
	"flag"
	"fmt"
	"ipstat/data"
	"ipstat/lms"
)

func main(){
	show_samples := flag.Bool("S",false,"Display the sample data")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: ipstat [-S] <IP>")
		os.Exit(1)
	}
	samples,err := data.CollectDataPoints(args[0],100,1500,100)
	if err != nil {
		fmt.Printf("Failed to collect samples, reason: %s\n", err.Error())
		os.Exit(2)
	}
	if *show_samples {
		for i,v := range samples {
			fmt.Printf("%d,%f,%f\n",i,v.RTT,v.Size)
		}
	}

	slope,intercept := lms.LMS_Perf(samples)
	fmt.Printf("Speed: %fbps, Propagation Delay: %fs\n",float64(1.0)/slope,intercept)
	return
}