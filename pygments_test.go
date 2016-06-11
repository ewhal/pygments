package pygments

import (
	"testing"
)

func TestBinary(t *testing.T) {
	Binary("/usr/bin/pygmentize")
}

func TestWhich(t *testing.T) {
	v := Which()
	e := "/usr/bin/pygmentize"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}

func TestHighlighting(t *testing.T) {
	v := Highlight("print \"Hello World!\"", "python", "html", "nowrap=True,lineseparator=<br/>,", "utf-8")
	e := "<span class=\"k\">print</span> <span class=\"s2\">&quot;Hello World!&quot;</span><br/>"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}
