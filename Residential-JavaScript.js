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

RequestElevator3(requestedFloor, direction,targetFloor){
  if(elevator1.state == 'idle' && elevator2.state == 'moving'){
     console.log('Elevator',elevator1.id, ' is sent');
    }else if(elevator2.state == 'idle' && elevator1.state == 'moving'){
     console.log('Elevator',elevator2.id, ' is sent');
}

};

   
 RequestElevator(requestedFloor, direction,targetFloor) {
   elevator1.getFloorsGap(requestedFloor);
   elevator2.getFloorsGap(requestedFloor);
 	 if(elevator1.state === 'idle' && elevator2.state === 'idle'){
 	 	if(elevator1.floorsGap < elevator2.floorsGap){
      console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 1');
 	 		console.log('Elevator:' + elevator1.id+ ' at floor '+elevator1.currentFloor+' is sent');
      elevator1.moveElevator(requestedFloor,targetFloor);
	    }else{
       console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 2');
	     console.log('Elevator:' + elevator2.id+ ' at floor '+elevator2.currentFloor+' is sent');
       elevator2.moveElevator(requestedFloor,targetFloor);
	    }
	        }
  	 if(elevator1.currentFloor === requestedFloor && elevator1.direction === direction){
  	 	if(elevator1.direction == 'up' && direction == 'up' && elevator1.currentFloor <= requestedFloor){
          console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 1');
  	 	    console.log('elevator ' + elevator1.id+ ' at floor '+elevator1.currentFloor+'is sent');
          elevator1.moveElevator(requestedFloor,targetFloor); 
  	    }else if(elevator1.direction == 'down' && direction =='down' && elevator1.currentFloor >= requestedFloor){
          console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 1');
  	    	console.log('elevator ' + elevator1.id+ ' at floor '+elevator1.currentFloor+' is sent');
          elevator1.moveElevator(requestedFloor,targetFloor);
  	    }
  	 }else if(elevator2.currentFloor === requestedFloor && elevator2.direction === direction){
  	 	if(elevator2.direction ==='up' && direction ==='up' && elevator2.currentFloor <= requestedFloor){
        console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 2');
  	 	  console.log('elevator ' + elevator2.id+ ' at floor '+elevator2.currentFloor+' is sent');
        elevator2.moveElevator(requestedFloor,targetFloor);
  	 	}else if(elevator2.direction === 'down' && direction ==='down' && elevator1.currentFloor >= requestedFloor){
        console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 2');
  	 	  console.log('elevator ' + elevator2.id+ ' at floor '+elevator2.currentFloor+' is sent');
        elevator2.moveElevator(requestedFloor,targetFloor);
  	 	}
  	 }else if(elevator1.floorsGap < elevator2.floorsGap){
  	 	if(elevator1.direction == direction){
  	 		if(elevator1.direction == 'up' && direction == 'up' && elevator1.currentFloor <= requestedFloor){
          console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 1');
  	 			console.log('elevator ' + elevator1.id+ ' at floor '+elevator1.currentFloor+' is sent');
          elevator1.moveElevator(requestedFloor,targetFloor);
  	 		}else if(elevator1.direction == 'down' && direction =='down' && elevator1.currentFloor >= requestedFloor){
          console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 1');
  	 			console.log('elevator ' + elevator1.id+ ' at floor '+elevator1.currentFloor+' is sent');
          elevator1.moveElevator(requestedFloor,targetFloor);
  	 	}else if(elevator2.direction == direction){
  	 		if(elevator1.direction == 'up' && direction == 'up' && elevator1.currentFloor <= requestedFloor){
          console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 2');
  	 			console.log('elevator ' + elevator2.id+ ' at floor '+elevator2.currentFloor+' is sent');
          elevator2.moveElevator(requestedFloor,targetFloor);
  	 		}else if(elevator1.direction == 'down' && direction =='down' && elevator1.currentFloor >= requestedFloor){
          console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 2');
  	 			console.log('elevator ' + elevator2.id+ ' at floor '+elevator2.currentFloor+' is sent');
          elevator2.moveElevator(requestedFloor,targetFloor);
  	 		}
  	 	}
  	 	}
  	 }
    if(elevator1.state == 'idle' && elevator2.state == 'moving'){
     console.log('Scenario - call from floor #',requestedFloor, 'going to #',targetFloor, 'should send elevator 1');
     console.log('elevator ' + elevator1.id+ ' at floor '+elevator1.currentFloor+' is sent');
     elevator1.moveElevator(requestedFloor,targetFloor);
    } 
    else if(elevator2.state == 'idle' && elevator1.state == 'moving'){
     console.log('Elevator',elevator2.id, ' is sent');
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

  
  RequestFloor(requestedFloor) {
    this.internalList.push(requestedFloor);
    if(requestedFloor>this.currentFloor){
    this.direction = 'up';
    console.log('Elevator ',this.id, 'direction:',this.direction);
    console.log('Moving to floor #',requestedFloor);
    while(this.currentFloor<requestedFloor){
      this.currentFloor++;
      sleep(2000);
      console.log('Elevator ',this.id, 'is at floor #', this.currentFloor);
    }
    }else if(requestedFloor<this.currentFloor){
    this.direction = 'down';
    console.log('Elvator direction:',this.direction);
    console.log('Moving to floor #',requestedFloor);
    while(requestedFloor<this.currentFloor){
      this.currentFloor--;
      sleep(200);
      console.log('elevator ',this.id, 'is at floor #', this.currentFloor);
      }
      
    }
    if(this.currentFloor === requestedFloor){
    console.log('Elvator stopped');
    sleep(2000);
    console.log('Elevator',this.id, 'arrived at target floor');
    sleep(2000);
    console.log('Opening doors...');
    sleep(2000);
    console.log('Person exits the elevator');
    sleep(2000);
    console.log('closing doors...');
    sleep(2000);
  }
  
  };



 getFloorsGap(requestedFloor){
 	this.floorsGap = Math.abs(this.currentFloor-requestedFloor);
 };

 getTargetFloorGap(internalList,currentFloor){
 	targetFloorGap = Math.abs(internalList-requestedFloor);
 };

 moveElevator(requestedFloor,targetFloor){
  if(requestedFloor>this.currentFloor){
    this.direction = 'up';
    console.log('Elevator direction:',this.direction);
    while(this.currentFloor<requestedFloor){
      this.currentFloor++;
      sleep(2000);
      console.log('Elevator ',this.id, 'is at floor #', this.currentFloor);
    }
  }else if(requestedFloor<this.currentFloor){
    this.direction = 'down';
    console.log('Elvator direction:',this.direction);
    while(requestedFloor<this.currentFloor){
      this.currentFloor--;
      sleep(200);
      console.log('Elevator ',this.id, 'is at floor #', this.currentFloor);
      }
      
    }
  
   if(this.currentFloor === requestedFloor){
    console.log('Elvator stopped');
    sleep(2000);
    console.log('Elevator',this.id, 'arrived at target floor');
    sleep(2000);
    console.log('Opening doors...');
    sleep(2000);
    console.log('Person enters the elevator');
    sleep(2000);
    console.log('closing doors...');
    sleep(2000);
    console.log('going to requested floor');
    sleep(2000);
   }
   if(targetFloor>this.currentFloor){
    this.direction = 'up';
    console.log('Elevator direction: '+ this.direction);
   }
   while(targetFloor>this.currentFloor){
    this.currentFloor++;
    sleep(200);
    console.log('Elevator ',this.id, 'is at floor #', this.currentFloor);
  }
   while(targetFloor<this.currentFloor){
    this.currentFloor--;
    sleep(2000);
    console.log('Elevator ',this.id, 'is at floor #', this.currentFloor);
 }
   if(this.currentFloor === targetFloor){
    sleep(2000);
    console.log('Elvator stopped');
    sleep(2000);
    console.log('Elevator ',this.id, 'arrived at target floor');
    sleep(2000);
    console.log('Opening doors...');
    sleep(2000);
    console.log('Person exits the elevator');
    sleep(2000);
    console.log('closing doors...');
   }

 }

};//close Elevator class

/////////////////////////////////
function sleep(milliseconds){
  var start = new Date().getTime();
  while ((new Date().getTime() - start)< milliseconds);
};

//**********************Test Section*****************************
var column1 = new Column(1,10,1,10,2,[2]);
var elevator1 = new Elevator(1);
var elevator2 = new Elevator(2);


function scenario1() {
  elevator1.state = 'idle';
  elevator1.currentFloor = 2;
  elevator2.state = 'idle';
  elevator2.currentFloor = 6;

  column1.RequestElevator(3,'up',7);

};

function scenario2() {
  
  elevator1.state = 'idle';
  elevator1.currentFloor = 10;
  elevator2.state = 'idle';
  elevator2.currentFloor = 3;
  sleep(1000);
  column1.RequestElevator(1,'up',6);//scenario 2 - 1st call

  //scenario 2 - 2nd call
  sleep(2000);
  elevator2.currentFloor = 6;
  column1.RequestElevator(3,'up',5);

  //scenario 2 - 3rd call
  sleep(2000);
  elevator2.currentFloor = 5;
  column1.RequestElevator(9,'down',2);
  
};

function scenario3(){
  elevator1.state = 'idle';
  elevator1.currentFloor = 10;
  elevator2.state = 'moving';
  elevator2.currentFloor = 3;
  column1.RequestElevator(3,'down',2);
  //2nd call
  elevator1.currentFloor = 2;
  elevator2.currentFloor = 6;
  elevator2.state = 'idle';
  column1.RequestElevator(10,'down',3);
};



//scenario1();
//scenario2();
//scenario3()


//To make your own test, just call the function: column1.RequestElevator(requestedFloor, direction,targetFloor);

/*********Internal Button Test*********/
elevator1.currentFloor = 1;
elevator1.RequestFloor(10);
//elevator2.currentFloor = 8;
//elevator2.RequestFloor(elevator1,5);

/**********************End Test Section********************/


//elevator1.internalList.push(1,5);

/* test

while(this.internalList>0){
  this.internalList.forEach(RequestFloor);
}

*/










