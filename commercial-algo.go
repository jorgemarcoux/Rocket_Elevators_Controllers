package main 

import (
	"fmt"
	"time"
	
)

var chosenElevator Elevator

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


func (c *Column) getSmallerFloorsGap(array []Elevator, floorNumber int){
	smallestGap := 100
	var chosenElevator = 1
	for i :=0; i < len(array); i++{
		array[i].floorsGap = positive(floorNumber - array[i].currentFloor)
		if array[i].floorsGap < smallestGap{
		  smallestGap = array[i].floorsGap
		  chosenElevator = array[i].id
		}
		
	  }
	  fmt.Println("Smallest floors gap:", smallestGap,".Elevator sent:",chosenElevator)
}


func (c Column) getIdleElevators(){
	var idleElevators = 0
	var idleElevList []Elevator
	for i :=0; i < len(c.listOfElevators); i++{
		if c.listOfElevators[i].state == "idle"{
			idleElevators ++
			idleElevList = append(idleElevList,c.listOfElevators[i])
		}
	}
	fmt.Println("Idle elevators:",idleElevators,"List of idle elevators:" ,idleElevList)
}

func (c *Column) requestElevator(floorNumber int, direction string){
	var elevInSameFloor []Elevator
	var elevatorsGoingUp []Elevator
	var elevatorsGoingDown []Elevator
	for i :=0; i < len(c.listOfElevators); i++{
		//Chcking if there's an elevator already at floorNumber going to same direction
	   if c.listOfElevators[i].currentFloor == floorNumber && c.listOfElevators[i].direction == direction{
			elevInSameFloor = append(elevInSameFloor,c.listOfElevators[i])
			fmt.Println("Scenario - call from floor #",floorNumber, "going", direction, "should send elevator",elevInSameFloor[0].id)
			fmt.Println("Elevator #",elevInSameFloor[0].id, "is already at floor #", floorNumber)
			elevInSameFloor[0].moveElevator(floorNumber)
		//Checking elevators going in same direction as requested direction	
	   }else if c.listOfElevators[i].direction == "up" && direction == "up" && c.listOfElevators[i].currentFloor <= floorNumber{
			elevatorsGoingUp = append(elevatorsGoingUp,c.listOfElevators[i])
			c.getSmallerFloorsGap(elevatorsGoingUp,floorNumber)
			fmt.Println("Scenario - call from floor #",floorNumber, "going", direction)
			fmt.Println("Elevator going up and requested direction up")
			elevatorsGoingUp[0].moveElevator(floorNumber)
	   }else if c.listOfElevators[i].direction == "down" && direction == "down" && c.listOfElevators[i].currentFloor >= floorNumber{
		    elevatorsGoingDown = append(elevatorsGoingDown,c.listOfElevators[i])
		    c.getSmallerFloorsGap(elevatorsGoingDown,floorNumber)
		    fmt.Println("Scenario - call from floor #",floorNumber, "going", direction)
			fmt.Println("Elevator going down and requested direction down. Elevator's current floor:",elevatorsGoingDown[0].currentFloor)
			elevatorsGoingDown[0].moveElevator(floorNumber)
	 //Prioritizing moving elevator versus idle one
	   }else if c.listOfElevators[i].state == "moving"{

	   }
	}//Closing for iterating c.listOfElevators
	

	
}//Closing requestElevator

type Elevator struct{
	id  int
	state string
	direction string
	currentFloor int
	floorsGap int
	
} 

func (e Elevator) moveElevator(requestedFloor int){
	if requestedFloor > e.currentFloor{
		e.direction = "up"
		fmt.Println("Elevator direction:", e.direction)
		for e.currentFloor < requestedFloor {
			e.currentFloor += 1
			fmt.Println("Elevator", e.id, "is at floor #", e.currentFloor)
		}
	}else if requestedFloor < e.currentFloor{
		e.direction = "down"
		fmt.Println("Elevator direction:", e.direction)
		for e.currentFloor > requestedFloor{
			e.currentFloor -= 1
			fmt.Println("Elevator", e.id, "is at floor #", e.currentFloor)
		}
	if e.currentFloor == requestedFloor{
		fmt.Println("Elevator stopped")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Elevator",e.id,"arrived at target floor")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Opening doors")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Person enters the elevator")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Closing doors...")



	}

	}
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
	//fmt.Println("Battery1 id is",bat.id)
	bat.createBattery(4,5)
	bat.listOfColumns[0].ChangeColumnValues(7,-6,1)
	bat.listOfColumns[1].ChangeColumnValues(20,1,20)
	bat.listOfColumns[2].ChangeColumnValues(21,1,40)
	bat.listOfColumns[3].ChangeColumnValues(21,1,60)
	//fmt.Println(bat.listOfColumns)
/* Test Smallest Floor Gap
	bat.listOfColumns[0].listOfElevators[0].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[1].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[2].currentFloor = 6
	bat.listOfColumns[0].listOfElevators[3].currentFloor = 2
	bat.listOfColumns[0].listOfElevators[4].currentFloor = 8
	//fmt.Println(bat.listOfColumns[0].listOfElevators)
	//bat.listOfColumns[0].getSmallerFloorsGap(10)
*/	


//******************************testing requestElevato***********************************
	//bat.listOfColumns[0].listOfElevators[2].direction = "down"
	//bat.listOfColumns[0].requestElevator(5, "down")
//****************SCENARIO 1*************************************
/*
	bat.listOfColumns[1].listOfElevators[0].currentFloor = 20
	bat.listOfColumns[1].listOfElevators[0].direction = "down"
	bat.listOfColumns[1].listOfElevators[1].currentFloor = 3
	bat.listOfColumns[1].listOfElevators[1].direction = "up"
	bat.listOfColumns[1].listOfElevators[2].currentFloor = 13
	bat.listOfColumns[1].listOfElevators[2].direction = "down"
	bat.listOfColumns[1].listOfElevators[3].currentFloor = 15
	bat.listOfColumns[1].listOfElevators[3].direction = "down"
	bat.listOfColumns[1].listOfElevators[4].currentFloor = 6
	bat.listOfColumns[1].listOfElevators[4].direction = "down"
	bat.listOfColumns[1].requestElevator(1, "up")
*/
//*********SCENARIO 2 WORKIN PROPERLY******************************************************************************
 /*
   bat.listOfColumns[2].listOfElevators[0].currentFloor = 1
   bat.listOfColumns[2].listOfElevators[0].direction = "up"
   bat.listOfColumns[2].listOfElevators[1].currentFloor = 23
   bat.listOfColumns[2].listOfElevators[1].direction = "up"
   bat.listOfColumns[2].listOfElevators[2].currentFloor =33
   bat.listOfColumns[2].listOfElevators[2].direction = "down"
   bat.listOfColumns[2].listOfElevators[3].currentFloor = 40
   bat.listOfColumns[2].listOfElevators[3].direction = "down"
   bat.listOfColumns[2].listOfElevators[4].currentFloor = 39
   bat.listOfColumns[2].listOfElevators[4].direction = "down"
   bat.listOfColumns[2].requestElevator(1, "up")
*/

//********SCENARIO 3 WORKING PROPERYL - BUT REPEATS PROCESS TWICE********************************************

  bat.listOfColumns[3].listOfElevators[0].currentFloor = 58
  bat.listOfColumns[3].listOfElevators[0].direction = "down"
  bat.listOfColumns[3].listOfElevators[1].currentFloor = 50
  bat.listOfColumns[3].listOfElevators[1].direction = "up"
  bat.listOfColumns[3].listOfElevators[2].currentFloor = 46
  bat.listOfColumns[3].listOfElevators[2].direction = "up"
  bat.listOfColumns[3].listOfElevators[3].currentFloor = 1
  bat.listOfColumns[3].listOfElevators[3].direction = "up"
  bat.listOfColumns[3].listOfElevators[4].currentFloor = 60
  bat.listOfColumns[3].listOfElevators[4].direction = "down"
  bat.listOfColumns[3].requestElevator(54, "down")

//********SCENARIO 4 WORKING PROPERLY - Missing opening doors at target floor***************************
/*
	bat.listOfColumns[0].listOfElevators[0].currentFloor = -4
	bat.listOfColumns[0].listOfElevators[0].state = "idle"
	bat.listOfColumns[0].listOfElevators[1].currentFloor = 1
	bat.listOfColumns[0].listOfElevators[1].state = "idle"
	bat.listOfColumns[0].listOfElevators[2].currentFloor = -3
	bat.listOfColumns[0].listOfElevators[2].direction = "down"
	bat.listOfColumns[0].listOfElevators[3].currentFloor = -6
	bat.listOfColumns[0].listOfElevators[3].direction = "up"
	bat.listOfColumns[0].listOfElevators[4].currentFloor = -1
	bat.listOfColumns[0].listOfElevators[4].direction = "down"
	bat.listOfColumns[0].requestElevator(-3, "up")
	
*/
//************************************************************************
	
}

















	