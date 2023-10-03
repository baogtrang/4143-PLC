module imagemod

go 1.21.0

require (
    github.com/fogleman/gg v1.3.0
    golang.org/x/image v0.12.0
)

require (
    github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
)


replace (
    imagemod/GetPic => ./GetPic
    imagemod/Grayscale => ./Grayscale
    imagemod/Text => ./Text
    imagemod/Colors => ./Colors
)
