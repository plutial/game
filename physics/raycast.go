package physics

import (
	// Raylib
    rl "github.com/gen2brain/raylib-go/raylib"
)

func RayVsBody(start rl.Vector2, velocity rl.Vector2, body Body) (collision bool, hitTime float32, contactNormal rl.Vector2) {
    // Calculate near and far distance
    near := rl.NewVector2(0, 0)
    near.X = (body.Position.X - start.X) / velocity.X
    near.Y = (body.Position.Y - start.Y) / velocity.Y

    far := rl.NewVector2(0, 0)
    far.X = (body.Position.X + body.Size.X - start.X) / velocity.X
    far.Y = (body.Position.Y + body.Size.Y - start.Y) / velocity.Y

    // Sort the distances
    if near.X > far.X {
        near.X, far.X = far.X, near.X
    }
    if near.Y > far.Y {
        near.Y, far.Y = far.Y, near.Y
    }

    // Check if the ray goes through the body
    if near.X > far.Y || near.Y > far.X {
        // Hit time is 1.0 when the ray does not go through the body 
        // As it means that the ray would have gone through the body fully
        // The contact normal would also be (0.0, 0.0)
        // As it means that there was no collision
        return false, 1.0, rl.NewVector2(0, 0)
    }

    // If the collision is somehow behind the ray, return false
    hitFar := min(far.X, far.Y)
    if hitFar < 0 {
        return false, 1.0, rl.NewVector2(0, 0)
    }

    // Get the hit time and the normal if there was a collision
    hitTime = max(near.X, near.Y)

    // Contact normal
    contactNormal = rl.NewVector2(0, 0)
   
    if (near.X > near.Y) {
        if (velocity.X < 0) {
            contactNormal.X = 1
            contactNormal.Y = 0
        } else {
            contactNormal.X = -1
            contactNormal.Y = 0
        }
    } else {
        if (velocity.Y < 0) {
            contactNormal.X = 0
            contactNormal.Y = 1
        } else {
            contactNormal.X = 0
            contactNormal.Y = -1
        }
    }

    return true, hitTime, contactNormal
}
