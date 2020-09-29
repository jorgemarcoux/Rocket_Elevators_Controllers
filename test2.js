switch (bestElevator) {
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


