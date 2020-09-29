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
   //if that works for scenario 3
 RequestElevator2(RequestedFloor, Direction) {
  	 if(elevator1.currentFloor === RequestedFloor && elevator1.direction === Direction){
  	 	console.log('elevator ' + elevator1.id); 
  	 }else if(elevator2.currentFloor === RequestedFloor && elevator2.direction === Direction){
  	 	console.log('elevator ' + elevator2.id);
  	 }
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


 

 getFloorsGap(currentFloor,requestedFloorsQueue){
 	floorsGap = Math.abs(this.currentFloor-column1.requestedFloorsQueue[0].floor);
 };

 getTargetFloorGap(InternalList,currentFloor){
 	targetFloorGap = Math.abs(InternalList-requestedFloorsQueue[i]);
 };


 elevatorDirectionCheck(Direction,RequestedFloor){
     if(this.direction =='up' && Direction == 'up')
     	if(RequestedFloor < this.currentFloor){
     		directionUpCheck = false;
     	}else if(RequestedFloor >= this.currentFloor){
     		directionUpCheck = true;
     	}
     else if(this.direction=='down' && Direction == 'down')
     	if(RequestedFloor > this.currentFloor){
     		directionDownCheck = false;
     	}else if(RequestedFloor <= this.currentFloor){
     		directionDownCheck = true;
     	}
 };

 getDirectionUpCheck(Direction,RequestedFloor){
 	if(this.direction =='up' && Direction == 'up'){
     	if(RequestedFloor < this.currentFloor){
     		directionUpCheck = false;
        }else{
        	directionUpCheck = true;
        }
        } 
     };
 


  



};//close Elevator class






class RequestFloorButtons {
	constructor(){
      status:'off'//pressed;
	};


}//close RequestFloorButtons class


//**********************Test Section*****************************
var column1 = new Column(1,10,1,10,2,[2]);
var elevator1 = new Elevator(1);
var elevator2 = new Elevator(2);

//*********Scenario 1**********
/*
elevator1.status = 'idle';
elevator1.currentFloor = 2;
elevator2.status = 'idle';
elevator2.currentFloor = 6;

column1.requestedFloorsQueue.push({'floor':3,'direction':'up'});


elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[0].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[0].floor);
column1.RequestElevator();
*/
//******End scenario 1*******
//*****Scenario 3**************

elevator1.status = 'idle';
elevator1.currentFloor = 10;
elevator2.status = 'moving';
elevator2.direction = 'up';
elevator2.currentFloor = 3;
elevator2.internalList.push(6);

column1.requestedFloorsQueue.push({'floor':3,'direction':'up'});

elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[0].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[0].floor);
column1.RequestElevator2(3,'up');

elevator2.getDirectionUpCheck(3,'up');








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















