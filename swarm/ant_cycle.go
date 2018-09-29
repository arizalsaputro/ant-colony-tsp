/**
*	Author : Muhammad Arizal Saputro
*	StudentId: 1301178325
*/

package swarm

import (
	"github.com/arizalsaputro/ant-system-tsp/model"
	"time"
	"math/rand"
	"gonum.org/v1/gonum/mat"
	"math"
	"log"
)


func init()  {
	rand.Seed(time.Now().Unix())
}

func random(max int) int {
	return rand.Intn(max)
}

type AntCycle struct {
	Iteration        int
	ColonySize       int
	Evaporation      float64
	DefaultPheromone float64
	Alpha            float64
	Beta             float64
	Q                float64
	Nodes            []model.City
	NumCities        int
	PheromoneMatrix  *mat.Dense
	DistanceMatrix   *mat.Dense
	AttractiveMatrix *mat.Dense
	BestAnt          *Ant
	Log              bool
}



//New instance for ant cycle tsp
func NewAntCycle(iteration, colonySize int,evaporation,alpha,beta,initPheromone,q float64,nodes []model.City)*AntCycle{
	return &AntCycle{
		Iteration:  iteration,
		ColonySize: colonySize,
		Evaporation:evaporation,
		Alpha:alpha,
		Beta:beta,
		DefaultPheromone:initPheromone,
		Nodes:nodes,
		NumCities:len(nodes),
		BestAnt:nil,
		Q:q,
	}
}

//set log status
func (c *AntCycle)SetLog(s bool){
	c.Log = s
}

// initialize pheromone
func (c *AntCycle) initializePheromone() {
	data := make([]float64,int(math.Pow(float64(c.NumCities + 1),2)))
	for i := range data {
		data[i] = c.DefaultPheromone
	}
	c.PheromoneMatrix = mat.NewDense(c.NumCities + 1,c.NumCities + 1,data)
}

// calculate the distance of vector between a <-> b
func (c *AntCycle) initializeDistance(){
	c.DistanceMatrix = mat.NewDense(c.NumCities + 1,c.NumCities + 1,nil)
	for _,i := range c.Nodes{
		for _,j:= range c.Nodes{
			c.DistanceMatrix.Set(i.ID,j.ID,i.DistanceTo(j))
		}
	}
}

// initialize attractiveness is 1/distance, use this to optimize computation
func (c *AntCycle) initializeAttractive(){
	c.AttractiveMatrix = mat.NewDense(c.NumCities + 1,c.NumCities + 1,nil)
	for _,i := range c.Nodes{
		for _,j:= range c.Nodes{
			if c.DistanceMatrix.At(i.ID,j.ID) != 0 {
				c.AttractiveMatrix.Set(i.ID,j.ID,1/c.DistanceMatrix.At(i.ID,j.ID))
			}
		}
	}
}

/*
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	log.Printf("Printing tabu\n%v\n", fa)
}
*/

// update pheromone after passing the path from a -> b
func (c *AntCycle) batchUpdatePheromone(ant *Ant,from,to model.City){
	for i := 0;i<c.NumCities+1;i++{
		for j := 0;j<c.NumCities+1;j++{
			deltaT := 0.0
			if (from.ID == i && to.ID == j) || (from.ID == j && to.ID == i){
				deltaT = c.Q / ant.TotalCost
			}
			decay := (1-c.Evaporation)*c.PheromoneMatrix.At(i,j) + deltaT
			c.PheromoneMatrix.Set(i,j,decay)
			if c.PheromoneMatrix.At(i,j) < 0 {
				c.PheromoneMatrix.Set(i,j,c.DefaultPheromone)
			}
		}
	}
	/*if c.Log {
		log.Println("[Ant System]: update pheromone")
		//matPrint(c.PheromoneMatrix)
	}*/
}

func (c *AntCycle)setBestAnt(ant *Ant){
	//
	if c.BestAnt == nil{
		if c.Log {
			log.Println("[Ant System]: Init best")
		}

		c.BestAnt = &Ant{TabuList:ant.TabuList,TotalCost:ant.TotalCost}
		return
	}
	if ant.TotalCost <  c.BestAnt.TotalCost {
		if c.Log {
			log.Printf("[Ant System]: New Best -> cost : %v",ant.TotalCost)
		}
		c.BestAnt = &Ant{TabuList:ant.TabuList,TotalCost:ant.TotalCost}
	}
}


// create ant agent then do tour
func (c *AntCycle) runCycle(){
	if c.Log {
		log.Printf("[Ant System]: Running Cycle")
	}
	listAnt := make([]Ant,c.ColonySize)
	for k:=0;k<c.ColonySize;k++{
		listAnt[k] = CreateNewAnt(c.Nodes[random(c.NumCities)])
	}
	for i := 0;i<c.Iteration;i++{
		if c.Log {
			log.Printf("[Ant System]: Iteration %d",i)
		}
		for _,ant := range listAnt{
			ant.Clear(c.Nodes)
			//ant.Random()
			ant.DoTour(c)
		}
	}
	if c.Log {
		log.Printf("[Ant System]: Cycle Done")
	}
}

// Run the ant cycle
// init all then run
func (c *AntCycle)Process(){
	c.initializePheromone()
	c.initializeDistance()
	c.initializeAttractive()
	c.runCycle()
}