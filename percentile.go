package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/clly/libsys"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/stat"
)

func main() {
	if ok, err := libsys.IsPiped(); err == nil {
		if !ok {
			fmt.Fprintf(os.Stderr, "Expecting content piped to %s\n", os.Args[0])
			os.Exit(1)
		}
	}

	percentiles, err := parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	distribution, err := readToFloat(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, v := range percentiles {
		p := stat.Quantile(v, stat.Empirical, distribution, nil)
		fmt.Printf("Pctl %.3f: %.2f\n", v*100, p)
	}
}

// intended to check os.Args for percentiles to check
func parse(str []string) ([]float64, error) {
	var percentiles []float64
	if len(str) == 0 {
		percentiles = append(percentiles, .50, .95)
	} else {
		return checkStringsForFloat(str)
	}
	return percentiles, nil
}

// intended to take os.Args
func checkStringsForFloat(str []string) ([]float64, error) {
	var args []float64
	for i := 0; i < len(str); i++ {
		// we want 64bit floats
		if fl, err := strconv.ParseFloat(str[i], 64); err == nil {
			if fl < 0 || fl > 100 {
				return nil, fmt.Errorf("percentile out of range. Should be greater than 0 or less than 100: %s", str[i])
			}
			args = append(args, fl/100)
		} else {
			return nil, errors.Wrap(err, "Failed to turn arg into float64 for percentile")
		}
	}
	return args, nil
}

// intended to take os.Stdin
func readToFloat(r io.Reader) ([]float64, error) {
	var distribution []float64
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		if fl, err := strconv.ParseFloat(t, 64); err == nil {
			distribution = append(distribution, fl)
		} else {
			return nil, err
		}
	}
	// sort here why not
	sort.Float64s(distribution)
	return distribution, nil
}
