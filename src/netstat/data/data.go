package data
import (
	"os/exec"
	"regexp"
	"strconv"
	"log"
)

var PING_MATCH = regexp.MustCompile("rtt=(\\S*)")

type PingResult struct{
	RTT  float64
	Size float64
}

func CollectSamples(ip string, size int) ([]PingResult,error){
	samples := make([]PingResult,0)
	output,err := exec.Command("/usr/sbin/hping3","-1","-n","--fast","-c",strconv.Itoa(100),"-d", strconv.Itoa(size - 28) , ip).Output()
	for _,v := range PING_MATCH.FindAllStringSubmatch(string(output),-1) {
		rtt,err := strconv.ParseFloat(v[1],64)
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			samples = append(samples,PingResult{rtt/float64(1000),float64(size*8)})
		}
	}
	return samples,err
}

func CollectDataPoints(ip string, minSize int, maxSize int, delta int) ([]PingResult,error){
	results := make([]PingResult,0)

	for i := minSize; i <= maxSize; i+=delta {
		rtts,err := CollectSamples(ip,i)
		if err != nil {
			return results,err
		}
		results = append(results,rtts...)
	}
	return results,nil
}