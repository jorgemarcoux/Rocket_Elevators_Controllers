package main 

import (
	"fmt"
	"time"
	
)

var chosenElevator Elevator
var smallestGap = 100

func positive(n int) int { //Converting negative numbers to absolute values
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
//Creating battery of columns.
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
	
}
//Update column values after creating each column with tame values with the createBattery method
func (c *Column) ChangeColumnValues(_floorNumber int, _minimumFloor int,_maximumFloor int){
	c.floorsNumber = _floorNumber
	c.minimumFloor = _minimumFloor
	c.maximumFloor = _maximumFloor

}

func (c *Column) createColumn(numberOfElevators int){
   for i :=0; i < numberOfElevators; i++{
	   c.listOfElevators = append(c.listOfElevators, Elevator{i+1,"idle","none",1,0,0})
   }
}

//Method to get the elevator with smallest floorGap (requestedFloor-elevator.currentFloor)
//It takes a list of elevator as a parameter to loop  throught it and find the smallestGap. 
func (c *Column) getSmallerFloorsGap(array []Elevator, floorNumber int){
	for i :=0; i < len(array); i++{
		array[i].floorsGap = positive(floorNumber - array[i].currentFloor)
		if array[i].floorsGap < smallestGap{
		  smallestGap = array[i].floorsGap
		  chosenElevator = array[i]//chosenElevator is the one witht the smallestGap
		}
		
	  }
}


func (c *Column) requestElevator(floorNumber int, direction string){
	var elevInSameFloor []Elevator//Contains elevators already in same floor as floorNumber
	var elevatorsGoingUp []Elevator//Contains elevators goinf up
	var elevatorsGoingDown []Elevator//Contains elevators going down
	var elevToNextFloor []Elevator//Contains elevators which next stop is floorNumber
	for i :=0; i < len(c.listOfElevators); i++{
		//Chcking if there's an elevator already at floorNumber going to same direction
	   if c.listOfElevators[i].currentFloor == floorNumber && c.listOfElevators[i].direction == direction{
			elevInSameFloor = append(elevInSameFloor,c.listOfElevators[i])
			fmt.Println("Scenario - call from floor #",floorNumber, "going", direction, "should send elevator",elevInSameFloor[0].id)
			fmt.Println("Elevator #",elevInSameFloor[0].id, "is already at floor #", floorNumber)
			elevInSameFloor[0].moveElevator(floorNumber)
			break
		//Checking elevators going in same direction as requested direction. Putting those elevators
        //inside the elevatorsGoingUp OR the elevatorsGoingDown array.
	   }else if c.listOfElevators[i].direction == "up" && direction == "up" && c.listOfElevators[i].currentFloor <= floorNumber{
			elevatorsGoingUp = append(elevatorsGoingUp,c.listOfElevators[i])
			
	   }else if c.listOfElevators[i].direction == "down" && direction == "down" && c.listOfElevators[i].currentFloor >= floorNumber{
			elevatorsGoingDown = append(elevatorsGoingDown,c.listOfElevators[i])
			
	   //Checking if elevator's next stop == floorNumber and choose smallest gap.
	   //Putting those elevators inside the elevToNextFloor array.
	   }else if c.listOfElevators[i].nextFloor == floorNumber{
			elevToNextFloor = append(elevToNextFloor, c.listOfElevators[i])		   
	   }
	  //Waiting till end of loop to call getSmallerFloorsGap and move chosen elevator. We call the 
	 //getSmallerFloorsGap method and pass the arrays as parameters to get 
	 //the 'chosenElevator' (the one with the smallest floorsGap)
	   if i == 4 && len(elevatorsGoingUp) > 0{
		c.getSmallerFloorsGap(elevatorsGoingUp,floorNumber)
		fmt.Println("Scenario - call from floor #",floorNumber, "going", direction)
		fmt.Println("Elevator going up and requested direction up")//DELETE
		fmt.Println("Elevator",chosenElevator.id,"should be sent")
		chosenElevator.moveElevator(floorNumber)

	  }else if i == 4 && len(elevatorsGoingDown) > 0{
		c.getSmallerFloorsGap(elevatorsGoingDown,floorNumber)
		fmt.Println("Scenario - call from floor #",floorNumber, "going", direction)
		fmt.Println("Elevator",chosenElevator.id, "is sent")
		fmt.Println("Elevator going down and requested direction down. Elevator's current floor:",chosenElevator.currentFloor)
		chosenElevator.moveElevator(floorNumber)
		  
	  }else if i == 4 && len(elevToNextFloor) > 0{
		c.getSmallerFloorsGap(elevToNextFloor,floorNumber)
		chosenElevator.moveElevator(floorNumber)
	  }

	}//Closing for iterating c.listOfElevators
	

}//Closing requestElevator

type Elevator struct{
	id  int
	state string
	direction string
	currentFloor int
	floorsGap int
	nextFloor int
	
} 

//Method to move elvator to requested floor
func (e Elevator) moveElevator(requestedFloor int){
	if requestedFloor > e.currentFloor{
		e.direction = "up"
		fmt.Println("Elevator",chosenElevator.id,"at floor",chosenElevator.currentFloor,"direction is", e.direction)
		for e.currentFloor < requestedFloor {
			e.currentFloor += 1 // Incrementing currentfloor since direction is up
			fmt.Println("Elevator", e.id, "is at floor #", e.currentFloor)
			if e.currentFloor == requestedFloor{ //When elevators arrives at the requested floor...
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
		}//Checking for direction down
	}else if requestedFloor < e.currentFloor{
		e.direction = "down"
		fmt.Println("Elevator",chosenElevator.id,"at floor",chosenElevator.currentFloor,"direction is", e.direction)
		for e.currentFloor > requestedFloor{
			e.currentFloor -= 1 // Decreasing currentfloor since direction is down
			fmt.Println("Elevator", e.id, "is at floor #", e.currentFloor)
			if e.currentFloor == requestedFloor{ //When elevators arrives at the requested floor...
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
 }//End moveElevator method


func (e *Elevator) requestfloor(requestedFloor int) {
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
		fmt.Println("Elevator",e.id, "is at floor #", e.currentFloor,"Direction =", e.direction)
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

}//End requestFloor method

func main(){
	bat := &Battery{1,4,[]Column{}}
	bat.createBattery(4,5)
	bat.listOfColumns[0].ChangeColumnValues(7,-6,1)
	bat.listOfColumns[1].ChangeColumnValues(20,1,20)
	bat.listOfColumns[2].ChangeColumnValues(21,1,40)
	bat.listOfColumns[3].ChangeColumnValues(21,1,60)
	



//******************************testing requestElevator***********************************
//****************SCENARIO 1*************************************
 /*
    fmt.Println("Scenario 1 - call from floor # 1 should send elevator 5")
	bat.listOfColumns[1].listOfElevators[0].currentFloor = 20
	bat.listOfColumns[1].listOfElevators[0].direction = "down"
	bat.listOfColumns[1].listOfElevators[1].currentFloor = 3
	bat.listOfColumns[1].listOfElevators[1].direction = "up"
	bat.listOfColumns[1].listOfElevators[1].nextFloor = 15
	bat.listOfColumns[1].listOfElevators[2].currentFloor = 13
	bat.listOfColumns[1].listOfElevators[2].direction = "down"
	bat.listOfColumns[1].listOfElevators[2].nextFloor = 1
	bat.listOfColumns[1].listOfElevators[3].currentFloor = 15
	bat.listOfColumns[1].listOfElevators[3].direction = "down"
	bat.listOfColumns[1].listOfElevators[4].currentFloor = 6
	bat.listOfColumns[1].listOfElevators[4].direction = "down"
	bat.listOfColumns[1].listOfElevators[4].nextFloor = 1
	bat.listOfColumns[1].requestElevator(1, "up")
 */

//*********SCENARIO 2 *******************************************************
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

//********SCENARIO 3 WORKING PROPERLY********************************************
 /*
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
 */
//********SCENARIO 4 ***************************
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
//******Testing requestFloor**********************************************
//bat.listOfColumns[1].listOfElevators[0].currentFloor = 8
//bat.listOfColumns[1].listOfElevators[0].requestfloor(1)
//************************************************************************
}

















	