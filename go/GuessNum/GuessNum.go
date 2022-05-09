package main

//3 poziomy skończone i początki 4, przez co jest zamienna hallOfFame która nic nie robi....
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type person struct {
	name   string
	points int
}

func Gra() person {
	// generowanie losowego wyniku
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	result := r1.Intn(100)
	//powitanie i zmienne
	fmt.Println("Zgadnij wylosowaną liczbę! (od 1 do 100)")
	fmt.Println("Możesz zakończyć program komendą \"koniec\" :(")
	var odpNick string
	var odp string
	i := true
	// loop programu
	for i {
		fmt.Print("Podaj liczbę: ")
		fmt.Scanf("%s\n", &odp)
		intOdp, _ := strconv.Atoi(odp)
		switch {
		case odp == "koniec":
			i = false
			fmt.Println("Wychodze z programu :(")
		case intOdp == 0:
			fmt.Println("bład")
		case intOdp == result:
			i = false
			fmt.Println("Zgadłeś!")
			fmt.Println("Podaj swój nick: ")
			fmt.Scanf("%s\n", &odpNick)
			return person{odpNick, 0}
		case intOdp > result:
			fmt.Println("Za dużo!")
		case intOdp < result:
			fmt.Println("Za mało!")
		default:
			fmt.Println("błąd")
		}
	}
	return person{}
}

func main() {
	Gra()
	j := true
	var odp2 string
	for j {
		fmt.Println("Czy gramy jeszcze raz? [T/N]")
		fmt.Scanf("%s\n", &odp2)
		switch {
		case odp2 == "T":
			Gra()
		case odp2 == "N":
			fmt.Println("Wychodze z programu :<")
			j = false
		default:
			fmt.Println("Błędna odpowiedź, wychodze z programu bo sie obraziłem :|")
			j = false
		}
	}
}
