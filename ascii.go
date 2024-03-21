package main

import (
    "fmt"
    "image"
    _"image/png"
    _"image/jpeg"
    "os"
    "strconv"
    "github.com/fatih/color"
)

func determineColor(pRed float64, pGreen float64, pBlue float64, input string) string{
    red := color.New(color.FgRed).SprintFunc()
    blue := color.New(color.FgBlue).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()
    yellow := color.New(color.FgYellow).SprintFunc()
    magenta := color.New(color.FgMagenta).SprintFunc()
    cyan := color.New(color.FgCyan).SprintFunc()
    white := color.New(color.FgWhite).SprintFunc()

    if pRed >= 0.6 {
        return red(input)
    } else if pGreen >= 0.6 {
        return green(input)
    } else if pBlue >= 0.6 {
        return blue(input)
    } else if pRed + pGreen >= 0.7 {
        return yellow(input)
    } else if pRed + pBlue >= 0.7 {
        return magenta(input)
    } else if pGreen + pBlue >= 0.7 {
        return cyan(input)
    } else {
        return white(input)
    }
}

func main() {
    // Checking a command-line argument was passed
    if len(os.Args) < 2 {
        fmt.Println("Please input file name")
        os.Exit(1)
    }

    // Reading the file name
    fileName := os.Args[1]

    // Opening the file
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)
    }

    defer file.Close()
    inputNumber := 5
    if len(os.Args) > 2 {
        inputNumber, err = strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Error reading condense factor:", err)
            os.Exit(1)
        }
    }
    
    // Decoding the file
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error decoding file:", err)
        os.Exit(1)
    }


    // getting the bounds for the image
    bounds := img.Bounds()
    condenseFactor := uint32(inputNumber)
    
    // colors


    // Iterating through each pixel
    for y := bounds.Min.Y; y < bounds.Max.Y; y+= int(condenseFactor) {
        line := ""
        for x := bounds.Min.X; x < bounds.Max.X; x+= int(condenseFactor) {
            var sum uint32 = 0
            var sumRed uint32 = 0
            var sumGreen uint32 = 0
            var sumBlue uint32 = 0
            for i := 0; i < int(condenseFactor); i++ {
                pixel := img.At(x + i, y + i)
                r, g, b, _ := pixel.RGBA()
                r = r >> 8
                g = g >> 8
                b = b >> 8
                sumRed += r
                sumGreen += g
                sumBlue += b
                sum += r + g + b
            }

            var pRed float64 = 0.0
            var pGreen float64 = 0.0
            var pBlue float64 = 0.0
            if sum != 0 {
                pRed = float64(sumRed)/float64(sum)
                pGreen = float64(sumGreen)/float64(sum)
                pBlue = float64(sumBlue)/float64(sum)
            }

            if sum > 700*condenseFactor {
               line += determineColor(pRed, pGreen, pBlue, "@@")
            } else if sum > 600*condenseFactor {
               line += determineColor(pRed, pGreen, pBlue, "%%")
            } else if sum > 500*condenseFactor {
               line += determineColor(pRed, pGreen, pBlue, "&&")
            } else if sum > 400*condenseFactor {
               line += determineColor(pRed, pGreen, pBlue, "==")
            } else if sum > 300*condenseFactor {
               line += determineColor(pRed, pGreen, pBlue, "**")
            } else if sum > 100*condenseFactor {
               line += determineColor(pRed, pGreen, pBlue, "^^")
            } else {
               line += determineColor(pRed, pGreen, pBlue, "``")
            }
        }
        fmt.Println(line)
    }
}   

