package commercialAlgoJava;

public class Battery {
	int id;
    int numberOfColumns;
    Column[] listOfColumns;
    
    //Battery constructor
    public Battery(int id, int numberOfColumns,Column[] listOfColumns) {
    	this.id = id;
        this.numberOfColumns = numberOfColumns;
        this.listOfColumns  = listOfColumns;
    }
    
    
	
}
