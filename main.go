package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
	WIDTH         = 64
	MAP_SIZE      = 24
)

type WorldMap [MAP_SIZE][MAP_SIZE]int

var (
	worldMap = WorldMap{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 0, 0, 0, 5, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	playerX    = 22.0
	playerY    = 12.0
	playerDirX = -1.0
	playerDirY = 0.0
	planeX     = 0.0
	planeY     = 0.66
	lineColor  rl.Color
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raycaster")
	rl.SetTargetFPS(60)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		dt := float64(rl.GetFrameTime())
		rotSpeed := dt

		if rl.IsKeyDown(rl.KeyW) {
			nextPosX := playerX + playerDirX*0.1
			nextPosY := playerY + playerDirY*0.1
			if nextPosX > 2 && nextPosX < MAP_SIZE {
				playerX = nextPosX

			}

			if nextPosY > 2 && nextPosY < MAP_SIZE {
				playerY = nextPosY
			}
		}

		if rl.IsKeyDown(rl.KeyS) {
			nextPosX := playerX - playerDirX*0.1
			nextPosY := playerY - playerDirY*0.1

			if nextPosX < MAP_SIZE-2 {
				playerX = nextPosX
			}

			if nextPosY < MAP_SIZE-2 {
				playerY = nextPosY
			}

		}

		if rl.IsKeyDown(rl.KeyD) {
			oldDirX := playerDirX
			playerDirX = playerDirX*math.Cos(-rotSpeed) - playerDirY*math.Sin(-rotSpeed)
			playerDirY = oldDirX*math.Sin(-rotSpeed) + playerDirY*math.Cos(-rotSpeed)
			oldPlaneX := planeX
			planeX = planeX*math.Cos(-rotSpeed) - planeY*math.Sin(-rotSpeed)
			planeY = oldPlaneX*math.Sin(-rotSpeed) + planeY*math.Cos(-rotSpeed)

		}

		if rl.IsKeyDown(rl.KeyA) {
			oldDirX := playerDirX
			playerDirX = playerDirX*math.Cos(rotSpeed) - playerDirY*math.Sin(rotSpeed)
			playerDirY = oldDirX*math.Sin(rotSpeed) + playerDirY*math.Cos(rotSpeed)
			oldPlaneX := planeX
			planeX = planeX*math.Cos(rotSpeed) - planeY*math.Sin(rotSpeed)
			planeY = oldPlaneX*math.Sin(rotSpeed) + planeY*math.Cos(rotSpeed)
		}

		rl.BeginDrawing()
		{

			rl.ClearBackground(rl.Black)
			for x := 0; x < SCREEN_WIDTH; x++ {
				var (
					cameraX = 2*float64(x)/float64(WIDTH) - 1
					rayDirX = playerDirX + planeX*cameraX
					rayDirY = playerDirY + planeY*cameraX

					deltaDistX   = math.Abs(1 / rayDirX)
					deltaDistY   = math.Abs(1 / rayDirY)
					perpWallDist float64
					stepX        int
					stepY        int
					hit          = 0
					side         int
					sideDistX    float64
					sideDistY    float64
					mapX         = int(playerX)
					mapY         = int(playerY)
				)

				if rayDirX < 0 {
					stepX = -1
					sideDistX = (playerX - float64(mapX)) * deltaDistX
				} else {
					stepX = 1
					sideDistX = (float64(mapX) + 1.0 - playerX) * deltaDistX
				}

				if rayDirY < 0 {
					stepY = -1
					sideDistY = (playerY - float64(mapY)) * deltaDistY
				} else {
					stepY = 1
					sideDistY = (float64(mapY) + 1.0 - playerY) * deltaDistY
				}

				for hit == 0 {
					if sideDistX < sideDistY {
						sideDistX += deltaDistX
						mapX += stepX
						side = 0
					} else {
						sideDistY += deltaDistY
						mapY += stepY
						side = 1
					}

					if worldMap[mapX][mapY] > 0 {
						hit = 1
					}
				}

				if side == 0 {
					perpWallDist = (float64(mapX) - playerX + (1-float64(stepX))/2) / rayDirX
				} else {
					perpWallDist = (float64(mapY) - playerY + (1-float64(stepY))/2) / rayDirY
				}

				h := float64(SCREEN_HEIGHT)

				lineHeight := h / float64(perpWallDist)
				drawStart := float64(-lineHeight)/2 + float64(h/2)
				draweEnd := float64(lineHeight)/2 + float64(h/2)
				if drawStart < 0 {
					drawStart = 0
				}

				if draweEnd < 0 {
					draweEnd = 0
				}
				switch worldMap[mapX][mapY] {
				case 1:
					lineColor = rl.Red
				case 2:
					lineColor = rl.Green
				case 3:
					lineColor = rl.Blue
				case 4:
					lineColor = rl.Brown
				default:
					lineColor = rl.Yellow
				}

				if side == 1 {
					lineColor = rl.ColorBrightness(lineColor, 0.5)
				}

				rl.DrawLine(int32(x), int32(drawStart), int32(x), int32(draweEnd), lineColor)
			}

		}
		rl.EndDrawing()
	}
}
