package bazica

import (
	"github.com/phawitb/bazi-go/internal/fourpillars"
	"github.com/phawitb/bazi-go/internal/luckpillars"
	"github.com/phawitb/bazi-go/internal/ultis"
	"github.com/phawitb/bazi-go/model"
	"time"
)

func GetBaziChart(dateTime time.Time, loc *time.Location, gender int, prefixPath ...string) (*model.BaziChart, error) {
	var baziChart model.BaziChart

	var path string
	if len(prefixPath) != 0 {
		path = prefixPath[0]
	}

	fourPillar, passed, remaining, err := fourpillars.GetFourPillars(dateTime, loc, path)
	if err != nil {
		return nil, err
	}
	baziChart.FourPillar = ultis.GetLifeCycleFromFourPillar(fourPillar)

	lucksPillar, err := luckpillars.GetLuckPillars(fourPillar, gender, passed, remaining, dateTime, path)
	if err != nil {
		return nil, err
	}
	baziChart.LuckPillars = lucksPillar

	return &baziChart, nil
}
