package main 

import (
	"fmt"
	"time"
)

func positive(n int) int {
    if n < 0 {
        n *= -1
    }
    return n
}

type Battery struct{
	id int
	numberOfColumns int
	listOfColumns []Column
}

func (b *Battery) createBattery(numberOfColumns int, _elevatorsInColumn int){
     for i :=0; i < numberOfColumns; i++{
		 b.listOfColumns = append(b.listOfColumns, Column{i+1,0,0,0,_elevatorsInColumn,[]Elevator{}})
		 b.listOfColumns[i].createColumn(_elevatorsInColumn)
	 }

}

type Column struct{
	id int
	floorsNumber int
	minimumFloor int
	maximumFloor int
	elevatorsInColumn int
	listOfElevators []Elevator
	//requestedFloorsQueue [] int
	
}

func (c *Column) ChangeColumnValues(_floorNumber int, _minimumFloor int,_maximumFloor int){
	c.floorsNumber = _floorNumber
	c.minimumFloor = _minimumFloor
	c.maximumFloor = _maximumFloor

}

func (c *Column) createColumn(numberOfElevators int){
   for i :=0; i < numberOfElevators; i++{
	   c.listOfElevators = append(c.listOfElevators, Elevator{i+1,"idle","none",1,0})
   }
}

func (c *Column) getSmallerFloorsGap(requestedFloor int){
	smallestGap := c.listOfElevators[0].floorsGap
	for i :=0; i < len(c.listOfElevators); i++{
		c.listOfElevators[i].floorsGap = positive(requestedFloor - c.listOfElevators[i].currentFloor)
		if c.listOfElevators[i].floorsGap < smallestGap{
		  smallestGap = c.listOfElevators[i].floorsGap
		}
		
	  }
	  fmt.Println("The smallest gap is:", smallestGap)
}

func (c Column) getIdleElevators(){
	var idleElevators = 0
	for i :=0; i < len(c.listOfElevators); i++{
		if c.listOfElevators[i].state == "idle"{
			idleElevators ++
			//TO DO: Append to new array IdleElevators to have access to those elevators
		}
	}
	fmt.Println("Idle elevators:", idleElevators)
}

func (c *Column) requestElevator(floorNumber int, direction string){

}

type Elevator struct{
	id  int
	state string
	direction string
	currentFloor int
	floorsGap int
	
} 

func (e *Elevator) requestfloor(requestedFloor int) {
	//falta push to array
	fmt.Println("REQUESTED FLOOR:#", requestedFloor)
	if requestedFloor > e.currentFloor {
		e.direction = "up"
		fmt.Println("Elevator",e.id,"is at floor #",e.currentFloor,"Direction =", e.direction)
		fmt.Println("Moving to floor #",requestedFloor)
		for e.currentFloor < requestedFloor{
			e.currentFloor +=1
			fmt.Println("Elevator",e.id,"is at floor #",e.currentFloor)
		} 
	} else if requestedFloor < e.currentFloor{
		e.direction = "down"
		fmt.Println("Elevator",e.id, "is at floor #", e.direction)
		fmt.Println("Moving to floor #",requestedFloor)
		for requestedFloor < e.currentFloor {
			e.currentFloor -=1
			fmt.Println("Elevator",e.id,"is at floor #", e.currentFloor)
		}
	} 
	if e.currentFloor == requestedFloor{
        fmt.Println("Elevator stopped");
        time.Sleep(100 * time.Millisecond)
        fmt.Println("Elevator",e.id, "arrived at target floor")
        time.Sleep(100 * time.Millisecond)
        fmt.Println("Opening doors...")
        time.Sleep(100 * time.Millisecond);
        fmt.Println("Person exits the elevator")
        time.Sleep(100 * time.Millisecond);
        fmt.Println("closing doors...")
        time.Sleep(100 * time.Millisecond);
	}

}//End requestElevator method

func main(){
	bat := &Battery{1,4,[]Column{}}
	fmt.Println("Battery1 id is",bat.id)
	bat.createBattery(4,5)
	bat.listOfColumns[0].ChangeColumnValues(7,-6,1)
	bat.listOfColumns[1].ChangeColumnValues(20,1,20)
	bat.listOfColumns[2].ChangeColumnValues(21,1,40)
	bat.listOfColumns[3].ChangeColumnValues(21,1,60)
	//fmt.Println(bat.listOfColumns)

	bat.listOfColumns[0].listOfElevators[0].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[1].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[2].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[3].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[4].currentFloor = 6
	fmt.Println(bat.listOfColumns[0].listOfElevators[0].currentFloor)
	bat.listOfColumns[0].getSmallerFloorsGap(10)

	bat.listOfColumns[0].listOfElevators[4].state = "moving"
	bat.listOfColumns[0].getIdleElevators()

	
}


//Future functions
/*
func (c Column) getIdleElevators(){
	var count = 0
	for i :=0; i < 6; i++{
		if c.listOfElevators[i].state == "idle"{
			count ++
		}
	}
	fmt.Println("Idle elevators:", count)
}
*/






	