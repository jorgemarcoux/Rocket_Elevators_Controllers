class Column{
  constructor(id,floorsNumber,minimumFloor,maximumFloor,elevatorsPerColumn){
     this.id = id;
     this.floorsNumber = floorsNumber;
     this.minimumFloor = minimumFloor;
     this.maximumFloor = maximumFloor;
     this.elevatorsPerColumn = elevatorsPerColumn;
     this.elevatorList = [];
  }

     requestedFloorsQueue = [];

     //if that works for 1st scenario
RequestElevator(RequestedFloor, Direction){
  if(elevator1.status == 'idle' && elevator2.status == 'idle') {
	if(elevator1.floorsGap < elevator2.floorsGap){
		console.log('elevator:' + elevator1.id);
	}else{
		console.log('elevator:' + elevator2.id);
	}
	}
};

 RequestElevator2(RequestedFloor, Direction) {
  	 if(internalList[i] || currentFloor == requestedFloorsQueue[i] && elevator1.directon || elevator2.direction == requestedFloorsQueue[i].direction && directionUpCheck || directionDownCheck == true){
  	 	console.log(elevator1.id) || console.log(elevator2.id);
  	 }
  };


  



};//closes Column class

class Elevator {
	constructor(id,column1){
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


 

 getFloorsGap(requestedFloorsQueue,currentFloor){
 	for(var i =0; i<=requestedFloorsQueue.lenght, i++)
 	floorsGap = Math.abs(this.currentFloor-column1.requestedFloorsQueue[i]);
 };

 getTargetFloorGap(InternalList,currentFloor){
 	targetFloorGap = Math.abs(InternalList-requestedFloorsQueue[i]);
 };

/*
 elevatorDirectionQueue(direction,requestedFloorsQueue){
     if(this.direction=='up')
 }
*/
 


  



}//close Elevator class

class CallElevatorButtons {
	constructor(callFrom,direction){
      status = 'off';
      this.callFrom = callFrom;
      this.direction = direction;//up,down
	};



}//close CallElevatorButtons class

class RequestFloorButtons {
	constructor(){
      status:'off';
	};


}//close RequestFloorButtons class


//**********************Test Section*****************************
var column1 = new Column(1,10,1,10,2,[2]);
var elevator1 = new Elevator(1);
var elevator2 = new Elevator(2);

//*********Scenario 1**********
elevator1.status = 'idle';
elevator1.currentFloor = 2;
elevator2.status = 'idle';
elevator2.currentFloor = 6;

column1.requestedFloorsQueue = 3;

elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[i]);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[i]);
column1.RequestElevator(3,'up');
//******End scenario 3***********

//*****Scenario 2**************





/*
    if(InternalList > currentFloor){
    	direction = 'up';
    	status = 'moving';
    }else if(InternalList < currentFloor){
    	direction = 'down';
    	status = 'moving';
    };

*/



//if(internalList[i])











/**********************Closing Test Section********************/















