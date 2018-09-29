package main

import (
	"time"
	"github.com/arizalsaputro/ant-system-tsp/util"
	"github.com/arizalsaputro/ant-system-tsp/swarm"
	"testing"
)

func TestTSP(t *testing.T)  {
	start := time.Now()
	/*list,err := util.ReadFile("Swarm016.tsp")
	if err != nil{
		log.Panic(err.Error())
	}

	avg := 0.0

	for i:=0;i<30;i++{
		pso := swarm.NewAntCycle(2,10,0.01,15,20,0.1,1,list)
		pso.SetLog(false)
		pso.Process()
		avg += pso.BestAnt.TotalCost
		fmt.Println("best cost",pso.BestAnt.TotalCost , pso.BestAnt.GetPath())
	}
	fmt.Println("average cost:",avg/30)*/

	list,err := util.ReadFile("Swarm096.tsp")
	if err != nil{
		t.Fatalf("Error read file %v",err.Error())
	}

	/*pso := swarm.NewAntCycle(1,96,0.001,15,20,0.1,0.1,list)
	pso.SetLog(true)
	pso.Process()
	if pso.BestAnt != nil {
		fmt.Println("best cost",pso.BestAnt.TotalCost , pso.BestAnt.GetPath())
	}*/

	avg := 0.0
	t.Log("Run...")
	for i:=0;i<1;i++{
		pso := swarm.NewAntCycle(5,20,0.001,15,20,0.1,0.1,list)
		pso.SetLog(false)
		pso.Process()
		avg += pso.BestAnt.TotalCost
		t.Logf("best cost %v %v \n",pso.BestAnt.TotalCost , pso.BestAnt.GetPath())
	}
	t.Log("average cost:",avg/30)

	elapsed := time.Since(start)
	t.Logf("Ant System took %s", elapsed)
}
