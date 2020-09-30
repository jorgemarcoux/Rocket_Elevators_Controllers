switch (Elevator) {
	case 1:
	 if(elevator1.status == 'idle' && elevator2.status == 'idle') {
	   if(elevator1.floorsGap < elevator2.floorsGap){
		   console.log('elevator:' + elevator1.id);
	   }else{
		console.log('elevator:' + elevator2.id);
	   }
	   }
	break;
}//end switch



//case 3


//Code Nic pour avoir floorsGap
elevator2.getFloorsGap(3);
console.log('floors gap: ',elevator2.floorsGap);


//**********Algo working for 1st and 2nd scenario*********
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
  	 }else if(elevator1.floorsGap < elevator2.floorsGap){
  	 	
	    console.log('elevator ' + elevator1.id);
     }else if(elevator1.floorsGap > elevator2.floorsGap && elevator1.status !== 'idle' || elevator2.status == 'idle'){
	   console.log('elevator ' + elevator2.id);
     }
  
};
//**********************************