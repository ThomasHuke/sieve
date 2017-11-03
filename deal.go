//Package sieve主要是处理
package sieve

type Sie interface {
	DealWith(ty int, src string) (string, error)
}
