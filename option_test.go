package option

import (
	"bytes"
	"os"
	"testing"

	"github.com/cdelorme/go-maps"
)

func TestPlacebo(t *testing.T) {
	t.Parallel()

	if !true {
		t.FailNow()
	}
}

func TestExample(t *testing.T) {
	t.Parallel()
	o := &App{}

	// test invalid input (empty string)
	o.Example("")
	if len(o.examples) != 0 {
		t.Log("failed to exclude empty-string example")
		t.FailNow()
	}

	// test valid input
	o.Example("test")
	if len(o.examples) == 0 || o.examples[0] != "test" {
		t.Logf("failed to register supplied example: %+v\n", o.examples)
		t.FailNow()
	}
}

func TestFlag(t *testing.T) {
	t.Parallel()
	o := &App{}

	// test three fail-cases: no name, no flags, invalid flags
	o.Flag("", "desc", "--flag")
	o.Flag("name", "desc")
	o.Flag("name", "desc", "badflag")

	// validate no entries registered
	if len(o.options) > 0 {
		t.Log("failed to ignore invalid input")
		t.FailNow()
	}

	// register valid record & validate
	o.Flag("test", "a flag", "-t")
	if len(o.options) != 1 || o.options[0].Name != "test" || o.options[0].Description != "a flag" || len(o.options[0].Flags) != 1 || o.options[0].Flags[0] != "-t" {
		t.Logf("Expected flag (%+v) but found (%+v)\n", o.options[0])
		t.FailNow()
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	// capture original arguments
	originalArgs := os.Args

	// test valid long flags of each type
	lo := &App{NoHelp: true}
	lo.Flag("lb", "long flag boolean", "--bool")
	lo.Flag("ls", "long flag string", "--string")
	lo.Flag("li", "long flag int", "--int")
	lo.Flag("lf", "long flag float", "--float")

	// set values to test
	os.Args = append(originalArgs, "--string=banana", "--float=6.4", "--bool", "--int=32")

	// parse & validate results
	lp := lo.Parse()
	lb, _ := maps.Bool(lp, false, "lb")
	if lb != true {
		t.Log("Long Bool Fail - Expected: true, Got: false")
		t.FailNow()
	}
	ls, _ := maps.String(lp, "", "ls")
	if ls != "banana" {
		t.Logf("Long String Fail - Expected: banana, Got: %s\n", ls)
		t.FailNow()
	}
	li, _ := maps.Int(lp, 0, "li")
	if li != int64(32) {
		t.Logf("Long Int Fail - Expected: 32, Got: %d\n", li)
		t.FailNow()
	}
	lf, _ := maps.Float(lp, 0, "lf")
	if lf != float64(6.4) {
		t.Logf("Long Float Fail - Expected: 32, Got: %f\n", lf)
		t.FailNow()
	}

	// test valid short flags of each type
	so := &App{NoHelp: true}
	so.Flag("sb", "short flag boolean", "-b")
	so.Flag("ss", "short flag string", "-s")
	so.Flag("si", "short flag int", "-i")
	so.Flag("sf", "short flag float", "-f")

	// set values to test
	os.Args = append(originalArgs, "-s", "banana", "-f", "6.4", "-b", "-i", "32")

	// parse & validate results
	sp := so.Parse()
	sb, _ := maps.Bool(sp, false, "sb")
	if sb != true {
		t.Log("Short Bool Fail - Expected: true, Got: false")
		t.FailNow()
	}
	ss, _ := maps.String(sp, "", "ss")
	if ss != "banana" {
		t.Logf("Short String Fail - Expected: banana, Got: %s\n", ss)
		t.FailNow()
	}
	si, _ := maps.Int(sp, 0, "si")
	if si != int64(32) {
		t.Logf("Short Int Fail - Expected: 32, Got: %d\n", si)
		t.FailNow()
	}
	sf, _ := maps.Float(sp, 0, "sf")
	if sf != float64(6.4) {
		t.Logf("Short Float Fail - Expected: 32, Got: %f\n", sf)
		t.FailNow()
	}

	// test combination short & long flags of each type

	// test valid short flags of each type
	co := &App{NoHelp: true}
	co.Flag("cb", "combo (short) flag boolean", "-b")
	co.Flag("cs", "combo (long) flag string", "--string")
	co.Flag("ci", "combo (short) flag int", "-i")
	co.Flag("cf", "combo (long) flag float", "--float")

	// set values to test
	os.Args = append(originalArgs, "-i", "32", "--string=banana", "--float=6.4", "-b")

	// parse & validate results
	cp := co.Parse()
	cb, _ := maps.Bool(cp, false, "cb")
	if cb != true {
		t.Log("Combo Bool Fail - Expected: true, Got: false")
		t.FailNow()
	}
	cs, _ := maps.String(cp, "", "cs")
	if cs != "banana" {
		t.Logf("Combo String Fail - Expected: banana, Got: %s\n", cs)
		t.FailNow()
	}
	ci, _ := maps.Int(cp, 0, "ci")
	if ci != int64(32) {
		t.Logf("Combo Int Fail - Expected: 32, Got: %d\n", ci)
		t.FailNow()
	}
	cf, _ := maps.Float(cp, 0, "cf")
	if cf != float64(6.4) {
		t.Logf("Combo Float Fail - Expected: 32, Got: %f\n", cf)
		t.FailNow()
	}
}

func TestHelp(t *testing.T) {
	t.Parallel()

	// override stdout and exit func
	var b bytes.Buffer
	stdout = &b
	exit = func(_ int) {}

	// create instance and trigger from forced help flag
	o := App{Description: "Test"}
	o.Flag("name", "desc", "--test")
	o.Example("--test")
	os.Args = append(os.Args, "-h")
	_ = o.Parse()

	// direct literal copy of what we expect to see
	expected := `[go-option.test]: Test

Flags:
help, -h, --help              	display help information
--test                        	desc

Usage:
go-option.test --test
`

	// compare against expected output
	if b.String() != expected {
		t.Log("Help output did not appear as expected...")
		t.FailNow()
	}
}
