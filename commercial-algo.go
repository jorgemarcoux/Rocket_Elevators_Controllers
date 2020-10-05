package main 

import (
	"fmt"
	"time"
)

type Battery struct{
	id int
	numberOfColumns int
}



type Column struct{
	id string
	floorsNumber int
	minimumFloor int
	maximumFloor int
	elevatorsInColumn int
	listOfElevators []Elevator 
	//requestedFloorsQueue [] int
	
}

func (c Column) addElevatorsToColumn(elevatorsAmount int){
   var listOfElevators []Elevator
   for i :=0; i < elevatorsAmount; i++{
	  columnElevator := Elevator{"A1","idle","none",1,0}
	  //listOfElevators: append(columnElevator)
	  listOfElevators = append(listOfElevators,columnElevator)
   }
}

func (c Column) requestElevator(floorNumber int, direction string){

}

type Elevator struct{
	id  string
	state string
	direction string
	currentFloor int
	floorsGap int
	
} 

func (e Elevator) requestfloor(requestedFloor int) {
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
	battery1 := Battery{1,4}
	fmt.Printf("Battery1 id is: %v",battery1.id)
	columnA := Column{"A",7,-6,1,5,}
	columnB := Column{"B",20,1,20,5}
	columnC := Column{"C",21,1,40,5}
	columnD := Column{"D",21,1,60,5}
	fmt.Println(columnA.floorsNumber)
	fmt.Println(columnB.floorsNumber)
	fmt.Println(columnC.floorsNumber)
	fmt.Println(columnD.floorsNumber)

	elevatorA1 := Elevator{"A1","idle","none",1,0}
	elevatorA2 := Elevator{"A2","idle","none",1,0}
	elevatorA3 := Elevator{"A3","idle","none",1,0}
	fmt.Println("Elevator 1 in Column A is at: ",elevatorA1.currentFloor)
	fmt.Println("Elevator 2 in Column A id is: ",elevatorA2.id)
	fmt.Println("Elevator 3 in Column A id is: ",elevatorA3.id)
	
	//elevatorA1.requestfloor(9)
	columnA.addElevatorsToColumn(5)
	fmt.Println(columnA.listOfElevators)
}