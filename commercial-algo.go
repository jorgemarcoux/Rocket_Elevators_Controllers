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
   
   for i :=1; i == elevatorsAmount; i++ {
	  columnElevator := Elevator{"A1","idle","none",1,0}
	  c.listOfElevators = append(c.listOfElevators,columnElevator)
   }
}

func (c Column) getSmallerFloorsGap(){
	smallestGap := c.listOfElevators[0].floorsGap
	for i :=0; i <= len(c.listOfElevators); i++{
		if i < smallestGap{
		  smallestGap = i
		  fmt.Println(smallestGap)
		}
	  }
}

func (c Column) getSmallerFloorsGap2(){
	smallestGap := c.listOfElevators[0].floorsGap
	for _, element := range c.listOfElevators{
		if element < smallestGap{
			smallestGap = element
		}
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
	columnA := Column{"A",7,-6,1,5,[]Elevator{}}
	columnB := Column{"B",20,1,20,5,[]Elevator{}}
	columnC := Column{"C",21,1,40,5,[]Elevator{}}
	columnD := Column{"D",21,1,60,5,[]Elevator{}}
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
	
	//elevatorA1.requestfloor(9) test requestFloor
	columnA.addElevatorsToColumn(5)
	elevatorA1.floorsGap = 5
	elevatorA2.floorsGap = 2
	columnA.listOfElevators = append(columnA.listOfElevators,elevatorA1)
	columnA.listOfElevators = append(columnA.listOfElevators,elevatorA2)
	fmt.Println(columnA.listOfElevators[0].floorsGap)
	fmt.Println(columnA.listOfElevators[1].floorsGap)
	fmt.Println(columnA.listOfElevators)

	columnA.getSmallerFloorsGap()
}

