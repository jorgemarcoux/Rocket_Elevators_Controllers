class Column
    attr_accessor :id, :floorsNumber, :minimumFloor, :maximumFloor, :elevatorsPerColumn 

    def initialize (id,floorsNumber,minimumFloor,maximumFloor,elevatorsPerColumn)
        @id = id
        @floorsNumber = floorsNumber
        @minimumFloor = minimumFloor
        @maximumFloor = maximumFloor
        @elevatorsPerColumn = elevatorsPerColumn
        @elevatorsInColumn =[1,2]
    end

    def requestElevator(requestedFloor,direction,targetFloor)
       elevator1.getFloorsGap(requestedFloor)
       elevator2.getFloorsGap(requestedFloor)
       if elevator1.state == 'idle' and elevator2.state == 'idle'
        if elevator1.floorsGap < elevator2.floorsGap
          puts("Scenario - call from floor # #{requestedFloor} going to ##{targetFloor} should send elevator 1")
          puts("Elevator: #{elevator1.id} at floor ##{elevator1.currentFloor} is sent")
          #elevator1.moveElevator(requestedFloor,targetFloor)
        elsif 
          puts("Scenario - call from floor ##{requestedFloor} going to # #{targetFloor} should send elevator 2")
          puts("Elevator: ##{elevator2.id} at floor #{elevator2.currentFloor} is sent")
          #elevator2.moveElevator(requestedFloor,targetFloor)
        end



       end
    end



end #Closes Column class



class Elevator
    attr_accessor :id, :floorsGap, :state, :direction, :currentFloor, :internalList

    def initialize(id)
        @id = id
        @state = 'idle' # idle, moving, stopped, offline
        @direction = 'none'# none, up, down
        @currentFloor = 1
        @floorsGap = 0
        @internalList = [] 
    end

    def requestFloor(requestedFloor)
        self.internalList.push(requestedFloor)
        if requestedFloor > self.currentFloor
            self.direction = 'up' 
            puts("Elevator #{self.id} at floor #{self.currentFloor} direction: #{self.direction}")
            sleep(1)
            puts("Moving to floor ##{requestedFloor}")
            while self.currentFloor < requestedFloor do
              self.currentFloor+=1
              sleep(1);
              puts("Elevator #{self.id} is at floor # #{self.currentFloor}")
            end
        elsif requestedFloor < self.currentFloor
            self.direction = 'down'
            puts("Elvator #{self.id} is at floor # #{self.currentFloor} direction: #{self.direction}")
            puts("Moving to floor #'#{requestedFloor}")
            while requestedFloor < self.currentFloor
                self.currentFloor-=
                sleep(1)
                puts("Elevator #{self.id} is at floor # #{self.currentFloor}")
            end
        end
        if self.currentFloor == requestedFloor
            puts('Elvator stopped')
            sleep(1)
            puts("Elevator #{self.id} arrived at target floor")
            sleep(1)
            puts('Opening doors...')
            sleep(1);
            puts('Person exits the elevator')
            sleep(1)
            puts('closing doors...')
            sleep(1)
        end
        
    end


end # closes Elevator class




column1.elevatorsInColumn.push(elevator1) 
puts("#{column1.elevatorsInColumn}")
# ---------Test Section --------------
column1 = Column.new(1,10,1,10,2)
elevator1 = Elevator.new(1)
elevator2 = Elevator.new(2)

def scenario1()
    elevator1.state = 'idle'
    elevator1.currentFloor = 2
    elevator2.state = 'idle'
    elevator2.currentFloor = 6
    column1.requestElevator(3,'up',7)
end






# -------------Internal Button Test-------
#elevator1.currentFloor = 1
#elevator1.requestFloor(5)
#elevator2.currentFloor = 8;
#elevator2.requestFloor(5)



