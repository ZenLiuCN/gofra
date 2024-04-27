package telemetry

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	buf := new(bytes.Buffer)
	s3 := new(strings.Builder)
	s4 := new(strings.Builder)
	s5 := new(strings.Builder)
	s6 := new(strings.Builder)
	s7 := new(strings.Builder)
	s8 := new(strings.Builder)
	for args := 0; args < 10; args++ {
		for ret := 0; ret < 10; ret++ {
			switch {
			case ret == 0 && args == 0:
				{
					_, _ = fmt.Fprintf(buf, `
func Use%[1]d%[2]d(fn func(context.Context), name string, opts ...trace.SpanStartOption) func(context.Context) {
	return func(ctx context.Context) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx)
		} else {
			fn(ctx)
		}
	}
}
func UseErr%[1]d%[2]d(fn func(context.Context) error, name string, opts ...trace.SpanStartOption) func(context.Context) error {
	return func(ctx context.Context) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x:=r.(type) {
					case error:
						err = x
					case string:
						err=errors.New(x)
					default:
						err=fmt.Errorf("%%v",x)
					}
				}
			}()
			err = fn(cx)
		} else {
			err = fn(ctx)
		}
		return
	}
}
`, args, ret)
				}
			case ret == 0:
				{
					s3.Reset()
					s4.Reset()
					s5.Reset()
					for i := 1; i <= args; i++ {
						if s3.Len() != 0 {
							s3.WriteByte(',')
							s4.WriteByte(',')
							s5.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s3, "A%d", i)
						_, _ = fmt.Fprintf(s4, "a%[1]d A%[1]d", i)
						_, _ = fmt.Fprintf(s5, "a%d", i)
					}
					_, _ = fmt.Fprintf(buf, `
func Use%[1]d%[2]d[%[3]s any](fn func(context.Context,%[3]s), name string, opts ...trace.SpanStartOption) func(context.Context,%[3]s) {
	return func(ctx context.Context,%[4]s) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx,%[5]s)
		} else {
			fn(ctx,%[5]s)
		}
	}
}
func UseErr%[1]d%[2]d[%[3]s any](fn func(context.Context,%[3]s) error, name string, opts ...trace.SpanStartOption) func(context.Context,%[3]s) error {
	return func(ctx context.Context,%[4]s) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x:=r.(type) {
					case error:
						err = x
					case string:
						err=errors.New(x)
					default:
						err=fmt.Errorf("%%v",x)
					}
				}
			}()
			err = fn(cx,%[5]s)
		} else {
			err = fn(ctx,%[5]s)
		}
		return
	}
}
`, args, ret, s3, s4, s5)
				}
			case args == 0:
				{
					s3.Reset()
					s4.Reset()
					s5.Reset()
					for i := 1; i <= ret; i++ {
						if s3.Len() != 0 {
							s3.WriteByte(',')
							s4.WriteByte(',')
							s5.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s3, "R%d", i)
						_, _ = fmt.Fprintf(s4, "r%[1]d R%[1]d", i)
						_, _ = fmt.Fprintf(s5, "r%d", i)
					}
					_, _ = fmt.Fprintf(buf, `
func Use%[1]d%[2]d[%[3]s any](fn func(context.Context)(%[3]s), name string, opts ...trace.SpanStartOption) func(context.Context)(%[3]s) {
	return func(ctx context.Context)(%[4]s) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			%[5]s=fn(cx)
		} else {
			%[5]s=fn(ctx)
		}
		return
	}
}
func UseErr%[1]d%[2]d[%[3]s any](fn func(context.Context)(%[3]s,error), name string, opts ...trace.SpanStartOption) func(context.Context)(%[3]s,error) {
	return func(ctx context.Context) (%[4]s,err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x:=r.(type) {
					case error:
						err = x
					case string:
						err=errors.New(x)
					default:
						err=fmt.Errorf("%%v",x)
					}
				}
			}()
			%[5]s,err = fn(cx)
		} else {
			%[5]s,err = fn(ctx)
		}
		return
	}
}
`, args, ret, s3, s4, s5)
				}
			default:
				{
					s3.Reset()
					s4.Reset()
					s5.Reset()
					s6.Reset()
					s7.Reset()
					s8.Reset()
					for i := 1; i <= args; i++ {
						if s3.Len() != 0 {
							s3.WriteByte(',')
							s4.WriteByte(',')
							s5.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s3, "A%d", i)
						_, _ = fmt.Fprintf(s4, "a%[1]d A%[1]d", i)
						_, _ = fmt.Fprintf(s5, "a%d", i)
					}
					for i := 1; i <= ret; i++ {
						if s6.Len() != 0 {
							s6.WriteByte(',')
							s7.WriteByte(',')
							s8.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s6, "R%d", i)
						_, _ = fmt.Fprintf(s7, "r%[1]d R%[1]d", i)
						_, _ = fmt.Fprintf(s8, "r%d", i)
					}
					_, _ = fmt.Fprintf(buf, `
func Use%[1]d%[2]d[%[3]s,%[6]s any](fn func(context.Context,%[3]s)(%[6]s), name string, opts ...trace.SpanStartOption) func(context.Context,%[3]s)(%[6]s) {
	return func(ctx context.Context,%[4]s)(%[7]s) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			%[8]s=fn(cx,%[5]s)
		} else {
			%[8]s=fn(ctx,%[5]s)
		}
		return
	}
}
func UseErr%[1]d%[2]d[%[3]s,%[6]s any](fn func(context.Context,%[3]s)(%[6]s,error), name string, opts ...trace.SpanStartOption) func(context.Context,%[3]s)(%[6]s,error) {
	return func(ctx context.Context,%[4]s) (%[7]s,err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x:=r.(type) {
					case error:
						err = x
					case string:
						err=errors.New(x)
					default:
						err=fmt.Errorf("%%v",x)
					}
				}
			}()
			%[8]s,err = fn(cx,%[5]s)
		} else {
			%[8]s,err = fn(ctx,%[5]s)
		}
		return
	}
}
`, args, ret, s3, s4, s5, s6, s7, s8)
				}
			}
		}
	}
	println(buf.String())
}
