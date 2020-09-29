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
 	 if(elevator1.status === 'idle' && elevator2.status === 'idle'){
 	 	if(elevator1.floorsGap < elevator2.floorsGap){
 	 		console.log('elevator:' + elevator1.id);
	    }else{
	    	console.log('elevator:' + elevator2.id);
	    }
	        }
  	 if(elevator1.currentFloor === RequestedFloor && elevator1.direction === Direction){
  	 	if(elevator1.direction == 'up' && Direction == 'up' && elevator1.currentFloor <= RequestedFloor){
  	 	    console.log('elevator ' + elevator1.id); 
  	    }else if(elevator1.direction == 'down' && Direction =='down' && elevator1.currentFloor >= RequestedFloor){
  	    	console.log('elevator ' + elevator1.id);
  	    }
  	 }else if(elevator2.currentFloor === RequestedFloor && elevator2.direction === Direction){
  	 	if(elevator2.direction ==='up' && Direction ==='up' && elevator2.currentFloor <= RequestedFloor){
  	 	  console.log('elevator ' + elevator2.id);
  	 	}else if(elevator2.direction === 'down' && Direction ==='down' && elevator1.currentFloor >= RequestedFloor){
  	 	  console.log('elevator ' + elevator2.id);
  	 	}
  	 }else if(elevator1.floorsGap < elevator2.floorsGap){
  	 	if(elevator1.direction == Direction){
  	 		if(elevator1.direction == 'up' && Direction == 'up' && elevator1.currentFloor <= RequestedFloor){
  	 			console.log('elevator ' + elevator1.id);
  	 		}else if(elevator1.direction == 'down' && Direction =='down' && elevator1.currentFloor >= RequestedFloor){
  	 			console.log('elevator ' + elevator1.id);
  	 	}else if(elevator2.direction == Direction){
  	 		if(elevator1.direction == 'up' && Direction == 'up' && elevator1.currentFloor <= RequestedFloor){
  	 			console.log('elevator ' + elevator2.id);
  	 		}else if(elevator1.direction == 'down' && Direction =='down' && elevator1.currentFloor >= RequestedFloor){
  	 			console.log('elevator ' + elevator2.id);
  	 		}
  	 	}
  	 	}
  	 }
  };


  



};//closes Column class

class Elevator {
	constructor(id,floorsGap,targetFloorGap,directionUpCheck,directionDownCheck){
       this.id = id;
       this.state = 'idle'; //idle, moving, stopped, offline
       this.direction = 'none';//none, up, down
       this.currentFloor =1;
       this.floorsGap = floorsGap || 0;
       this.targetFloorGap = targetFloorGap || 0;
       this.directionDownCheck = directionDownCheck;

	}


    directionUpCheck;
    maximumWeight = 10000;
    currentWeight;
    idleTime = 0;
    doorsOpen = false;
    requestFloorButtons = [1,2,3,4,5,6,7,8,9,10];
    internalList = [];
    doorClosingPathClear = true;
    overload = false;

  
  RequestFloor(Elevator, RequestedFloor) {
    //must move the elevators in its treatment.
  };


 

 getFloorsGap(RequestedFloor){
 	this.floorsGap = Math.abs(this.currentFloor-RequestedFloor);
 };

 getTargetFloorGap(InternalList,currentFloor){
 	targetFloorGap = Math.abs(InternalList-requestedFloorsQueue[i]);
 };

/*
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
*/
/*
 getDirectionUpCheck(Direction,RequestedFloor){
 	if(this.direction =='up' && Direction == 'up'){
 		return
     	if(RequestedFloor < this.currentFloor){
     		return directionUpCheck = false;
        }else{
        	return directionUpCheck = true;
        }
        } 
     };

*/
 

  



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
//EVERYTHING WORKS IN SCENARIO 1
/*
elevator1.status = 'idle';
elevator1.currentFloor = 2;
elevator2.status = 'idle';
elevator2.currentFloor = 6;

column1.requestedFloorsQueue.push({'floor':3,'direction':'up'});


elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[0].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[0].floor);
column1.RequestElevator2(3,'up');
console.log('scenario 1 should send elevator 1');
*/
//******End scenario 1*******

//*****Scenario 2**************
//EVERYTHING WORKS IN SCENARIO 2
/*
elevator1.status = 'idle';
elevator1.currentFloor = 10;
elevator2.status = 'idle';
elevator2.currentFloor = 3;

column1.requestedFloorsQueue.push({'floor':1,'direction':'up'});

elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[0].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[0].floor);
column1.RequestElevator2(1,'up');//scenario 1st call
console.log('scenario 2 - 1st call send elevator 2');

//scenario 2 - 2nd call
elevator2.currentFloor = 6;

column1.requestedFloorsQueue.push({'floor':3,'direction':'up'});

elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[1].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[1].floor);


column1.RequestElevator2(3,'up');
console.log('scenario 2 - 2nd call send elevator 2');

//scenario 2 - 3rd call
elevator2.currentFloor = 5;

column1.requestedFloorsQueue.push({'floor':9,'direction':'down'});

elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[2].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[2].floor);

column1.RequestElevator2(9,'down');
console.log('scenario 2 - 3rd call send elevator 1');

*/
//******End scenario 2*******


//*****Scenario 3**************

elevator1.status = 'idle';
elevator1.currentFloor = 10;
elevator2.status = 'moving';
elevator2.direction = 'up';
elevator2.currentFloor = 3;
//elevator2.internalList.push(6);

column1.requestedFloorsQueue.push({'floor':3,'direction':'down'});

elevator1.floorsGap = Math.abs(elevator1.currentFloor-column1.requestedFloorsQueue[0].floor);
elevator2.floorsGap = Math.abs(elevator2.currentFloor-column1.requestedFloorsQueue[0].floor);
column1.RequestElevator2(3,'down');
//console.log("scenario 3 - 1st call schould send elevator 1")








/**********************End Test Section********************/















