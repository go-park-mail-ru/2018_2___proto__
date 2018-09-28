package main

//просто проверяю, что интерфейс полностью реализован, потом уберу
func test(ctx IContext) {

}

func main() {
	c := NewContext(nil, nil)
	test(c)
}
