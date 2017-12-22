# Percentile

[![Go Report Card](https://goreportcard.com/badge/github.com/clly/percentile)](https://goreportcard.com/report/github.com/clly/percentile)
[![Build Status](https://travis-ci.org/clly/percentile.svg?branch=travis)](https://travis-ci.org/clly/percentile)
[![Godoc](https://godoc.org/github.com/clly/percentile?status.svg)](https://godoc.org/github.com/clly/percentile)
[![License](https://img.shields.io/github/license/clly/percentile.svg)](LICENSE)

Percentile is a unix tool to calculate percentile while on the command line. It accepts a set of float's separated by new lines piped into the tool.
It then calculates a set of percentiles that are given as command line arguments. If no arguments are given it will calculate the median and 95th 
percentile.

```
  --> seq 1 10000| ./percentile 10 15 20.5 50 99 99.999
Pctl 10.000: 1000.00
Pctl 15.000: 1500.00
Pctl 20.500: 2050.00
Pctl 50.000: 5000.00
Pctl 99.000: 9900.00
Pctl 99.999: 10000.00
```
