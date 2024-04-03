package main

import (
	"fmt"
	"math"
	"time"
)

const (
	thetaSpacing  = 0.01
	phiSpacing    = 0.01
	R1            = 1
	R2            = 2
	K2            = 5
	screenWidth   = 80
	screenHeight  = 24
)

func renderFrame(A, B float64) {
	
	K1 := float64(screenWidth) * K2 * 3 / (8 * (R1 + R2))

	cosA, sinA := math.Cos(A), math.Sin(A)
	cosB, sinB := math.Cos(B), math.Sin(B)

	output := make([][]rune, screenWidth)
	zbuffer := make([][]float64, screenWidth)
	for i := range output {
		output[i] = make([]rune, screenHeight)
		zbuffer[i] = make([]float64, screenHeight)
	}

	for theta := 0.0; theta < 2*math.Pi; theta += thetaSpacing {
	
		costheta, sintheta := math.Cos(theta), math.Sin(theta)
	
		for phi := 0.0; phi < 2*math.Pi; phi += phiSpacing {
	
			cosphi, sinphi := math.Cos(phi), math.Sin(phi)

			circlex := R2 + R1*costheta
			circley := R1*sintheta

			x := circlex*(cosB*cosphi+sinA*sinB*sinphi) - circley*cosA*sinB
			y := circlex*(sinB*cosphi-sinA*cosB*sinphi) + circley*cosA*cosB
			z := K2 + cosA*circlex*sinphi + circley*sinA
			ooz := 1 / z 

			// possivel efeito 3d, 2d
			xp := int(screenWidth/2 + K1*ooz*x)
			yp := int(screenHeight/2 - K1*ooz*y)


			if xp >= 0 && xp < screenWidth && yp >= 0 && yp < screenHeight {

				if ooz > zbuffer[xp][yp] {
					zbuffer[xp][yp] = ooz

					L := cosphi*costheta*sinB - cosA*costheta*sinphi - sinA*sintheta + cosB*(cosA*sintheta-costheta*sinA*sinphi)

					luminanceIndex := int((L + 1) * 4)
					if luminanceIndex < 0 {
						luminanceIndex = 0
					} else if luminanceIndex >= len(".,-~:;=!*#$@") {
						luminanceIndex = len(".,-~:;=!*#$@") - 1
					}
					output[xp][yp] = rune(".,-~:;=!*#$@"[luminanceIndex])
				}
			}
		}
	}

	for j := 0; j < screenHeight; j++ {
		for i := 0; i < screenWidth; i++ {
			fmt.Printf("%c", output[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// loop de movimento
	for {

		t := float64(time.Now().UnixNano()) / 1e9

		A := t
		B := t / 2
		renderFrame(A, B)
		time.Sleep(50 * time.Millisecond)
	}
}
