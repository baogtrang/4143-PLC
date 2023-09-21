package main

import (
	"fmt"
	"imagemod/imageManipulator"
)

func main() {
    // Create an ImageManipulator instance with an existing image
    im, err := imageManipulator.NewImageManipulatorWithImage("pig.jpg")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Draw a rectangle
    im.DrawRectangle(150, 50, 560, 411)

    // Save the modified image to a file
    err = im.SaveToFile("pig_with_rectangle.png")
    if err != nil {
        fmt.Println("Failed to save the modified image:", err)
    }
}
