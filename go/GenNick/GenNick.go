package main

import (
	"flag"
	"fmt"
)

func main() {
	dayAdjc := map[int]string{
		1: "dark", 2: "the darkest", 3: "sneaky",
		4: "spooky", 5: "evil", 6: "ashy",
		7: "creepy", 8: "spontanious", 9: "creative",
		10: "shadowy", 11: "misplaced", 12: "lost",
		13: "forgotten", 14: "sharp", 15: "elegant",
		16: "misty", 17: "demonic", 18: "ghastly",
		19: "spiritual", 20: "celestial", 21: "divine",
		22: "ghostly", 23: "dreadful", 24: "spectral",
		25: "grusome", 26: "nameless", 27: "fallen",
		28: "astral", 29: "cursed", 30: "deppresed",
		31: "sad",
	}
	nameWord := map[string]string{
		"a": "angel", "b": "devil", "c": "demon",
		"d": "hunter", "e": "destroyer", "f": "killer",
		"g": "swordsman", "h": "socerer", "i": "mage",
		"j": "shaman", "k": "hero", "l": "lord",
		"m": "lich", "n": "archangel", "o": "brute",
		"p": "undead", "r": "creature", "s": "being",
		"t": "conqueror", "u": "assasin", "w": "monk",
		"z": "swordmaster",
	}
	surnameWord := map[string]string{
		"a": "of darkness", "b": "the shadow stealer", "c": "the eater of worlds",
		"d": "of destruction", "e": "the demon hunter", "f": "the skin stealer",
		"g": "the hand of god", "h": "the nameless one", "i": "the one who wonders",
		"j": "the one who knows", "k": "of dark plague", "l": "the shapeless",
		"m": "the world ender", "n": "the world creator", "o": "the mighty one",
		"p": "the king of all beings", "r": "the sculptor of reality", "s": "of scorching sun",
		"t": "the star thats shines", "u": "the mistake in reality", "w": "the rift in normality",
		"z": "of blue sky",
	}
	var dzien int
	var namestring, surnamestring string
	flag.IntVar(&dzien, "day", 0, "dzien urodzenia")
	flag.StringVar(&namestring, "namechar", "", "Pierwsza litera imienia")
	flag.StringVar(&surnamestring, "surnamechar", "", "Pierwsza litera nazwiska")
	flag.Parse()
	fmt.Println("Callsign")
	if dzien == 0 {
		fmt.Print("Podaj dzien urodzenia: ")
		fmt.Scanf("%d\n", &dzien)
	}
	if namestring == "" {
		fmt.Print("Podaj pierwsza litere imienia(lowercase, only english): ")
		fmt.Scanf("%s\n", &namestring)
	}
	if surnamestring == "" {
		fmt.Print("Podaj pierwsza litere nazwiska(lowercase, only english): ")
		fmt.Scanf("%s\n", &surnamestring)
	}
	fmt.Print(dayAdjc[dzien] + " ")
	fmt.Print(nameWord[namestring] + ", ")
	fmt.Print(surnameWord[surnamestring])
}
