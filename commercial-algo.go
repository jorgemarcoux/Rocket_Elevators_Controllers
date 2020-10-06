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
	smallestGap := 100
	for i :=0; i < len(c.listOfElevators); i++{
		diff := positive(requestedFloor - c.listOfElevators[i].currentFloor)
		if diff < smallestGap{
		  smallestGap = diff
		}
		fmt.Println("The smallest gap is:", smallestGap)
	  }
}
/*
func (c Column) getSmallerFloorsGap2(){
	smallestGap := c.listOfElevators[0].floorsGap
	for _, element := range c.listOfElevators{
		if smallestGap < element{
			smallestGap = element
		}
	}
}
*/
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
}








//battery1.listOfColumns[0].listOfElevators = append(battery1.listOfColumns[0].listOfElevators,colA.listOfElevators[0])
	//battery1.listOfColumns[0].listOfElevators = append(battery1.listOfColumns[0].listOfElevators,elevatorA2)
	//fmt.Println(battery1.listOfColumns[0].listOfElevators[0].floorsGap)
	//fmt.Println(battery1.listOfColumns[0].listOfElevators[1].floorsGap)
	//fmt.Println("Elevator 3 in Column A id is: ",elevatorA3.id)
	
	//elevatorA1.requestfloor(9) test requestFloor
	//battery1.listOfColumns[0].addElevatorsToColumn(5)
	//fmt.Println("Elevator 3 in Column A id is: ",elevatorA3.id)
	
	//elevatorA1.requestfloor(9) test requestFloor
	//battery1.listOfColumns[0].addElevatorsToColumn(5)
	//colA.listOfElevators = append(colA.listOfElevators, Elevator{"A1","idle","none",1,0})
	//colA.listOfElevators = append(colA.listOfElevators, Elevator{"A2","idle","none",1,0})
	//elevatorA2 := Elevator{"A2","idle","none",1,0}
	//elevatorA3 := Elevator{"A3","idle","none",1,0}
	//fmt.Println(columnB.floorsNumber)
	//fmt.Println(columnC.floorsNumber)
	//fmt.Println(columnD.floorsNumber)
	//columnB := Column{"B",20,1,20,5,[]Elevator{}}
	//columnC := Column{"C",21,1,40,5,[]Elevator{}}
	//columnD := Column{"D",21,1,60,5,[]Elevator{}}

	/* battery1.listOfColumns = append(battery1.listOfColumns, Column{"A",7,-6,1,5,[]Elevator{}})
	
	fmt.Println(battery1.listOfColumns[0].floorsNumber)
	
	colA := battery1.listOfColumns[0]
	
	
	fmt.Println("Elevator 1 in Column A is at: ",colA.listOfElevators[0].currentFloor)
	fmt.Println("Elevator 2 in Column A id is: ",colA.listOfElevators[1].id)
	
	colA.listOfElevators[0].currentFloor = 2
	colA.listOfElevators[1].currentFloor = 5
	
	fmt.Println(colA.listOfElevators)

	colA.getSmallerFloorsGap(10) */

	