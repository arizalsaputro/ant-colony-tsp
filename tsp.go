/**
*	Author : Muhammad Arizal Saputro
*	StudentId: 1301178325
*/

package main

import (
	"flag"
	"os"
	"fmt"
	"time"
	"github.com/arizalsaputro/ant-system-tsp/util"
	"github.com/arizalsaputro/ant-system-tsp/swarm"
)

func main() {
	// Sub commands
	runCommand := flag.NewFlagSet("run",flag.ExitOnError)


	// Calculate sub command flag pointer
	runTextFileName := runCommand.String("file","","File .tsp to read. (Required)")
	runIntIteration := runCommand.Int("iteration",5,"Number of iteration. (Optional)")
	runIntColonySize := runCommand.Int("ant",20,"Number of ant size. (Optional)")
	runFloatEvap := runCommand.Float64("evap",0.001,"Number of evaporation (1-p). (Optional)")
	runFloatAlpha := runCommand.Float64("alpha",15,"Alpha constant variable. (Optional)")
	runFloatBeta := runCommand.Float64("beta",20,"Beta constant variable. (Optional)")
	runFloatInitPherom := runCommand.Float64("init-pheromone",0.1,"Initial Pheromone. (Optional)")
	runFloatInitQ := runCommand.Float64("q",0.1,"Q constant variable. (Optional)")
	runBoolSetLog := runCommand.Bool("log",false,"Set logging. (Optional)")
	runIntAvgOf := runCommand.Int("avg-of",0,"Run n-times and get average")

	if len(os.Args) < 2 {
		fmt.Println("run subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("Unknown Command %v",os.Args[1])
		os.Exit(1)
	}

	if runCommand.Parsed(){
		if *runTextFileName == ""{
			runCommand.PrintDefaults()
			os.Exit(1)
		}
		start := time.Now()
		fmt.Println("Run...")
		list,err := util.ReadFile(*runTextFileName)
		if err != nil{
			fmt.Printf("Error read file '%v'",err.Error())
			os.Exit(1)
		}

		if *runIntAvgOf != 0{
			avg := 0.0
			for i:=0;i<*runIntAvgOf;i++{
				pso := swarm.NewAntCycle(5,20,0.001,15,20,0.1,0.1,list)
				pso.SetLog(false)
				pso.Process()
				avg += pso.BestAnt.TotalCost
				fmt.Printf("best cost %v %v \n",pso.BestAnt.TotalCost , pso.BestAnt.GetPath())
			}

			fmt.Println("average cost:",avg/float64(*runIntAvgOf))
		}else{
			pso := swarm.NewAntCycle(*runIntIteration,*runIntColonySize,*runFloatEvap,*runFloatAlpha,*runFloatBeta,*runFloatInitPherom,*runFloatInitQ,list)
			pso.SetLog(*runBoolSetLog)
			pso.Process()
			fmt.Printf("best cost -> %v \npath -> %v \n",pso.BestAnt.TotalCost , pso.BestAnt.GetPath())
		}

		elapsed := time.Since(start)
		fmt.Printf("Ant System took %s", elapsed)
	} else{

	}
}