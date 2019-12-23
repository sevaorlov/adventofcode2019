package day12

import (
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

const numIteration = 1000

type Coords struct {
	x, y, z int
}

func (c *Coords) string() string {
	return fmt.Sprintf("%d%d%d", c.x, c.y, c.z)
}

type Moon struct {
	position Coords
	velocity Coords
}

func (m *Moon) applyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m *Moon) potentialEnergy() int {
	return int(math.Abs(float64(m.position.x)) + math.Abs(float64(m.position.y)) + math.Abs(float64(m.position.z)))
}

func (m *Moon) kineticEnergy() int {
	return int(math.Abs(float64(m.velocity.x)) + math.Abs(float64(m.velocity.y)) + math.Abs(float64(m.velocity.z)))
}

func (m *Moon) energy() int {
	return m.potentialEnergy() * m.kineticEnergy()
}

func (m *Moon) equalPosition(other Moon) bool {
	return m.position.x == other.position.x && m.position.y == other.position.y && m.position.z == other.position.z
}

func (m *Moon) equalVelocity(other Moon) bool {
	return m.velocity.x == other.velocity.x && m.velocity.y == other.velocity.y && m.velocity.z == other.velocity.z
}

func (m *Moon) string() string {
	return fmt.Sprintf("%s%s", m.position.string(), m.velocity.string())
}

func moonsFromFile(filename string) []Moon {
	var moons []Moon

	input.ReadFile(filename, func(line string) {
		re := regexp.MustCompile(`=(-?\d+)`)

		submatches := re.FindAllStringSubmatch(line, -1)
		x, err := strconv.Atoi(submatches[0][1])
		if err != nil {
			log.Fatal("failed to parse coordinate x")
		}
		y, err := strconv.Atoi(submatches[1][1])
		if err != nil {
			log.Fatal("failed to parse coordinate x")
		}
		z, err := strconv.Atoi(submatches[2][1])
		if err != nil {
			log.Fatal("failed to parse coordinate x")
		}

		moons = append(moons, Moon{position: Coords{x: x, y: y, z: z}})
	})

	return moons
}

func Part1(filename string) string {
	moons := moonsFromFile(filename)

	i := 0

	for {
		i++
		timeStep(moons)

		if i == numIteration {
			break
		}
	}

	logger.Debug("moons", moons)

	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += moon.energy()
	}

	return strconv.Itoa(totalEnergy)
}

func timeStep(moons []Moon) {
	for n := 0; n < len(moons); n++ {
		for m := n + 1; m < len(moons); m++ {
			if moons[n].position.x > moons[m].position.x {
				moons[n].velocity.x--
				moons[m].velocity.x++
			} else if moons[n].position.x < moons[m].position.x {
				moons[n].velocity.x++
				moons[m].velocity.x--
			}

			if moons[n].position.y > moons[m].position.y {
				moons[n].velocity.y--
				moons[m].velocity.y++
			} else if moons[n].position.y < moons[m].position.y {
				moons[n].velocity.y++
				moons[m].velocity.y--
			}

			if moons[n].position.z > moons[m].position.z {
				moons[n].velocity.z--
				moons[m].velocity.z++
			} else if moons[n].position.z < moons[m].position.z {
				moons[n].velocity.z++
				moons[m].velocity.z--
			}

		}

		moons[n].applyVelocity()
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func Part2(filename string) string {
	moons := moonsFromFile(filename)

	var xPeriod, yPeriod, zPeriod int

	i := 0
	for {
		i++

		timeStep(moons)

		if xPeriod == 0 {
			otherZeros := false
			for _, moon := range moons {
				if moon.velocity.x != 0 {
					otherZeros = true
					break
				}
			}

			if !otherZeros {
				xPeriod = i * 2
			}
		}

		if yPeriod == 0 {
			otherZeros := false
			for _, moon := range moons {
				if moon.velocity.y != 0 {
					otherZeros = true
					break
				}
			}

			if !otherZeros {
				yPeriod = i * 2
			}
		}

		if zPeriod == 0 {
			otherZeros := false
			for _, moon := range moons {
				if moon.velocity.z != 0 {
					otherZeros = true
					break
				}
			}

			if !otherZeros {
				zPeriod = i * 2
			}
		}

		if xPeriod != 0 && yPeriod != 0 && zPeriod != 0 {
			break
		}
	}

	logger.Debug(moons)

	return strconv.Itoa(lcm(lcm(xPeriod, yPeriod), zPeriod))
}
