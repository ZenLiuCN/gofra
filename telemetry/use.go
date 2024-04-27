package telemetry

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/trace"
)

func Use00(fn func(context.Context), name string, opts ...trace.SpanStartOption) func(context.Context) {
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
func UseErr00(fn func(context.Context) error, name string, opts ...trace.SpanStartOption) func(context.Context) error {
	return func(ctx context.Context) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
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

func Use01[R1 any](fn func(context.Context) R1, name string, opts ...trace.SpanStartOption) func(context.Context) R1 {
	return func(ctx context.Context) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx)
		} else {
			r1 = fn(ctx)
		}
		return
	}
}
func UseErr01[R1 any](fn func(context.Context) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, error) {
	return func(ctx context.Context) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx)
		} else {
			r1, err = fn(ctx)
		}
		return
	}
}

func Use02[R1, R2 any](fn func(context.Context) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2) {
	return func(ctx context.Context) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx)
		} else {
			r1, r2 = fn(ctx)
		}
		return
	}
}
func UseErr02[R1, R2 any](fn func(context.Context) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx)
		} else {
			r1, r2, err = fn(ctx)
		}
		return
	}
}

func Use03[R1, R2, R3 any](fn func(context.Context) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx)
		} else {
			r1, r2, r3 = fn(ctx)
		}
		return
	}
}
func UseErr03[R1, R2, R3 any](fn func(context.Context) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx)
		} else {
			r1, r2, r3, err = fn(ctx)
		}
		return
	}
}

func Use04[R1, R2, R3, R4 any](fn func(context.Context) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx)
		} else {
			r1, r2, r3, r4 = fn(ctx)
		}
		return
	}
}
func UseErr04[R1, R2, R3, R4 any](fn func(context.Context) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx)
		} else {
			r1, r2, r3, r4, err = fn(ctx)
		}
		return
	}
}

func Use05[R1, R2, R3, R4, R5 any](fn func(context.Context) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx)
		}
		return
	}
}
func UseErr05[R1, R2, R3, R4, R5 any](fn func(context.Context) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx)
		}
		return
	}
}

func Use06[R1, R2, R3, R4, R5, R6 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx)
		}
		return
	}
}
func UseErr06[R1, R2, R3, R4, R5, R6 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx)
		}
		return
	}
}

func Use07[R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx)
		}
		return
	}
}
func UseErr07[R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx)
		}
		return
	}
}

func Use08[R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx)
		}
		return
	}
}
func UseErr08[R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx)
		}
		return
	}
}

func Use09[R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx)
		}
		return
	}
}
func UseErr09[R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx)
		}
		return
	}
}

func Use10[A1 any](fn func(context.Context, A1), name string, opts ...trace.SpanStartOption) func(context.Context, A1) {
	return func(ctx context.Context, a1 A1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1)
		} else {
			fn(ctx, a1)
		}
	}
}
func UseErr10[A1 any](fn func(context.Context, A1) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1) error {
	return func(ctx context.Context, a1 A1) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1)
		} else {
			err = fn(ctx, a1)
		}
		return
	}
}

func Use11[A1, R1 any](fn func(context.Context, A1) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1) R1 {
	return func(ctx context.Context, a1 A1) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1)
		} else {
			r1 = fn(ctx, a1)
		}
		return
	}
}
func UseErr11[A1, R1 any](fn func(context.Context, A1) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1)
		} else {
			r1, err = fn(ctx, a1)
		}
		return
	}
}

func Use12[A1, R1, R2 any](fn func(context.Context, A1) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1)
		} else {
			r1, r2 = fn(ctx, a1)
		}
		return
	}
}
func UseErr12[A1, R1, R2 any](fn func(context.Context, A1) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1)
		} else {
			r1, r2, err = fn(ctx, a1)
		}
		return
	}
}

func Use13[A1, R1, R2, R3 any](fn func(context.Context, A1) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1)
		} else {
			r1, r2, r3 = fn(ctx, a1)
		}
		return
	}
}
func UseErr13[A1, R1, R2, R3 any](fn func(context.Context, A1) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1)
		} else {
			r1, r2, r3, err = fn(ctx, a1)
		}
		return
	}
}

func Use14[A1, R1, R2, R3, R4 any](fn func(context.Context, A1) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1)
		}
		return
	}
}
func UseErr14[A1, R1, R2, R3, R4 any](fn func(context.Context, A1) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1)
		}
		return
	}
}

func Use15[A1, R1, R2, R3, R4, R5 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1)
		}
		return
	}
}
func UseErr15[A1, R1, R2, R3, R4, R5 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1)
		}
		return
	}
}

func Use16[A1, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1)
		}
		return
	}
}
func UseErr16[A1, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1)
		}
		return
	}
}

func Use17[A1, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1)
		}
		return
	}
}
func UseErr17[A1, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1)
		}
		return
	}
}

func Use18[A1, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1)
		}
		return
	}
}
func UseErr18[A1, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1)
		}
		return
	}
}

func Use19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1)
		}
		return
	}
}
func UseErr19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1)
		}
		return
	}
}

func Use20[A1, A2 any](fn func(context.Context, A1, A2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) {
	return func(ctx context.Context, a1 A1, a2 A2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2)
		} else {
			fn(ctx, a1, a2)
		}
	}
}
func UseErr20[A1, A2 any](fn func(context.Context, A1, A2) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) error {

	return func(ctx context.Context, a1 A1, a2 A2) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2)
		} else {
			err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use21[A1, A2, R1 any](fn func(context.Context, A1, A2) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) R1 {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2)
		} else {
			r1 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr21[A1, A2, R1 any](fn func(context.Context, A1, A2) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2)
		} else {
			r1, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use22[A1, A2, R1, R2 any](fn func(context.Context, A1, A2) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2)
		} else {
			r1, r2 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr22[A1, A2, R1, R2 any](fn func(context.Context, A1, A2) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2)
		} else {
			r1, r2, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use23[A1, A2, R1, R2, R3 any](fn func(context.Context, A1, A2) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr23[A1, A2, R1, R2, R3 any](fn func(context.Context, A1, A2) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use24[A1, A2, R1, R2, R3, R4 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr24[A1, A2, R1, R2, R3, R4 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use25[A1, A2, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr25[A1, A2, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use26[A1, A2, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr26[A1, A2, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use27[A1, A2, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr27[A1, A2, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2)
		}
		return
	}
}
func UseErr29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2)
		}
		return
	}
}

func Use30[A1, A2, A3 any](fn func(context.Context, A1, A2, A3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3)
		} else {
			fn(ctx, a1, a2, a3)
		}
	}
}
func UseErr30[A1, A2, A3 any](fn func(context.Context, A1, A2, A3) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3)
		} else {
			err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use31[A1, A2, A3, R1 any](fn func(context.Context, A1, A2, A3) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3)
		} else {
			r1 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr31[A1, A2, A3, R1 any](fn func(context.Context, A1, A2, A3) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3)
		} else {
			r1, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use32[A1, A2, A3, R1, R2 any](fn func(context.Context, A1, A2, A3) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr32[A1, A2, A3, R1, R2 any](fn func(context.Context, A1, A2, A3) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use33[A1, A2, A3, R1, R2, R3 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr33[A1, A2, A3, R1, R2, R3 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use34[A1, A2, A3, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr34[A1, A2, A3, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use35[A1, A2, A3, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr35[A1, A2, A3, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use36[A1, A2, A3, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr36[A1, A2, A3, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use37[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr37[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use38[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr38[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use39[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3)
		}
		return
	}
}
func UseErr39[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

func Use40[A1, A2, A3, A4 any](fn func(context.Context, A1, A2, A3, A4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4)
		} else {
			fn(ctx, a1, a2, a3, a4)
		}
	}
}
func UseErr40[A1, A2, A3, A4 any](fn func(context.Context, A1, A2, A3, A4) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4)
		} else {
			err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use41[A1, A2, A3, A4, R1 any](fn func(context.Context, A1, A2, A3, A4) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr41[A1, A2, A3, A4, R1 any](fn func(context.Context, A1, A2, A3, A4) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use42[A1, A2, A3, A4, R1, R2 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr42[A1, A2, A3, A4, R1, R2 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use43[A1, A2, A3, A4, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr43[A1, A2, A3, A4, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use44[A1, A2, A3, A4, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr44[A1, A2, A3, A4, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use45[A1, A2, A3, A4, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr45[A1, A2, A3, A4, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use46[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr46[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use47[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr47[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use48[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr48[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use49[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}
func UseErr49[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

func Use50[A1, A2, A3, A4, A5 any](fn func(context.Context, A1, A2, A3, A4, A5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5)
		} else {
			fn(ctx, a1, a2, a3, a4, a5)
		}
	}
}
func UseErr50[A1, A2, A3, A4, A5 any](fn func(context.Context, A1, A2, A3, A4, A5) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use51[A1, A2, A3, A4, A5, R1 any](fn func(context.Context, A1, A2, A3, A4, A5) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr51[A1, A2, A3, A4, A5, R1 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use52[A1, A2, A3, A4, A5, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr52[A1, A2, A3, A4, A5, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use53[A1, A2, A3, A4, A5, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr53[A1, A2, A3, A4, A5, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use54[A1, A2, A3, A4, A5, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr54[A1, A2, A3, A4, A5, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use55[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr55[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use56[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr56[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use57[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr57[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use58[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr58[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use59[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}
func UseErr59[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

func Use60[A1, A2, A3, A4, A5, A6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6)
		}
	}
}
func UseErr60[A1, A2, A3, A4, A5, A6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use61[A1, A2, A3, A4, A5, A6, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr61[A1, A2, A3, A4, A5, A6, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use62[A1, A2, A3, A4, A5, A6, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr62[A1, A2, A3, A4, A5, A6, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use63[A1, A2, A3, A4, A5, A6, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr63[A1, A2, A3, A4, A5, A6, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use64[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr64[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use65[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr65[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use66[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr66[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use67[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr67[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use68[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr68[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use69[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}
func UseErr69[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

func Use70[A1, A2, A3, A4, A5, A6, A7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
	}
}
func UseErr70[A1, A2, A3, A4, A5, A6, A7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use71[A1, A2, A3, A4, A5, A6, A7, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr71[A1, A2, A3, A4, A5, A6, A7, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use72[A1, A2, A3, A4, A5, A6, A7, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr72[A1, A2, A3, A4, A5, A6, A7, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use73[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr73[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use74[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr74[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use75[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr75[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use76[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr76[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use77[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr77[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use78[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr78[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use79[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}
func UseErr79[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

func Use80[A1, A2, A3, A4, A5, A6, A7, A8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
	}
}
func UseErr80[A1, A2, A3, A4, A5, A6, A7, A8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use81[A1, A2, A3, A4, A5, A6, A7, A8, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr81[A1, A2, A3, A4, A5, A6, A7, A8, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use82[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr82[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use83[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr83[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use84[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr84[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use85[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr85[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use86[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr86[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use87[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr87[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use88[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr88[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use89[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}
func UseErr89[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

func Use90[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
	}
}
func UseErr90[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) error, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use91[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) R1, name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr91[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use92[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr92[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use93[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr93[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use94[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr94[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use95[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr95[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use96[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr96[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use97[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr97[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use98[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr98[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

func Use99[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
func UseErr99[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), name string, opts ...trace.SpanStartOption) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := ByContext(ctx, opts...); t != nil {
			cx, span := t.StartSpan(name, ctx)
			defer func() {
				defer span.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = fmt.Errorf("%s", x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
