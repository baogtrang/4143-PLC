package main

import (
    "imagemod/GetPic"
    "imagemod/Grayscale"
    "imagemod/Text"
    "imagemod/Colors"
)


func main() {
    GetPic.Download()
    Grayscale.Convert("downloaded_image.jpg")
    Colors.ApplyColor()
    Text.AddText()
    // ... any other operations you want to perform...
}
