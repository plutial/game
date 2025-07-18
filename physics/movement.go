package physics

type Jump struct {
	// Jump "buffers" (like Coyote time)
	AirTime        int
	JumpRegistered int

	// Number of jumps available
	Jumps int
}

func BodyMove(force *Force, moveLeft bool, moveRight bool) {
	if moveLeft {
		force.Acceleration.X = max(-force.Speed, force.Acceleration.X-0.3)
	}

	if moveRight {
		force.Acceleration.X = min(force.Speed, force.Acceleration.X+0.3)
	}

	if !moveLeft && !moveRight {
		if force.Acceleration.X < 0 {
			force.Acceleration.X = min(0, force.Acceleration.X+0.6)
		} else {
			force.Acceleration.X = max(0, force.Acceleration.X-0.6)
		}
	}
}

func BodyJump(force *Force, jump *Jump, jumpPressed bool) {
	if jumpPressed {
		// Register a jump
		// A jump can be registered even if the body has not yet touched the ground
		jump.JumpRegistered = 3

		// Fix to enable multiple jumps
		// Currently not working, coyote time doesn't work after using this fix
		// jump.AirTime = 0
	}

	// Lower the gravity if the body is on the ground
	if force.Collisions.Down {
		// Set the gravity
		force.Acceleration.Y = min(5, force.Acceleration.Y)
	}

	if force.Collisions.Up {
		// Reset the velocity when the body bonks on its top
		force.Acceleration.Y = max(0, force.Acceleration.Y)
	}

	// Body hits the ground
	if force.Collisions.Down {
		// How long the body has been in the air for
		jump.AirTime = 0

		// Numbers of jumps possible at a time
		// 2 for double jump, 3 for triple jump ...
		// Currently not working but I'm not bother to fix it
		jump.Jumps = 1
	} else {
		// If the body is not touching the ground, it's in the air
		jump.AirTime += 1
	}

	// If the body does not hit the ground in time, it won't jump
	if jump.JumpRegistered > 0 {
		// If the body can jump, and if it is on the ground (kind of... coyote time)
		if jump.Jumps > 0 && jump.AirTime < 5 {
			// How high it goes (and the actual jump part)
			force.Acceleration.Y = -5

			// Take off an available jump
			jump.Jumps -= 1

			// Stop registering the jumps
			jump.JumpRegistered = 0
		} else {
			// Tick down the timer of the register
			jump.JumpRegistered -= 1
		}
	}
}

func BodyDash(force *Force, moveLeft, moveRight, dash bool) {
	if dash && moveRight {
		force.Velocity.X = 30
	} else if dash && moveLeft {
		force.Velocity.X = -30
	}
}
