package main


import (
	"fmt"
)

type Produto struct {
	Nome string
	Preco float64
	Quantidade int
}

type Carrinho struct {
	Itens []Produto
}


func adicionarCarrinho(c *Carrinho, produto Produto) {
	nc := make([]Produto, len(c.Itens))
	copy(nc, c.Itens)
	
	encontrado := false
	for i, item := range c.Itens {
		if item.Nome == produto.Nome {
			nc[i].Quantidade += produto.Quantidade
			encontrado = true
			break
		}
	}

	if !encontrado {
		nc  = append(nc, produto)
	}
	c.Itens = nc
}

func calcularValorTotal(c Carrinho) float64 {
	var total float64
	for _, item := range c.Itens {
		total += item.Preco * float64(item.Quantidade)
	}
	return total
}

func criarProduto(prods *[]Produto) {

	var nome string
	fmt.Println("Digite o nome do produto")
	fmt.Scanln(&nome)

	var preco float64
	fmt.Println("Digite o preço do produto")
	fmt.Scanln(&preco)

	var quantidade int
	fmt.Println("Digite a quantadade")
	fmt.Scanln(&quantidade)

	var produto Produto = Produto{
		Nome: nome,
		Preco: preco,
		Quantidade: quantidade,
	}

	*prods = append(*prods, produto)
}

func listarProdutos(prods []Produto) {
	for i, prod := range prods {
		fmt.Printf("%d - %s | Preço: R$ %.2f | Quantidade: %d\n", i, prod.Nome, prod.Preco, prod.Quantidade)
	}
}

func listarCarrinho(c *Carrinho) {
	for i, prod := range c.Itens {
		fmt.Printf("%d - %s | Preço: R$ %.2f | Quantidade: %d\n", i, prod.Nome, prod.Preco, prod.Quantidade)
		fmt.Println("____")
	}
	fmt.Println("")
}

func showMenu() {
	fmt.Println(`
		1 - Criar Produto
		2 - Listar Produtos
		3 - Adicionar Produto ao Carrinho
		4 - Mostrar Menu
		5 - Mostrar Carrinho
		0 - Finalizar Compra
	`)
}

func tela() {
	var carrinho Carrinho = Carrinho{}
	var produtos []Produto = []Produto{}

	showMenu();

	var option int = 99
	for option != 0 {
		fmt.Scanln(&option)
		switch option {
			case 1:
				criarProduto(&produtos)
			case 2: 
				listarProdutos(produtos)
			case 3:
				fmt.Println("Digite o indice do produto")
				var ind int
				fmt.Scanln(&ind)
				adicionarCarrinho(&carrinho, produtos[ind])
			case 4:
				showMenu()
			case 5:
				listarCarrinho(&carrinho)
			case 0:
				option = 0
		}
	}
}

func main() {
	tela()
}
