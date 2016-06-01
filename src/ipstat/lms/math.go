package lms
import(
	"math"
	"..//data"
)

func LMS(size []float64, rtt []float64) (float64,float64){
	sx := float64(0.0)
	sxx := float64(0.0)
	sy := float64(0.0)
	sxy := float64(0.0)
	for _,x := range size {
		sx += x
		sxx += math.Pow(x,2.0)
	}
	sx /= float64(len(size))
	sxx /= float64(len(size))
	sx2 := math.Pow(sx,2.0)
	for i,y := range rtt {
		sy += y
		sxy += y * size[i]
	}
	sy /= float64(len(size))
	sxy /= float64(len(size))
	slope := (sxy - sx*sy)/(sxx-sx2)
	intercept := sy - slope*sx
	return slope,intercept
}

func LMS_Perf(samples []data.PingResult) (float64,float64){
	sx := float64(0.0)
	sxx := float64(0.0)
	sy := float64(0.0)
	sxy := float64(0.0)
	for _,s := range samples {
		sx += s.Size
		sxx += math.Pow(s.Size,2.0)
		sy += s.RTT
		sxy += s.RTT * s.Size
	}
	sx /= float64(len(samples))
	sxx /= float64(len(samples))
	sx2 := math.Pow(sx,2.0)
	sy /= float64(len(samples))
	sxy /= float64(len(samples))
	slope := (sxy - sx*sy)/(sxx-sx2)
	intercept := sy - slope*sx
	return slope,intercept
}