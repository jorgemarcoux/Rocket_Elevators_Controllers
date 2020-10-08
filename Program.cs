using System;
using System.Threading;
using System.Collections.Generic;

namespace Rocket_Elevators_Controllers
{

 

     public class Battery     
    {
      public int id;
      public int numberOfColumns; 
      public Column[] listOfColumns; 
      

      //Battery constructor
      public Battery(int _id,int _numberOfColumns, Column[] listOfColumns)
      {
        this.id = _id;
        this.numberOfColumns = _numberOfColumns;
        this.listOfColumns  = listOfColumns;
      }
      


    }//Closing Battery class


    public class Column 
    {
      public int id;
      public int floorsNumber;
      public int minimumFloor;
      public int maximumFloor;
      public int elevatorsInColumn;
      public Elevator[] listOfElevators = new Elevator[5];
      public List<Elevator> elevInSameFloor = new List<Elevator>();//Contains elevators already in same floor as floorNumber
      public List<Elevator> elevatorsGoingUp = new List<Elevator>();//Contains elevators goinf up
      public List<Elevator> elevatorsGoingDown = new List<Elevator>();//Contains elevators going down
      public List<Elevator> elevToNextFloor = new List<Elevator>();//Contains elevators which next stop is floorNumber
      

      
     //Column constructor
      public Column(int _id, int _floorsNumber, int _minimumFloor, int _maximumFloor, Elevator[] listOfElevators)
      {
        this.id = _id;
        this.floorsNumber = _floorsNumber;
        this.minimumFloor = _minimumFloor;
        this.maximumFloor = _maximumFloor;
        this.listOfElevators = listOfElevators;
      }
      
      //Method to create elevators in column
      public void createElevsInCol(int _numberOfElevators)
      {
         for (int i = 0; i < _numberOfElevators; i++ ){
            listOfElevators[i] = new Elevator(i+1,"idle","none",1,0,0);
         };
         
      }

      //RequestElevator method
      public void requestElevator(int floorNumber, string direction)
      {
        for( int i =0;i < this.listOfElevators.Length;i++){
          //Chcking if there's an elevator already at floorNumber going to same direction
          if(this.listOfElevators[i].currentFloor == floorNumber && this.listOfElevators[i].direction == direction){
             this.elevInSameFloor.Add(this.listOfElevators[i]);
             Console.WriteLine("Scenario - call from floor # "+floorNumber+ " going "+ direction+ " should send elevator "+ elevInSameFloor[0].id);
             Console.WriteLine("Elevator # "+ this.elevInSameFloor[0].id + " is already at floor #"+ floorNumber);
             this.elevInSameFloor[0].moveElevator(floorNumber);
             
          //Checking elevators going in same direction as requested direction	
          }else if (this.listOfElevators[i].direction == "up" && direction == "up" && this.listOfElevators[i].currentFloor <= floorNumber){
             this.elevatorsGoingUp.Add(this.listOfElevators[i]);
          }else if(this.listOfElevators[i].direction == "down" && direction == "down" && this.listOfElevators[i].currentFloor >= floorNumber){
             this.elevatorsGoingDown.Add(this.listOfElevators[i]);
          //Checking if elevator's next stop == floorNumber and choose smallest gap
          }else if(this.listOfElevators[i].nextFloor == floorNumber){
             this.elevToNextFloor.Add(this.listOfElevators[i]);
          }
          //Waiting till end of loop to call getSmallerFloorsGap and move chosen elevator
          if(i == 4 && this.elevatorsGoingUp.Count > 0){
             this.elevatorsGoingUp.Sort((x, y) => x.floorsGap.CompareTo(y.floorsGap));//sorting elevators to get the one with the smallest floorsGap
             Console.WriteLine("Scenario - call from floor # "+ floorNumber+ " going "+ direction);
             Console.WriteLine("Elevator "+this.elevatorsGoingUp[0].id+ " should be sent");
             this.elevatorsGoingUp[0].moveElevator(floorNumber);

          }else if (i == 4 && elevatorsGoingDown.Count > 0){
             this.elevatorsGoingDown.Sort((x, y) => x.floorsGap.CompareTo(y.floorsGap));
             Console.WriteLine("Scenario - call from floor # "+ floorNumber+ " going "+ direction);
             Console.WriteLine("Elevator "+this.elevatorsGoingDown[0].id+ " should be sent");
             this.elevatorsGoingDown[0].moveElevator(floorNumber);

          }else if(i == 4 && elevToNextFloor.Count > 0){
             this.elevToNextFloor.Sort((x, y) => x.floorsGap.CompareTo(y.floorsGap));
             this.elevToNextFloor[0].moveElevator(floorNumber);
          }

        }//End for
         
      }//Closing requestElevator


    
           

    }//Closing column class


    public class Elevator
    {
      public int id;
      public string state;
      public string direction;
      public int currentFloor;
      public int floorsGap;
      public int nextFloor;

    
      //Elevator constructor
      public Elevator(int _id, string _state, string _direction, int _currentFloor, int _floorsGap, int _nextFloor)
      {
        this.id = _id;
        this.state = _state;
        this.direction = _direction;
        this.currentFloor = _currentFloor;
        this.floorsGap = _floorsGap;
        this.nextFloor = _nextFloor; 
      }

      public void requestFloor(int requestedFloor)
      {
        Console.WriteLine("Requested floor: #"+requestedFloor);
        if (requestedFloor > this.currentFloor) {
           this.direction = "up";
           Console.WriteLine("Elevator "+ this.id+ " is at floor # "+ this.currentFloor+ " Direction = "+ this.direction);
           Thread.Sleep(1000);
           Console.WriteLine("Moving to floor # "+ requestedFloor);
           while (this.currentFloor < requestedFloor){
             this.currentFloor ++;
             Console.WriteLine("Elevator "+ this.id+ " is at floor # "+this.currentFloor);
           }
        } else if (requestedFloor < this.currentFloor){
           this.direction = "down";
           Console.WriteLine("Elevator "+ this.id+ " is at floor # "+ this.currentFloor+ " Direction = "+ this.direction);
           Thread.Sleep(1000);
           Console.WriteLine("Moving to floor # "+ requestedFloor);
           while (this.currentFloor > requestedFloor){
              this.currentFloor --;
              Console.WriteLine("Elevator "+ this.id+ " is at floor # "+ this.currentFloor);
           }
        }
        if (this.currentFloor == requestedFloor){
          Console.WriteLine("Elevator stopped");
          Thread.Sleep(1000);
          Console.WriteLine("Elevator "+ this.id+ " arrived at target floor");
          Thread.Sleep(1000);
          Console.WriteLine("Opening doors...");
          Thread.Sleep(1000);
          Console.WriteLine("Person exits the elevator");
          Thread.Sleep(1000);
          Console.WriteLine("closing doors...");

        }
      }//CLosing requestFloor

      public void moveElevator(int requestedFloor)
      {
      if(requestedFloor > this.currentFloor){
        this.direction = "up";
        Console.WriteLine("Elevator "+ this.id+ " at floor #"+ this.currentFloor+ " direction is " + this.direction);
        while (this.currentFloor < requestedFloor){
          this.currentFloor ++;
          Console.WriteLine("Elevator "+ this.id+ " is at floor #"+ this.currentFloor);
          if (this.currentFloor == requestedFloor){
            Console.WriteLine("Elevator stopped");
            Thread.Sleep(1000);
            Console.WriteLine("Elevator "+ this.id+ " arrived to target floor");
            Thread.Sleep(1000);
            Console.WriteLine("Opening doors");
            Thread.Sleep(1000);
            Console.WriteLine("Person enters the elevator");
            Thread.Sleep(1000);
            Console.WriteLine("Closing doors...");

          }else if (requestedFloor < this.currentFloor){
             this.direction = "down";
             Console.WriteLine("Elevator "+this.id+ " at floor #"+ this.currentFloor+ " direction is "+ this.direction);
             while (this.currentFloor > requestedFloor){
               this.currentFloor--;
               Console.WriteLine("Elevator "+this.id+ " is at floor #"+this.currentFloor);
               if (this.currentFloor == requestedFloor){
                 Console.WriteLine("Elevator stopped");
                 Thread.Sleep(1000);
                 Console.WriteLine("Elevator "+ this.id+ " arrived to target floor");
                 Thread.Sleep(1000);
                 Console.WriteLine("Opening doors");
                 Thread.Sleep(1000);
                 Console.WriteLine("Person enters the elevator");
                 Thread.Sleep(1000);
                 Console.WriteLine("Closing doors...");
               }
             }
          }

        }
      }

      }//Closing moveElevator



    }//Closing Elevator class







   









    class Program
    {
        static void Main(string[] args)
        {
            //Creating battery and columns 
            Battery bat = new Battery(1,4, new Column[5]);
            Column col1 = new Column(1,7,-6,1,new Elevator[5]);
            Column col2 = new Column(2,20,1,20,new Elevator[5]);
            Column col3 = new Column(3,21,1,40,new Elevator[5]);
            Column col4 = new Column(4,21,1,60,new Elevator[5]);
            //Creating 5 elevators inside each column
            col1.createElevsInCol(5);
            col2.createElevsInCol(5);
            col3.createElevsInCol(5);
            col4.createElevsInCol(5);

            //****************Testing Section***************
            //Scenario 1 - NOT WORKING YET
           /* 
            Console.WriteLine("Scenario 1 - call from floor # 1 should send elevator 5");
	          col1.listOfElevators[0].currentFloor = 20;
	          col1.listOfElevators[0].direction = "down";
	          col1.listOfElevators[1].currentFloor = 3;
	          col1.listOfElevators[1].direction = "up";
	          col1.listOfElevators[1].nextFloor = 15;
	          col1.listOfElevators[2].currentFloor = 13;
	          col1.listOfElevators[2].direction = "down";
	          col1.listOfElevators[2].nextFloor = 1;
	          col1.listOfElevators[3].currentFloor = 15;
	          col1.listOfElevators[3].direction = "down";
	          col1.listOfElevators[4].currentFloor = 6;
	          col1.listOfElevators[4].direction = "down";
	          col1.listOfElevators[4].nextFloor = 1;
	          col1.requestElevator(1, "down");
            Console.WriteLine("End");
           */
           //***************************************************
           //scenario 2 - WORKING PROPERLY
           /*
            col3.listOfElevators[0].currentFloor = 1;
            col3.listOfElevators[0].direction = "up";
            col3.listOfElevators[1].currentFloor = 23;
            col3.listOfElevators[1].direction = "up";
            col3.listOfElevators[2].currentFloor =33;
            col3.listOfElevators[2].direction = "down";
            col3.listOfElevators[3].currentFloor = 40;
            col3.listOfElevators[3].direction = "down";
            col3.listOfElevators[4].currentFloor = 39;
            col3.listOfElevators[4].direction = "down";
            col3.requestElevator(1, "up");
           */

           //***************************************************
           //Scenario 3 - not workin properly
           
            col4.listOfElevators[0].currentFloor = 58;
            col4.listOfElevators[0].direction = "down";
            col4.listOfElevators[1].currentFloor = 50;
            col4.listOfElevators[1].direction = "up";
            col4.listOfElevators[2].currentFloor = 46;
            col4.listOfElevators[2].direction = "up";
            col4.listOfElevators[3].currentFloor = 1;
            col4.listOfElevators[3].direction = "up";
            col4.listOfElevators[4].currentFloor = 60;
            col4.listOfElevators[4].direction = "down";
            col4.requestElevator(54, "down");
           
          //***************************************************
          //Scenario 4 - WORKING PROPERLY
          /*
            col1.listOfElevators[0].currentFloor = -4;
	          col1.listOfElevators[0].state = "idle";
	          col1.listOfElevators[1].currentFloor = 1;
	          col1.listOfElevators[1].state = "idle";
	          col1.listOfElevators[2].currentFloor = -3;
	          col1.listOfElevators[2].direction = "down";
	          col1.listOfElevators[3].currentFloor = -6;
	          col1.listOfElevators[3].direction = "up";
	          col1.listOfElevators[4].currentFloor = -1;
	          col1.listOfElevators[4].direction = "down";
	          col1.requestElevator(-3, "up");
           */
          //***************************************************

          //       
            
        }
      
    }
}//Closing  namespace

    
