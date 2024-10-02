package initialize

import "github.com/hainguyen27798/open-search-engine/internal/domain/products"

func InitChangeStream() {
	go products.InitProductChangeStream()
}
