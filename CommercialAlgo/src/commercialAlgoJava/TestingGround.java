package commercialAlgoJava;

public class TestingGround {

	public static void main(String[] args) {
		
		//Creating battery and columns
		Battery bat = new Battery(1,4,new Column[4]);
		Column col1 = new Column(1,7,-6,1,new Elevator[5]);
        Column col2 = new Column(2,20,1,20,new Elevator[5]);
        Column col3 = new Column(3,21,1,40,new Elevator[5]);
        Column col4 = new Column(4,21,1,60,new Elevator[5]);
        //Creating 5 elevators for each column
        col1.createElevsInCol(5);
        col2.createElevsInCol(5);
        col3.createElevsInCol(5);
        col4.createElevsInCol(5);
        
		//******Testing area*************************************
        //Scenario1
        /*
        System.out.println("Scenario 1 - call from floor # 1 should send elevator 5");
        col2.listOfElevators[0].currentFloor = 20;
        col2.listOfElevators[0].direction = "down";
        col2.listOfElevators[1].currentFloor = 3;
        col2.listOfElevators[1].direction = "up";
        col2.listOfElevators[1].nextFloor = 15;
        col2.listOfElevators[2].currentFloor = 13;
        col2.listOfElevators[2].direction = "down";
        col2.listOfElevators[2].nextFloor = 1;
        col2.listOfElevators[3].currentFloor = 15;
        col2.listOfElevators[3].direction = "down";
        col2.listOfElevators[4].currentFloor = 6;
        col2.listOfElevators[4].direction = "down";
        col2.listOfElevators[4].nextFloor = 1;
        col2.requestElevator(1, "up");
        */
        //*******************
        
        //Scenario 2
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
        //*************************
        
        //Scenario 3
        /*
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
        */
        
        //*************************
        
        //Scenario 4
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
        //**************************
        
        //Tests requestFloor method
        //col1.listOfElevators[2].currentFloor = 1;
        //col1.listOfElevators[2].requestFloor(5); 

        //col1.listOfElevators[2].currentFloor = 10;
        //col1.listOfElevators[2].requestFloor(5); 
       
	}

}
