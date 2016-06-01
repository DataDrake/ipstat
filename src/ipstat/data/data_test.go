package data

import (
	"testing"
)

func TestCollectSamples(t *testing.T){
	rtts,err := CollectSamples("129.21.171.72",100)
	if err != nil {t.Error(err.Error())}
	if len(rtts) != 50 {t.Error("Did not finish collecting samples")}
}

func TestCollectDataPoints(t *testing.T){
	rtts,err := CollectDataPoints("129.21.171.72",100,1500,100)
	if err != nil {t.Error(err.Error())}
	if len(rtts) != 50*15 {t.Error("Did not finish collecting samples")}
}

