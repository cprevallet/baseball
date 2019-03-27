//Package main calls trajectory to calculate the trajectory of a baseball with air drag.
//Original author(Fortran90) - Ralph L. Carmichael, Public Domain Aeronautical Software
//Rewritten in Go
package main

import (
	"fmt"
	"github.com/cprevallet/baseball/trajectory"
)

func main() {
	fmt.Println("Baseball trajectory from Public Domain Aeronautical Software. Go version")
	var history []trajectory.TrajectoryPoint
	var dt = 0.1 // time step
	var normalized = true //make initial altitude the reference point in results
	var k = 0
	// Instructions for use of this program: put your values for the three
	//  initial conditions on the next three lines. Recompile. Run bb.
	// Hints: Denver=1609  Mexico City=2420  La Paz=3650  Everest=8850
	initialAltitude := 1609.0 // meters
	initialAngle := 40.0      // degrees from horizontal
	initialVelocity := 35.0   // m/s
	fmt.Println("      t           x           y          vx           vy          ax         ay")
	history = trajectory.Trajectory(initialAltitude, initialVelocity,
                initialAngle, dt, normalized)
	for k = 0; k < len(history); k++ {
		fmt.Printf("%9.2f %11.2f %11.2f %11.2f %11.2f %11.2f %11.2f \n",
                        history[k].Time,
                        history[k].Position[0],
                        history[k].Position[1],
                        history[k].Velocity[0],
                        history[k].Velocity[1],
                        history[k].Acceleration[0],
                        history[k].Acceleration[1])
	}
	fmt.Println(" End of Baseball")
}
