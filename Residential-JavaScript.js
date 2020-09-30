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


   
 RequestElevator(RequestedFloor, Direction) {
 	 if(elevator1.status === 'idle' && elevator2.status === 'idle'){
 	 	if(elevator1.floorsGap < elevator2.floorsGap){
 	 		console.log('elevator:' + elevator1.id+' is sent');
	    }else{
	     console.log('elevator:' + elevator2.id+' is sent');
	    }
	        }
  	 if(elevator1.currentFloor === RequestedFloor && elevator1.direction === Direction){
  	 	if(elevator1.direction == 'up' && Direction == 'up' && elevator1.currentFloor <= RequestedFloor){
  	 	    console.log('elevator ' + elevator1.id+'is sent'); 
  	    }else if(elevator1.direction == 'down' && Direction =='down' && elevator1.currentFloor >= RequestedFloor){
  	    	console.log('elevator ' + elevator1.id+' is sent');
  	    }
  	 }else if(elevator2.currentFloor === RequestedFloor && elevator2.direction === Direction){
  	 	if(elevator2.direction ==='up' && Direction ==='up' && elevator2.currentFloor <= RequestedFloor){
  	 	  console.log('elevator ' + elevator2.id+' is sent');
  	 	}else if(elevator2.direction === 'down' && Direction ==='down' && elevator1.currentFloor >= RequestedFloor){
  	 	  console.log('elevator ' + elevator2.id+' is sent');
  	 	}
  	 }else if(elevator1.floorsGap < elevator2.floorsGap){
  	 	if(elevator1.direction == Direction){
  	 		if(elevator1.direction == 'up' && Direction == 'up' && elevator1.currentFloor <= RequestedFloor){
  	 			console.log('elevator ' + elevator1.id+' is sent');
  	 		}else if(elevator1.direction == 'down' && Direction =='down' && elevator1.currentFloor >= RequestedFloor){
  	 			console.log('elevator ' + elevator1.id+' is sent');
  	 	}else if(elevator2.direction == Direction){
  	 		if(elevator1.direction == 'up' && Direction == 'up' && elevator1.currentFloor <= RequestedFloor){
  	 			console.log('elevator ' + elevator2.id+' is sent');
  	 		}else if(elevator1.direction == 'down' && Direction =='down' && elevator1.currentFloor >= RequestedFloor){
  	 			console.log('elevator ' + elevator2.id+' is sent');
  	 		}
  	 	}
  	 	}
  	 }
  
};



};//closes Column class

class Elevator {
	constructor(id){
       this.id = id;
       this.state = 'idle'; //idle, moving, stopped, offline
       this.direction = 'none';//none, up, down
       this.currentFloor = 1;
       this.floorsGap;
       this.targetFloorGap;
       this.directionDownCheck;

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

  
  RequestFloor(Elevator,RequestedFloor) {
    this.internalList.push(RequestedFloor);
    this.currentFloor == RequestedFloor;
  };


 getFloorsGap(RequestedFloor){
 	this.floorsGap = Math.abs(this.currentFloor-RequestedFloor);
 };

 getTargetFloorGap(internalList,currentFloor){
 	targetFloorGap = Math.abs(internalList-RequestedFloor);
 };

 moveElevator(Elevator,RequestedFloor){
  var m = this.currentFloor;
  if(RequestedFloor>this.currentFloor){
    while(this.currentFloor<RequestedFloor){
      m++;
      console.log(this.elevator, 'is at floor #', m);
    }
  }else if(RequestedFloor<this.currentFloor){
    while(RequestedFloor<this.currentFloor){
      m--;
      console.log(this.elevator, 'is at floor #', m);
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

/*
function scenario1() {
  elevator1.status = 'idle';
  elevator1.currentFloor = 2;
  elevator2.status = 'idle';
  elevator2.currentFloor = 6;

  elevator1.getFloorsGap(3);
  elevator2.getFloorsGap(3);

  column1.RequestElevator(3,'up');
  console.log('scenario 1 should send elevator 1');

};
*/
function scenario2() {
  elevator1.status = 'idle';
  elevator1.currentFloor = 10;
  elevator2.status = 'idle';
  elevator2.currentFloor = 3;
  elevator1.getFloorsGap(1);
  elevator2.getFloorsGap(1);

  column1.RequestElevator(1,'up');//scenario 2 1st call
  console.log('scenario 2 - 1st call sends elevator 2');

  //scenario 2 - 2nd call
  elevator2.currentFloor = 6;
  elevator1.getFloorsGap(3);
  elevator2.getFloorsGap(3);

  column1.RequestElevator(3,'up');
  console.log('scenario 2 - 2nd call sends elevator 2');

  //scenario 2 - 3rd call
  elevator2.currentFloor = 5;
  elevator1.getFloorsGap(9);
  elevator2.getFloorsGap(9);

  column1.RequestElevator(9,'down');
  console.log('scenario 2 - 3rd call send elevator 1');
};

//scenario1();
//scenario2();

/*********Internal Button Test*********/
//elevator1.RequestFloor(elevator1,5);






/**********************End Test Section********************/













