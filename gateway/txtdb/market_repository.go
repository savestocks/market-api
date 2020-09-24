package txtdb

import (
	"encoding/json"
	"errors"
    "fmt"
	"log"
	"time"

    "github.com/andersonlira/market-api/domain"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/str"
	"sort"
)

//GetMarketList return all items 
func GetMarketList() []domain.Market {
	list := []domain.Market{}
    fileName := fmt.Sprintf("bd/%ss.json", "Market");
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

//GetMarketByID return all items 
func GetMarketByID(ID string) (domain.Market, error) {
	list := GetMarketList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			return list[idx],nil
		}
	}
	return domain.Market{}, errors.New("NOT_FOUND")
}



//SaveMarket saves a Market object
func SaveMarket(it domain.Market) domain.Market {
	list := GetMarketList()
	it.ID = str.NewUUID()
	it.CreatedAt = time.Now()
	list = append(list, it)
	writeMarket(list)
	return it
}

//UpdateMarket( updates a Market object
func UpdateMarket(ID string, it domain.Market) domain.Market{
	list := GetMarketList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].UpdatedAt = time.Now()
			writeMarket(list)
			return list[idx]
		}
	}
	return it
}

//DeleteMarket delete object by giving ID
func DeleteMarket(ID string) bool {
	list := GetMarketList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writeMarket(list)
			return true
		}
	}
	return false
}

func writeMarket(list []domain.Market) {
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%ss.json", "Market"), string(b))
}

