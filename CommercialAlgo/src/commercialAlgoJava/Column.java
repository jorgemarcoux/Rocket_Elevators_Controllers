package commercialAlgoJava;
import java.util.*;

public class Column {

	public int id;
    public int floorsNumber;
    public int minimumFloor;
    public int maximumFloor;
    public int elevatorsInColumn;
    public Elevator[] listOfElevators = new Elevator[5];
    public Elevator chosenElevator;
    public int smallestGap = 100;
    public List<Elevator> elevInSameFloor = new ArrayList<Elevator>(); //Contains elevators already in same floor as floorNumber
    public List<Elevator> elevatorsGoingUp = new ArrayList<Elevator>();//Contains elevators going up 
    public List<Elevator> elevatorsGoingDown = new ArrayList<Elevator>();//Contains elevators going down
    public List<Elevator> elevToNextFloor = new ArrayList<Elevator>();//Contains elevators which next stop is floorNumber
    
	
    
	//Column constructor
	public Column(int _id, int _floorsNumber, int _minimumFloor, int _maximumFloor, Elevator[] listOfElevators) {
		this.id = _id;
        this.floorsNumber = _floorsNumber;
        this.minimumFloor = _minimumFloor;
        this.maximumFloor = _maximumFloor;
        this.listOfElevators = listOfElevators;
	}
	
	//Method to create elevators in column
    public void createElevsInCol(int _numberOfElevators){
       for (int i = 0; i < _numberOfElevators; i++ ){
          listOfElevators[i] = new Elevator(i+1,"idle","none",1,0,0);
       };
       
    }
     
	
  //Method to get the elevator with smallest floorGap (requestedFloor-elevator.currentFloor)
  //It takes a list of elevator as a parameter to loop it and find the smallestGap.
    public void getSmallerFloorsGap(List<Elevator>ArrayList, int floorNumber){
       for(int i=0; i < ArrayList.size(); i++){
    	   ArrayList.get(i).floorsGap = Math.abs(floorNumber -ArrayList.get(i).currentFloor);
           if(ArrayList.get(i).floorsGap < smallestGap){
             smallestGap = ArrayList.get(i).floorsGap;
             this.chosenElevator = ArrayList.get(i);//chosenElevator is the one witht the smallestGap
           }
       } 
    }
    
  //RequestElevator method
    public void requestElevator(int floorNumber, String direction) {
    	for( int i =0;i < this.listOfElevators.length;i++){
    		//Cheking if there's an elevator already at floorNumber going to same direction
            if(this.listOfElevators[i].currentFloor == floorNumber && this.listOfElevators[i].direction == direction){
               this.elevInSameFloor.add(this.listOfElevators[i]);
               System.out.println("Scenario - call from floor # "+floorNumber+ " going "+ direction+ " should send elevator "+ this.elevInSameFloor.get(0).id);
               System.out.println("Elevator # "+ this.elevInSameFloor.get(0).id + " is already at floor #"+ floorNumber);
               this.elevInSameFloor.get(0).moveElevator(floorNumber);
               
             //Checking elevators going in same direction as requested direction. Putting those elevators
             //inside the elevatorsGoingUp OR the elevatorsGoingDown array.	
            }else if (this.listOfElevators[i].direction == "up" && direction == "up" && this.listOfElevators[i].currentFloor <= floorNumber){
               this.elevatorsGoingUp.add(this.listOfElevators[i]);
            }else if(this.listOfElevators[i].direction == "down" && direction == "down" && this.listOfElevators[i].currentFloor >= floorNumber){
               this.elevatorsGoingDown.add(this.listOfElevators[i]);
               //Checking if elevator's next stop == floorNumber and choose smallest gap and putting those elevators
               //inside the elevToNextFloor array.
            }else if(this.listOfElevators[i].nextFloor == floorNumber){
               this.elevToNextFloor.add(this.listOfElevators[i]);
            }
          //Waiting till end of loop to call getSmallerFloorsGap and move chosen elevator. We call the 
          //getSmallerFloorsGap method and pass the arrays as parameters to 
          //get the 'chosenElevator' (the one with the smallest floorsGap)
            if(i == 4 && this.elevatorsGoingUp.size() > 0){
               this.getSmallerFloorsGap(this.elevatorsGoingUp,floorNumber);
               System.out.println("Scenario - call from floor # "+ floorNumber+ " going "+ direction);
               System.out.println("Elevator "+this.chosenElevator.id+ " is sent");
               this.chosenElevator.moveElevator(floorNumber); 

            }else if (i == 4 && this.elevatorsGoingDown.size() > 0){
               this.getSmallerFloorsGap(this.elevatorsGoingDown,floorNumber);
               System.out.println("Scenario - call from floor # "+ floorNumber+ " going "+ direction);
               System.out.println("Elevator "+this.chosenElevator.id+ " is sent");
               this.chosenElevator.moveElevator(floorNumber);
               

            }else if(i == 4 && this.elevToNextFloor.size() > 0){
               this.getSmallerFloorsGap(this.elevToNextFloor,floorNumber);
               this.chosenElevator.moveElevator(floorNumber);
            }

          }//End for
    }//Closing requestElevator
	
    
	

}//Closing Column class
