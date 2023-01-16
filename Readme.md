## Setup
- Máquina: Macbook Pro
- Chip: M1 Max
- Ram: 32gb
- Arch: darwin/arm64
- MacOs 13
- Go: 1.19 version


## Run
- Install Golang 1.19
- Configure input.json
```
go run main.go
```


## How it works
Cria uma simulação de como funciona em memória e como os processos interagem com a memória virtual. Quando acontece uma falta de página, ele executa o algoritmo Working Set para fazer a substituição. A memória é representada como uma sequência de páginas, onde cada página possui uma identificação, uma estrutura de dados, um bit de referência e o seu último acesso. A memória virtual é onde será mapeado cada processo para a memória principal por meio de mapas. O algoritmo do Working Set implementa a troca de páginas, verificando a janela de cada processo, caso a página não estiver dentro da janela, ele retornará o índice dessa página e a substituição será iniciada, caso contrário, ele retornará a pagina mais antiga em memória.