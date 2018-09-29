/**
*	Author : Muhammad Arizal Saputro
*	StudentId: 1301178325
*/

package swarm

import (
	"github.com/arizalsaputro/ant-system-tsp/model"
	"math/rand"
	"time"
	"math"
	"log"
	"fmt"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}


type Ant struct {
	InitialCity model.City
	CurrentCity model.City
	CityList []model.City
	TotalCity int
	TabuList []model.City
	TotalCost float64
}

// creating agent
func CreateNewAnt(city model.City)Ant{
	return Ant{
		InitialCity:city,
		CurrentCity:city,
		TotalCost:0,
	}
}

// clearing tabulist and
func (a *Ant)Clear(list []model.City){
	a.TabuList = a.TabuList[:0]
	a.TotalCost = 0
	a.TotalCity = len(list)
	mL := make([]model.City,a.TotalCity)
	copy(mL,list)
	a.CityList = mL
	a.CurrentCity = a.InitialCity
}

// to random initial city again
func (a *Ant)Random(){
	rand.Seed(time.Now().UnixNano())
	a.InitialCity = a.CityList[random(a.TotalCity)]
	a.CurrentCity = a.InitialCity
}

// get ant path
func (a *Ant) GetPath()string{
	s := ""
	for idx,i := range a.TabuList {
		if idx == len(a.TabuList)-1{
			s = fmt.Sprintf("%s%v",s,i.ID)
		}else{
			s = fmt.Sprintf("%s%v-",s,i.ID)
		}
	}
	return s
}

//Add city to tabu list
func (a *Ant) addCityToTabuList(city model.City){
	a.TabuList = append(a.TabuList,city)
}

// remove current city from list of destination city
func (a *Ant) normalizeCity(){
	for i:=0;i<len(a.CityList);i++{
		if a.CityList[i].ID == a.CurrentCity.ID{
			a.CityList = a.CityList[:i+copy(a.CityList[i:], a.CityList[i+1:])]
			break
		}
	}
}

// calculate of function (Tij^a*nij^b)
func (a *Ant) calculatePheromoneAttraction(c *AntCycle,from,to model.City)float64{
	return math.Pow(c.PheromoneMatrix.At(from.ID,to.ID),c.Alpha) * math.Pow(c.AttractiveMatrix.At(from.ID,to.ID),c.Beta)
}

// calculate the sum of (Tij^a*nij^b) from destination city
func (a *Ant) batchCalculate(c *AntCycle)float64{
	var sum float64
	for _,item := range a.CityList{
		sum += a.calculatePheromoneAttraction(c,a.CurrentCity,item)
	}
	return sum
}

// move from current city to destination city
// update pheromone path
func (a *Ant) moveCity(c *AntCycle,city model.City){
	/*if c.Log {
		log.Printf("[Ant Agent]: %d -> %d",a.CurrentCity.ID,city.ID)
	}*/
	a.addCityToTabuList(city)
	a.TotalCost += c.DistanceMatrix.At(a.CurrentCity.ID,city.ID)
	c.batchUpdatePheromone(a,a.CurrentCity,city)
	a.CurrentCity = city
	a.normalizeCity()
}

// ant do tour
// visit all city by calculating the probability
// update pheromone if ant visit the path
func (a *Ant)DoTour(c *AntCycle){
	if c.Log {
		log.Printf("[Ant Agent]: start-> %d",a.InitialCity.ID)
	}
	a.addCityToTabuList(a.CurrentCity)
	a.normalizeCity()
	stuck := 0
	for{
		if len(a.TabuList) == a.TotalCity{
			a.moveCity(c,a.InitialCity)
			c.setBestAnt(a)
			if c.Log {
				log.Printf("[Ant Agent]: End -> cost:  %v ->tour %s",a.TotalCost,a.GetPath())
			}
			break
		}
		r := rand.Float64()
		sum := a.batchCalculate(c)
		for _,to := range a.CityList{
			p := a.calculatePheromoneAttraction(c,a.CurrentCity,to) / sum
			if p >= r  {
				a.moveCity(c,to)
				stuck = 0
				break
			}
		}
		stuck++
		if stuck > 200 {
			if c.Log{
				log.Printf("{stuck} this ant get stuck,kill this ant")
			}
			break
		}
	}
}

