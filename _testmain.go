package main

import "qprintable"
import "testing"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"qprintable.TestUniversalEncode", qprintable.TestUniversalEncode},
	{"qprintable.TestEOLEncode", qprintable.TestEOLEncode},
	{"qprintable.TestWrapEncode", qprintable.TestWrapEncode},
	{"qprintable.TestUniversalDecode", qprintable.TestUniversalDecode},
	{"qprintable.TestEOLDecode", qprintable.TestEOLDecode},
	{"qprintable.TestWrapDecode", qprintable.TestWrapDecode},
}
var benchmarks = []testing.InternalBenchmark{}

func main() {
	testing.Main(__regexp__.MatchString, tests)
	testing.RunBenchmarks(__regexp__.MatchString, benchmarks)
}
