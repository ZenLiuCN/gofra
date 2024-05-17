package units

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/ZenLiuCN/fn"
	"golang.org/x/exp/constraints"
	"io"
	"net/http"
)

func ReadGabsProperty[T constraints.Ordered | ~bool](g *gabs.Container, p string, def ...T) (v T, ok bool) {
	if g != nil && g.ExistsP(p) {
		v, ok = g.Path(p).Data().(T)
		return
	}
	if len(def) != 0 {
		return def[0], false
	}
	return
}
func ReadGabsIndex[T constraints.Ordered | ~bool](g *gabs.Container, p int, def ...T) (v T, ok bool) {
	if g != nil && g.Index(p) != nil {
		v, ok = g.Index(p).Data().(T)
		return
	}
	if len(def) != 0 {
		return def[0], false
	}
	return
}
func ReadGabsPropertySlice[T constraints.Ordered | ~bool](g *gabs.Container, p string, def ...T) (v []T, ok bool) {
	if g != nil && g.ExistsP(p) {
		var a []any
		a, ok = g.Path(p).Data().([]any)
		if !ok {
			return
		}
		var f int
		v, f = fn.ConvAnySlice[T](a)
		if ok = f <= -1; !ok {
			return
		}
	}
	if len(def) != 0 {
		return def, false
	}
	return
}
func ReadGabsIndexSlice[T constraints.Ordered | ~bool](g *gabs.Container, p int, def ...T) (v []T, ok bool) {
	if g != nil && g.Index(p) != nil {
		var a []any
		a, ok = g.Index(p).Data().([]any)
		if !ok {
			return
		}
		var f int
		v, f = fn.ConvAnySlice[T](a)
		if ok = f <= -1; !ok {
			return
		}
		return
	}
	if len(def) != 0 {
		return def, false
	}
	return
}
func n2n[V any, F, T constraints.Integer | constraints.Float](fn func(V, string, ...F) (F, bool)) func(V, string, ...T) (T, bool) {
	return func(v V, p string, t ...T) (o T, ok bool) {
		var x F
		if len(t) == 0 {
			x, ok = fn(v, p)
			if !ok {
				return
			}
			return T(x), true
		}
		var f = make([]F, len(t))
		for i, t2 := range t {
			f[i] = F(t2)
		}
		x, ok = fn(v, p, f...)
		if !ok {
			return
		}
		return T(x), true
	}
}
func ns2ns[V any, F, T constraints.Integer | constraints.Float, FS ~[]F, TS ~[]T](fn func(V, string, ...F) (FS, bool)) func(V, string, ...T) (TS, bool) {
	return func(v V, p string, t ...T) (o TS, ok bool) {
		var x FS
		if len(t) == 0 {
			x, ok = fn(v, p)
			if !ok {
				return
			}
			if len(x) == 0 {
				return
			}
			o = make(TS, len(x))
			for i, f := range x {
				o[i] = T(f)
			}
		}
		var fs = make([]F, len(t))
		for i, t2 := range t {
			fs[i] = F(t2)
		}
		x, ok = fn(v, p, fs...)
		if !ok {
			return
		}
		if len(x) == 0 {
			return
		}
		o = make(TS, len(x))
		for i, f := range x {
			o[i] = T(f)
		}
		ok = true
		return
	}
}

var (
	GabsString   = ReadGabsProperty[string]
	GabsStrings  = ReadGabsPropertySlice[string]
	GabsBoolean  = ReadGabsProperty[bool]
	GabsBooleans = ReadGabsPropertySlice[bool]
	GabsNumber   = ReadGabsProperty[float64]
	GabsNumbers  = ReadGabsPropertySlice[float64]
	GabsInteger  = n2n[*gabs.Container, float64, int](GabsNumber)
	GabsIntegers = ns2ns[*gabs.Container, float64, int, []float64, []int](GabsNumbers)
)

type Gabs struct {
	*gabs.Container
}

func (g Gabs) String(p string, def ...string) (string, bool) {
	return GabsString(g.Container, p, def...)
}
func (g Gabs) Boolean(p string, def ...bool) (bool, bool) {
	return GabsBoolean(g.Container, p, def...)
}
func (g Gabs) Number(p string, def ...float64) (float64, bool) {
	return GabsNumber(g.Container, p, def...)
}
func (g Gabs) Integer(p string, def ...int) (int, bool) {
	return GabsInteger(g.Container, p, def...)
}
func (g Gabs) Strings(p string, def ...string) ([]string, bool) {
	return GabsStrings(g.Container, p, def...)
}
func (g Gabs) Booleans(p string, def ...bool) ([]bool, bool) {
	return GabsBooleans(g.Container, p, def...)
}
func (g Gabs) Numbers(p string, def ...float64) ([]float64, bool) {
	return GabsNumbers(g.Container, p, def...)
}
func (g Gabs) Integers(p string, def ...int) ([]int, bool) {
	return GabsIntegers(g.Container, p, def...)
}

func ReadGabs(r *http.Request) (g Gabs) {
	fn.IgnoreClose(r.Body)
	buf := fn.GetBuffer()
	defer fn.PutBuffer(buf)
	fn.Panic1(io.Copy(buf, r.Body))
	g = Gabs{fn.Panic1(gabs.ParseJSONBuffer(buf))}
	return
}
