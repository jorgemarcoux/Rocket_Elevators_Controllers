class Column{
  constructor(id,floorsNumber,minimumFloor,maximumFloor,elevatorsPerColumn){
     this.id = id;
     this.floorsNumber = floorsNumber;
     this.minimumFloor = minimumFloor;
     this.maximumFloor = maximumFloor;
     this.elevatorsPerColumn = elevatorsPerColumn;
     this.elevatorList = [];
  }

     floorsList =[1,2,3,4,5,6,7,8,9,10]

  RequestElevator(RequestedFloor, Direction) {
  	 //must return the chosen elevator and move the elevators in its treatment
  };


  


};//closes Column class

class Elevator {
	constructor(id){
       this.id = id;
	}

	state = 'onService';
    direction = 'None'; //none, up, down
    status = 'idle'; //idle, moving, stopped, offline
    currentFloor;
    maximumWeight = 10000;
    currentWeight;
    idleTime = 0;
    doorsOpen = false;
    requestFloorButtons = [1,2,3,4,5,6,7,8,9,10];
    internalList = [];
    floorsGap;
    targetFloorGap;
    doorClosingPathClear = true;
    overload = false;
    directionUpCheck;
    directionDownCheck;

  
  RequestFloor(Elevator, RequestedFloor) {
    //must move the elevators in its treatment.
  };


  getElevatorDirection(){
  	if(internalList > currentFloor){
  		direction = 'up';
  		status = 'moving';
  	}else if(internalList < currentFloor){
  		direction = 'down';
  		status = 'moving';
  	}
  }
  
  


  



}//close Elevator class

class CallElevatorButtons {
	constructor(minimumFloor,maximumFloor){
      status: 'off';
      callFrom;
      dierction;
	}



}//close CallElevatorButtons class

class RequestFloorButtons {
	constructor(){
      status:'off';
	}


}//close RequestFloorButtons class



var column1 = new Column(1,10,1,10,2,[2]);
var elevator1 = new Elevator(1);












