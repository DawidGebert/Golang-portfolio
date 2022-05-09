//zadanie z mrówkami, podnoszą liście jak je spotkają i kładą obok innego liścia, nie mogą wchodzic na siebie ani na liście
//wyświetlanie planszy działa tylko na VSCode, a przynajmniej nie działa w zwykłym terminalu, ponieważ użyłem emoji

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//all parameters
const n = 35
const antCount = 70
const leafCount = 300
const NumberOfTurns = 1500

//helper functions
func contains(list [][]int, cords []int) bool {
	for _, a := range list {
		if a[0] == cords[0] && a[1] == cords[1] {
			return true
		}
	}
	return false
}
func isGood(list []int) bool {
	if list[0] >= 0 && list[0] < n && list[1] >= 0 && list[1] < n {
		return true
	}
	return false
}

func RemoveIndex(s [][]int, index int) [][]int {
	return append(s[:index], s[index+1:]...)
}

func Filter(arr [][]int, cond func(int, int) bool) [][]int {
	result := [][]int{}
	for i := range arr {
		if cond(arr[i][0], arr[i][1]) {
			result = append(result, arr[i])
		}
	}
	return result
}

//one turn of ants movement
func AntsMove(Ants [][]int, Leaves [][]int, turn int) ([][]int, [][]int) {
	for i := 0; i <= len(Ants)-1; i++ {
		ant_copy := []int{Ants[i][0], Ants[i][1], Ants[i][2], Ants[i][3]}
		x := rand.Intn(4)
		switch {
		case x == 0:
			Ants[i][0] += 1
		case x == 1:
			Ants[i][0] -= 1
		case x == 2:
			Ants[i][1] += 1
		case x == 3:
			Ants[i][1] -= 1
		}
		if !isGood(Ants[i]) {
			Ants[i] = ant_copy
		}
		if Ants[i][3] > 0 {
			Ants[i][3] -= 1
		}
		if contains(Ants[:i], Ants[i]) || contains(Ants[i+1:], Ants[i]) {
			Ants[i] = ant_copy
		}
		// jeśli mrówka spotka liścia
		if contains(Leaves, []int{Ants[i][0], Ants[i][1]}) {
			state := Ants[i][2]
			//podnosi go
			if Ants[i][2] == 0 && Ants[i][3] < 1 && turn <= NumberOfTurns-20 {
				state = 1
				Ants[i][3] = 7
				Leaves = Filter(Leaves, func(val1 int, val2 int) bool {
					return !(val1 == Ants[i][0] && val2 == Ants[i][1])
				})
			}
			//kładzie obok drugiego napotkanego
			if Ants[i][2] == 1 && Ants[i][3] < 1 {
				state = 0
				Ants[i] = ant_copy
				Ants[i][3] = 7
				target_leaf := []int{Ants[i][0], Ants[i][1]}
				j := true
				for j {
					x = rand.Intn(8)
					switch x {
					case 0:
						bonus_leaf := []int{target_leaf[0] - 1, target_leaf[1]}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 1:
						bonus_leaf := []int{target_leaf[0] + 1, target_leaf[1]}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 2:
						bonus_leaf := []int{target_leaf[0] - 1, target_leaf[1] - 1}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 3:
						bonus_leaf := []int{target_leaf[0] + 1, target_leaf[1] - 1}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 4:
						bonus_leaf := []int{target_leaf[0], target_leaf[1] - 1}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 5:
						bonus_leaf := []int{target_leaf[0] - 1, target_leaf[1] + 1}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 6:
						bonus_leaf := []int{target_leaf[0] + 1, target_leaf[1] + 1}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					case 7:
						bonus_leaf := []int{target_leaf[0], target_leaf[1] + 1}
						if !contains(Leaves, bonus_leaf) && isGood(bonus_leaf) {
							Leaves = append(Leaves, bonus_leaf)
							j = false
						}
					}
				}
			}
			Ants[i][2] = state
		}
	}
	return Ants, Leaves
}

var Ants [][]int
var Leaves [][]int

func main() {
	rand.Seed(time.Now().UnixNano())
	var pole [n][n]string

	//generate list of leaves and ants
	for i := 1; i <= leafCount; i++ {
		x := rand.Intn(n)
		y := rand.Intn(n)
		leaf := []int{x, y}
		if !contains(Leaves, leaf) {
			Leaves = append(Leaves, leaf)
		}
	}
	for i := 1; i <= antCount; i++ {
		x := rand.Intn(n)
		y := rand.Intn(n)
		// x  ,  y  ,  hasLeaf?   ,   turnsTired
		ant := []int{x, y, 0, 0}
		if !contains(Ants, ant) {
			Ants = append(Ants, ant)
		}
		Leaves = Filter(Leaves, func(val1 int, val2 int) bool {
			return !(val1 == ant[0] && val2 == ant[1])
		})
	}
	//wygenerowane pole
	for i, v := range pole {
		for j := range v {
			//puste pola
			pole[i][j] = "⬛"
		}
	}
	for _, v := range Leaves {
		//liście
		pole[v[0]][v[1]] = "🟩"
	}
	for _, v := range Ants {
		//mrówki
		pole[v[0]][v[1]] = "🐜"
		if v[2] == 1 {
			//mrówki z liściem
			pole[v[0]][v[1]] = "🦗"
		}
		if v[3] > 0 && v[2] == 0 {
			//zmęczone mrówki
			pole[v[0]][v[1]] = "🟫"
		}
	}
	for _, v := range pole {
		fmt.Println(v)
	}
	//ruch mrówek
	for i := 1; i <= NumberOfTurns; i++ {
		Ants, Leaves = AntsMove(Ants, Leaves, i)
	}
	//pole po wszystkich turach
	fmt.Println("--------------------------------------------------------------")
	for i, v := range pole {
		for j := range v {
			pole[i][j] = "⬛"
		}
	}
	for _, v := range Leaves {
		pole[v[0]][v[1]] = "🟩"
	}
	for _, v := range Ants {
		pole[v[0]][v[1]] = "🐜"
		if v[2] == 1 {
			pole[v[0]][v[1]] = "🦗"
		}
		if v[3] > 0 && v[2] == 0 {
			pole[v[0]][v[1]] = "🟫"
		}
	}
	for _, v := range pole {
		fmt.Println(v)
	}
	// 🦗🟫

}
