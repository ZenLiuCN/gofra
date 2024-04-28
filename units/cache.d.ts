export interface Cache<K, V> {
    emptyValue(): V

    timeToLive(): time.Duration

    measureUnit(): Measure

    put(k: K, v: V)

    putTTL(k: K, v: V, ttl: time.Duration)

    /**
     * @return [value,exists]
     */
    get(k: K): (V | boolean)[]

    invalidate(k: K)

    invalidateAll()

    all(): Entry<K, V>[]
    count():number
    purify()

    close()
}

export interface Entry<K, V> {
    key(): K

    data(): V

    goString(): string
}

//0 nanos 1 micros 2 mills 3 seconds
export type Measure = 0 | 1 | 2 | 3

export interface Option {
}

// @ts-ignore
import * as time from "go/time"

export function withMaxSize(n: number): Option

export function withExpireAfterAccess(n: time.Duration): Option

export function newStringKeyCache(freq: time.Duration, ttl: time.Duration, unit: Measure, ...opts: Option[]): Cache<string, any>

export function newNumberKeyCache(freq: time.Duration, ttl: time.Duration, unit: Measure, ...opts: Option[]): Cache<number, any>
