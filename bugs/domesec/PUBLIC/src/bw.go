package main

import "os"
import "fmt"
import "strings"
import "errors"
import "strconv"
import "io/ioutil"

type colour struct {
	r, g, b int
}

type pixelmap struct {
	width, height int
	data          []colour
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("Usage ./sobel in.ppm out.ppm\n")
		return
	}
	data, error := ioutil.ReadFile(args[0])
	if error != nil {
		fmt.Printf("Error reading data from file %s\n", args[0])
		return
	}
	matrix, err := ppm_parse(data)
	if err != nil {
		fmt.Printf("Error, bad formatting for PPM %s\n", err)
		return
	}
	pix, err := luminance_map(matrix)
	if err != nil {
		fmt.Printf("Error when comp, got %s\n", err)
	}
	err = write_to_ppm(pix, args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func Append(slice []colour, el colour) []colour {
	ln := len(slice)
	slice = slice[0 : ln+1]
	slice[ln] = el
	return slice
}

func ppm_parse(raw []byte) (*pixelmap, error) {
	data := string(raw[:len(raw)-1])
	next_lf = strings.Index(data, "\n")
	first_ln = data[:next_lf]
	rem_str = data[next_lf+1:]
	if first_ln != "P3" {
		return nil, errors.New("Bad PPM format, supporting only P3 PPM type")
	}
	//fmt.Printf("%s\n", first_ln)
	next_lf = strings.Index(rem_str, "\n")
	first_ln = rem_str[:next_lf]
	rem_str = rem_str[next_lf+1:]
	parts = strings.Split(first_ln, " ")
	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	matsz := width * height
	ret_matrix = make([]colour, 0, matsz)
	next_lf = strings.Index(rem_str, "\n")
	first_ln = rem_str[:next_lf]
	rem_str = rem_str[next_lf+1:]
	for len(rem_str) > 0 {
		ln = strings.Index(rem_str, "\n")
		ln_data string
		if ln == -1 {
			ln = len(rem_str)
			ln_data = rem_str[:ln]
			rem_str = ""
		} else {
			ln_data = rem_str[:ln]
			rem_str = rem_str[ln+1:]
		}
		//fmt.Printf("Got line %s\n", ln_data)
		var colours_els = strings.Split(strings.TrimSpace(ln_data), " ")
		var pos = 0
		for pos < len(colours_els) {
			r, err := strconv.Atoi(colours_els[pos])
			if err != nil {
				fmt.Printf("Could not convert %s to int (red)\n", colours_els[pos])
				return nil, err
			}
			g, err := strconv.Atoi(colours_els[pos+1])
			if err != nil {
				fmt.Printf("Could not convert %s to int (green)\n", colours_els[pos+1])
				return nil, err
			}
			b, err := strconv.Atoi(colours_els[pos+2])
			if err != nil {
				fmt.Printf("Could not convert %s to int (blue)\n", colours_els[pos+2])
				return nil, err
			}
			ret_matrix = Append(ret_matrix, colour{r, g, b})
			pos += 3
		}
		//fmt.Printf("%s\n", ln_data)
	}
	//fmt.Printf("Made a %d long matrix\n", len(ret_matrix))
	pixmap := pixelmap{width, height, ret_matrix}
	return &pixmap, nil
}

func write_to_ppm(inmap *pixelmap, outpath string) error {
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	f.Write([]byte("P3\n"))
	width_str := strconv.Itoa(inmap.width)
	height_str := strconv.Itoa(inmap.height)
	f.Write([]byte(width_str))
	f.Write([]byte{0x20})
	f.Write([]byte(height_str))
	f.Write([]byte{0x0A})
	f.Write([]byte("255\n"))
	x := 0
	y := 0
	for x <= inmap.width {
		for y <= inmap.height {
			colour, err = pixel_at(inmap, x, y)
			if err != nil {
				return err
			}
			f.Write([]byte(strconv.Itoa(colour.r)))
			f.Write([]byte{0x20})
			f.Write([]byte(strconv.Itoa(colour.g)))
			f.Write([]byte{0x20})
			f.Write([]byte(strconv.Itoa(colour.b)))
			f.Write([]byte{0x20})
			y += 1
		}
		y = 0
		x += 1
	}
	return nil
}

func pixel_at(mat *pixelmap, x, y int) (*colour, error) {
	coord = x*mat.height + y
	//fmt.Printf("Coords (%d, %d) = %d\n", x, y, coord)
	if coord >= len(mat.data) {
		return nil, errors.New("Coordinates are out of bounds")
	}
	return &mat.data[coord], nil
}

func luminance(col colour) colour {
	lum := int(0.2126*float32(col.r) + 0.7152*float32(col.g) + 0.0722*float32(col.b))
	return colour{lum, lum, lum}
}

func luminance_map(inmat *pixelmap) (*pixelmap, error) {
	x = 0
	y = 0
	retmap = make([]colour, 0, len(inmat.data))
	for x <= inmat.width {
		for y <= inmat.height {
			col, err := pixel_at(inmat, x, y)
			if err != nil {
				return nil, err
			}
			retmap = Append(retmap, luminance(*col))
			y += 1
		}
		y = 0
		x += 1
	}
	px = pixelmap{inmat.width, inmat.height, retmap}
	return &px, nil
}