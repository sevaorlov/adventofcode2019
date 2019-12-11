package day10

import (
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	asteroidStr = "#"
	asteroidID  = 1
	upAngle = 3.141592653589793
	betAsteroid = 200
)

type Point struct {
	x, y int
}

func Part1(filename string) string {
	asteroidMap := readMap(filename)

	max, _ := bestAsteroidByVisibleAsteroids(asteroidMap)

	return strconv.Itoa(max)
}

func Part2(filename string) string {
	asteroidMap := readMap(filename)

	_, point := bestAsteroidByVisibleAsteroids(asteroidMap)
	byAngles := asteroidsByAngles(asteroidMap, point.x, point.y)
	angles := sortedAngles(byAngles)

	var pointsSorted map[float64]bool

	destroyedCount := 0
	i := 0
	started := false

	for {
		angle := angles[i]

		i++
		if i >= len(angles) {
			i = 0
		}

		if angle == upAngle {
			started = true
		}

		if !started {
			continue
		}

		if len(byAngles[angle]) > 0 {
			if !pointsSorted[angle] {
				sortPointsSlice(byAngles[angle], point)
			}

			point := popPoints(byAngles[angle])
			destroyedCount++

			if destroyedCount == betAsteroid {
				return strconv.Itoa(point.x * 100 + point.y)
			}
		}
	}
}

func popPoints(points []Point) Point {
	var point Point
	point, points = points[len(points)-1], points[:len(points)-1]
	return point
}

func sortedAngles(asteroidMapByAngles map[float64][]Point) []float64 {
	keys := make([]float64, 0, len(asteroidMapByAngles))
	for k := range asteroidMapByAngles {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	return keys
}

func bestAsteroidByVisibleAsteroids(asteroidMap [][]int) (int, Point) {
	max := 0
	var x, y int

	for i := 0; i < len(asteroidMap); i++ {
		for j := 0; j < len(asteroidMap[i]); j++ {
			if asteroidMap[i][j] == asteroidID {
				count := len(asteroidsByAngles(asteroidMap, j, i))
				if count > max {
					x = j
					y = i
					max = count
				}
			}

		}
	}

	logger.Debug(asteroidMap)
	logger.Debug("best", x, y)

	return max, Point{x, y}
}

func asteroidsByAngles(asteroidMap [][]int, x, y int) map[float64][]Point {
	asteroids := make(map[float64][]Point)

	for y1 := 0; y1 < len(asteroidMap); y1++ {
		for x1 := 0; x1 < len(asteroidMap[y1]); x1++ {
			dx := x1 - x
			dy := y1 - y

			if x1 == x && y1 == y {
				continue
			}

			if asteroidMap[y1][x1] != asteroidID {
				continue
			}

			angle := math.Atan2(0, 1) - math.Atan2(float64(dx), float64(dy))
			if angle < 0 {
				angle += 2 * math.Pi
			}

			asteroids[angle] = append(asteroids[angle], Point{x1, y1})
		}
	}

	return asteroids
}

func sortPointsSlice(points []Point, point Point) {
	sort.Slice(points, func(i, j int) bool {
		return distanceBetweenPoints(points[i], point) < distanceBetweenPoints(points[j], point)
	})
}

func distanceBetweenPoints(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x - p2.x), 2) + math.Pow(float64(p1.y - p2.y), 2))
}

func readAllLines(filename string) []string {
	var lines []string
	input.ReadFile(filename, func(line string) {
		lines = append(lines, line)
	})
	return lines
}

func readMap(filename string) [][]int {
	lines := readAllLines(filename)

	asteroidMap := make([][]int, len(lines))

	for i, line := range lines {
		asteroidMap[i] = make([]int, len(line))
		strA := strings.Split(line, "")

		for j, item := range strA {
			if item == asteroidStr {
				asteroidMap[i][j] = 1
			}
		}
	}

	return asteroidMap
}
