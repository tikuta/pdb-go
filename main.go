package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Atom struct {
	Serial     int
	Name       string
	AltLoc     string
	ResName    string
	ChainId    string
	ResSeq     int
	InsCode    string
	X          float64
	Y          float64
	Z          float64
	Occupancy  float64
	TempFactor float64
	Element    string
	Charge     string
}

func Read(r io.Reader) (atoms []Atom, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "ATOM") {
			atom := Atom{}
			atom.Serial, err = strconv.Atoi(strings.TrimSpace(s[6:11]))
			atom.Name = strings.TrimSpace(s[12:16])
			atom.AltLoc = strings.TrimSpace(string(s[16]))
			atom.ResName = strings.TrimSpace(s[17:20])
			atom.ChainId = strings.TrimSpace(string(s[21]))
			atom.ResSeq, err = strconv.Atoi(strings.TrimSpace(s[22:26]))
			atom.InsCode = strings.TrimSpace(string(s[26]))
			atom.X, err = strconv.ParseFloat(strings.TrimSpace(s[30:37]), 32)
			atom.Y, err = strconv.ParseFloat(strings.TrimSpace(s[38:46]), 32)
			atom.Z, err = strconv.ParseFloat(strings.TrimSpace(s[46:54]), 32)
			atom.Occupancy, err = strconv.ParseFloat(strings.TrimSpace(s[54:60]), 32)
			atom.TempFactor, err = strconv.ParseFloat(strings.TrimSpace(s[60:66]), 32)
			atom.Element = strings.TrimSpace(s[76:78])
			atom.Charge = strings.TrimSpace(s[78:80])

			if err != nil {
				return nil, err
			}

			atoms = append(atoms, atom)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return atoms, nil
}

func Write(atoms []Atom) {
	for _, e := range atoms {
		s := fmt.Sprintf("ATOM  % 5d % 4s% 1s% 3s % 1s% 4d% 1s   % 8s% 8s% 8s% 6s% 6s          % 2s% 2s", e.Serial, e.Name, e.AltLoc, e.ResName, e.ChainId, e.ResSeq, e.InsCode, fmt.Sprintf("%.3f", e.X), fmt.Sprintf("%.3f", e.Y), fmt.Sprintf("%.3f", e.Z), fmt.Sprintf("%.2f", e.Occupancy), fmt.Sprintf("%.2f", e.TempFactor), e.Element, e.Charge)
		fmt.Println(s)
	}
}

func main() {
	var atoms []Atom

	atoms, err := Read(os.Stdin)
	if err != nil {
		panic(err)
	}

	Write(atoms)
}
