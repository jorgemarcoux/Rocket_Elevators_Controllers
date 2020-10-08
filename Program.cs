using System;

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
      public createElevsInCol(int _numberOfElevators)
      {
         for (int i = 0; i < _numberOfElevators; i++ );
         this.listOfElevators.Add() 
      }


    
           

    }//Closing column class


    public class Elevator
    {
      public int id;
      public string state;
      public string direction;
      public string currentFloor;
      public int floorsGap;
      public int nextFloor;

    
      //Elevator constructor
      public Elevator(int _id, string _state, string _direction, string _currentFloor, int _floorsGap, int _nextFloor)
      {
        this.id = _id;
        this.state = _state;
        this.direction = _direction;
        this.currentFloor = _currentFloor;
        this.floorsGap = _floorsGap;
        this.nextFloor = _nextFloor; 
      }



    }//Closing Elevator class







   









    class Program
    {
        static void Main(string[] args)
        {
            //Creating battery, columns and elevators
            Battery bat =  new Battery(1,4, new Column[5]);
            Column col1 = new Column(1,7,-6,1,new Elevator[5]);
            Column col2 = new Column(2,20,1,20,new Elevator[5]);
            Column col3 = new Column(3,21,1,40,new Elevator[5]);
            Column col4 = new Column(4,21,1,60,new Elevator[5]);

             

  //Elevator(int _id, string _state, string _direction, string _currentFloor, int _floorsGap, int _nextFloor)
   
            Console.WriteLine(col1.maximumFloor);
        }
    }
}//Closing  namespace

    
