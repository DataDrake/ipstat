package netperf

import (
	"testing"
	"math"
)

func TestLMS1(t *testing.T) {
	size := []float64{100.0,100.0,100.0,100.0,100.0}
	rtt := []float64{0.046,0.0449,0.0439,0.0428,0.0507}
	m,b := LMS(size,rtt)
	if !math.IsNaN(m) {t.Error("Slope should be NaN")}
	if !math.IsNaN(b) {t.Error("Intercept should be NaN")}
}

func TestLMS2(t *testing.T) {
	size := []float64{100.0,200.0,300.0,400.0,500.0,600.0,700.0,800.0,900.0,1000.0}
	rtt := []float64{0.045725,0.046599,0.050748,0.056776,0.047153,0.047296,0.049783,0.049109,0.048043,0.051660}
	m,b := LMS(size,rtt)
	if !math.IsNaN(m) {t.Errorf("Slope should be NaN, found: %f\n",m)}
	if !math.IsNaN(b) {t.Errorf("Intercept should be NaN,found: %f\n",b)}
	t.Logf("Average Bandwidth: %fbps\n",float64(8.0)/m)
}