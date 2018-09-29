/**
*	Author : Muhammad Arizal Saputro
*	StudentId: 1301178325
*/

package util

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/arizalsaputro/ant-system-tsp/model"
)

func ParseCity(line []string)(*model.City,error)  {
	city := new(model.City)
	var err error
	city.ID,err = strconv.Atoi(line[0])
	if err != nil{
		return nil,err
	}
	city.X,err = strconv.ParseFloat(line[1], 64)
	if err != nil{
		return nil,err
	}
	city.Y,err =  strconv.ParseFloat(line[2], 64)
	if err != nil{
		return nil,err
	}
	return city,nil
}

func ReadFile(loc string)([]model.City,error){
	f, err := os.Open(loc)
	if err != nil {
		return nil,err
	}
	defer f.Close()

	var listOfCity []model.City

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		id := strings.Split(scanner.Text()," ")
		if len(id) == 3 {
			cty,err := ParseCity(id)
			if err == nil{
				listOfCity = append(listOfCity,*cty)
			}
		}
	}
	return listOfCity,nil
}
