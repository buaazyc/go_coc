package scene

import (
	"go_coc/dao"
	"go_coc/parser"
	"go_coc/time"
	"log"
)

// SumWar 部落战汇总
func SumWar(clan string, season string) (*parser.SumWarRsp, error) {
	wars, err := queryMonthWars(clan, season)
	if err != nil {
		return nil, err
	}
	_ = wars
	// log.Print(wars)
	// for _, r := range wars {
	// 	log.Printf("%+v", r)
	// }
	return &parser.SumWarRsp{}, nil
}

// queryMonthWars 从数据库中获取部落整个月的战绩
func queryMonthWars(clan string, season string) ([]string, error) {
	wars, err := dao.QueryAllWarsFor(clan)
	if err != nil {
		return []string{}, err
	}
	var res []string
	for _, war := range wars {
		monthTime, err := time.TimeToMonth(war.Time)
		if err != nil {
			log.Print(err)
			continue
		}
		if monthTime == season {
			res = append(res, war.Info)
		}
	}
	return res, nil
}
