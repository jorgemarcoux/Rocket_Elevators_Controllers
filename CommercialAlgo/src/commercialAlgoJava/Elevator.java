package commercialAlgoJava;

public class Elevator {
	public int id;
	public String state;
	public String direction;
    public int currentFloor;
    public int floorsGap;
    public int nextFloor;
	
	//Elevator constructor
    public Elevator(int _id, String _state, String _direction, int _currentFloor, int _floorsGap, int _nextFloor) {
    	this.id = _id;
        this.state = _state;
        this.direction = _direction;
        this.currentFloor = _currentFloor;
        this.floorsGap = _floorsGap;
        this.nextFloor = _nextFloor;
        
    }
    
    
    public void requestFloor(int floorNumber)
    {
    	System.out.println("Requested floor: #"+floorNumber);
      if (floorNumber > this.currentFloor) {
         this.direction = "up";
         System.out.println("Elevator "+ this.id+ " is at floor # "+ this.currentFloor+ " Direction = "+ this.direction);
         System.out.println("Moving to floor # "+ floorNumber);
         while (this.currentFloor < floorNumber){
           this.currentFloor ++;
           System.out.println("Elevator "+ this.id+ " is at floor # "+this.currentFloor);
         }
      } else if (floorNumber < this.currentFloor){
         this.direction = "down";
         System.out.println("Elevator "+ this.id+ " is at floor # "+ this.currentFloor+ " Direction = "+ this.direction);
         System.out.println("Moving to floor # "+ floorNumber);
         while (this.currentFloor > floorNumber){
            this.currentFloor --;
            System.out.println("Elevator "+ this.id+ " is at floor # "+ this.currentFloor);
         }
      }
      if (this.currentFloor == floorNumber){//When elevator has arrived to its destination:
    	  System.out.println("Elevator stopped");
        System.out.println("Elevator "+ this.id+ " arrived at target floor");
        System.out.println("Opening doors...");
        System.out.println("Person exits the elevator");
        System.out.println("closing doors...");

      }
    }//CLosing requestFloor
    
    //Method to move elevator to requested floor
    public void moveElevator(int floorNumber)
    {
    if(floorNumber > this.currentFloor){
      this.direction = "up";
      System.out.println("Elevator "+ this.id+ " at floor #"+ this.currentFloor+ " direction is " + this.direction);
      while (this.currentFloor < floorNumber){
        this.currentFloor ++;// Incrementing currentfloor since direction is up
        System.out.println("Elevator "+ this.id+ " is at floor #"+ this.currentFloor);
        if (this.currentFloor == floorNumber){
        	System.out.println("Elevator stopped");
          System.out.println("Elevator "+ this.id+ " arrived to target floor");
          System.out.println("Opening doors");
          System.out.println("Person enters the elevator");
          System.out.println("Closing doors...");
      
        }
             }
           }else if (floorNumber < this.currentFloor){
           this.direction = "down";//Checking for direction down
           System.out.println("Elevator "+this.id+ " at floor #"+ this.currentFloor+ " direction is "+ this.direction);
           while (this.currentFloor > floorNumber){
             this.currentFloor--;// Decreasing currentfloor since direction is down
             System.out.println("Elevator "+this.id+ " is at floor #"+this.currentFloor);
             if (this.currentFloor == floorNumber){//When elevator has arrived to its destination:
            	 System.out.println("Elevator stopped");
               System.out.println("Elevator "+ this.id+ " arrived to target floor");
               System.out.println("Opening doors");
               System.out.println("Person enters the elevator");
               System.out.println("Closing doors...");
        }

      }
    }

    }//Closing moveElevator
    
    
    
    
    
}//Closing Elevator class
