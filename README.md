# Rocket_Elevators_Controllers
Contains the algorithm files for the elevator controllers for the New Rocket Elevator Solutions for both Residential and Commercial Offers

<h2>Residential Alogrithm</h2>
The algorithm for the residential services provided by Rocket Elevators is found in the <i>Residential_Controller.algo</i> file. I also built it on <storng>Javascript and <strong>Python </strong>
applying it to a residential building with 10 stories and 2 elevators.
<br/>
<br>
Please notice that several test scenarios are provided in each file, so you can see the algorithm in action. In order to run these tests, just uncomment the functions <strong>scenario1()</strong>, 
<strong>scenario2()</strong> and <strong>scenario3()</strong> in the "test section" (line 281 to 283 in the <a href src="https://github.com/jorgemarcoux/Rocket_Elevators_Controllers/blob/master/Residential-JavaScript.js">JavaScript file</a> and 221 to 214 in the <a href="https://github.com/jorgemarcoux/Rocket_Elevators_Controllers/blob/master/Residential-Py">Python</a> one)
<br/>
<br>
To see the results of the JS test, you'll need to:
<ol>
<li>Open the <a href src="Resi-JS-Test.html">Resi-JS-Test.html<a> file in your browser</li>
<li>Go to your developer console</li>
<li>Uncomment the scenarios functions</li>
<li>Reaload your browser tab to see the scenarios in action</li>
</ol>
<br>
  <storng>For the Python version, you'll need to install <a href="https://www.python.org/downloads/">Python</a> in your computer and use a command terminal to run the tests</strong>.
<br>  
To make your own test, just call the function: <strong>column1.RequestElevator(requestedFloor, direction,targetFloor)</strong>. Please indicate the direction 'up' or 'down' in lowercase)
  <br>
You can also test the requestFloor method (the one that brings an user to a requestedFloor pressing the internal buttons of the elevator) by uncommenting the test under "Internal Button Test", also provided in the test section.

