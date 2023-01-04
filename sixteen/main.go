package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type tunnel struct {
	flowrate  int
	tunnels   []string
	distances map[string][]string
}

type action struct {
	tunnel string
	time   int
}

var tmap map[string]tunnel

func calculateDistance(current string, dest string, route []string) {
	if current == dest {
		newRoute := append(route, current)

		if len(route) > 0 {
			topRoute := tmap[route[0]].distances[dest]
			if len(topRoute) == 0 || len(topRoute) > len(newRoute) {
				tmap[route[0]].distances[dest] = newRoute[1:]
			}
		}

		return
	}

	for _, tunnel := range tmap[current].tunnels {
		visited := false

		for _, past := range route {
			if past == tunnel {
				visited = true
			}
		}

		if !visited {
			newRoute := append(route, current)
			calculateDistance(tunnel, dest, newRoute)
		}
	}
}

func getTmapKeys() []string {
	keys := make([]string, 0, len(tmap))
	for key := range tmap {
		keys = append(keys, key)
	}

	return keys
}

func compileTunnelInfo(src []string) {
	re, err := regexp.Compile(`([A-Z][A-Z]|\d)+`)
	check(err)

	tmap = map[string]tunnel{}

	for _, line := range src {
		values := re.FindAllString(line, -1)

		flowrate, err := strconv.Atoi(values[1])
		check(err)

		tunnel_info := tunnel{
			flowrate:  flowrate,
			tunnels:   values[2:],
			distances: map[string][]string{},
		}

		tmap[string(values[0])] = tunnel_info
	}

	keys := getTmapKeys()

	for _, t := range keys {
		for _, c := range keys {
			calculateDistance(t, c, []string{})
		}
	}
}

func advance(time int, current string, dest string, route []string, actions []action, closed_valves []string) []action {
	if time <= 0 {
		return actions
	}

	duration := 1
	target := dest
	next_route := route
	new_closed_valves := closed_valves

	if current == dest {
		eligible := true
		for _, action := range actions {
			if dest == action.tunnel {
				eligible = false
			}
		}

		if time < 30 && eligible {
			duration = 2
			actions = append(actions, action{
				tunnel: dest,
				time:   time - 1,
			})
		}

		new_closed_valves = []string{}
		for _, valve := range closed_valves {
			if valve != dest {
				new_closed_valves = append(new_closed_valves, valve)
			}
		}

		top_weight := float32(0)
		for dest_key, new_dest := range tmap[dest].distances {
			eligible := false
			for _, valve := range new_closed_valves {
				if valve == dest_key {
					eligible = true
				}
			}

			if eligible {
				squared_distance := math.Pow(float64(len(tmap[dest].distances[dest_key])), 2)
				weight := float32(tmap[dest_key].flowrate) / float32(squared_distance)
				if weight > top_weight {
					top_weight = weight
					target = dest_key
					next_route = new_dest
				}
			}
		}
	} else {
		next_route = next_route[1:]
	}

	return advance(time-duration, next_route[0], target, next_route, actions, new_closed_valves)
}

func main() {
	tunnels, err := os.ReadFile("tunnels.txt")
	check(err)

	src := strings.Split(string(tunnels), "\n")

	compileTunnelInfo(src)

	optimal_route := advance(30, "AA", "AA", []string{}, []action{}, getTmapKeys())

	subtotal := 0
	for i := 30; i > 0; i-- {
		for _, part := range optimal_route {
			if part.time >= i {
				subtotal += tmap[part.tunnel].flowrate
			}
		}
	}

	fmt.Println(optimal_route, subtotal)
}
