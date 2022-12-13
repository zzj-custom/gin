package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	houhaiStation := NewSubwayStation("后海站")
	fmt.Println(EnhanceEnterStationProcess(houhaiStation, false, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(houhaiStation, true, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(houhaiStation, true, true).Enter())
}

// EnhanceEnterStationProcess 根据是否有行李，是否处于疫情，增加进站流程
func EnhanceEnterStationProcess(station Station, hasLuggage bool, hasEpidemic bool) Station {
	if hasLuggage {
		station = NewSecurityCheckDecorator(station)
	}
	if hasEpidemic {
		station = NewEpidemicProtectionDecorator(station)
	}
	return station
}
