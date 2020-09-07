# Playlist Youtube  
Go – Aprenda a Programar (Curso)  
https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg  

>> Capítulos  
	- 1  
	- 2  
	- 3  

# Exercícios realizados  
	1) https://play.golang.org/p/2BupCtzk2E4  
	2) https://play.golang.org/p/LOip7_w9ytn  
	3) https://play.golang.org/p/kFZurcbreW5  
	4) https://play.golang.org/p/W53m7DSZld3  
	5) https://play.golang.org/p/Kdj3xVYLITF  
	6) Prova:  
		a. https://docs.google.com/forms/u/0/d/e/1FAIpQLSchZktObQvMOW0HThOow4alvRYtkSAcpjLR2urxz8gZoxaXjw/formResponse  

>> Capítulos  
	- 4 (Fundamentos da programação)  
	- 5 (Exercícios)  

   Geração dos computadores iniciou com o ENIAC como a primeira geração criada nos anos 40. Composto com válvulas termiônicas e depois evoluído para transistores.  
		○ Obs.: As válvulas tinham como condutores os gases de depois os transistores utilizaram de condutores elétricos.  
Para entendimento mais aprofundado sobre a história dos computadores e sua evolução ver o filme Alan Turing (The Immitation Game).  

Apresentação complementar:  
https://docs.google.com/presentation/d/1aVytiGOBVDMISFW-ZARJ5iFY1osU2XJIw0hQpNICXm8/edit#slide=id.g1e92cb3441_0_264  


Tipos em GO:  
# Strings  
https://blog.golang.org/strings  

# Sistemas numéricos  
https://docs.google.com/document/d/1GqXpubhMMIr4Sy5xwgiPIDh5PGVmVpF2u0c9vDrvykE/  

# IOTA - Go Lang  
https://golang.org/ref/spec  
Toddy - Resolução com Iota e visualizando o passeio dos bits para os lados e mostrando seu resultado em decimal  
https://play.golang.org/p/7MOnbhx4R4  

Conceitos  
https://splice.com/blog/iota-elegant-constants-golang/  
https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827  

# Exercícios realizados  
	1) https://play.golang.org/p/39-fx9ofpeE  
	2) https://play.golang.org/p/8g-1ubhUWeS  
	3) https://play.golang.org/p/LZSCArHIeVk  
	4) https://play.golang.org/p/HhLnFgt_Qvs  
	5) https://play.golang.org/p/twogk7PJrX2  
	6) https://play.golang.org/p/6LHX24QuypN  
	7) Prova:  
		a. CleanForm: https://docs.google.com/forms/d/e/1FAIpQLSchSZmzSvOu8lC9hfgLGkxsjEbW9pN5B10LBWJzJ676u0KVSA/viewform  
		b. ResponseForm: https://docs.google.com/forms/d/e/1FAIpQLSchSZmzSvOu8lC9hfgLGkxsjEbW9pN5B10LBWJzJ676u0KVSA/viewscore?viewscore=AE0zAgBF8PTkE6a6TuU0UvV7IHzfEc8kYVJSnh5ddJyUU5ZRW5pican8iRRHcsQgmtp6P8g  

>> Capítulos  
	- 6  
	- 7  

# Anotações:  
Estrutura de controle de repetição  
https://gobyexample.com/for  
https://golang.org/ref/spec#For_statements  

Desafio surpresa  
https://play.golang.org/p/Y0PZrwRU13A  

# Exercícios realizados  
	1) https://play.golang.org/p/2owMVaqMIcg  
	2) https://play.golang.org/p/1jiKBnP_RR5 - string  
	https://play.golang.org/p/gL6iXZWkLNv - Unicode point  
	https://play.golang.org/p/Ju9bfX8ChNg - Unicode point 3x  
	3) https://play.golang.org/p/Tr7hI3ih_tn  
	4) https://play.golang.org/p/zO7FwxyiTYP  
	5) https://play.golang.org/p/jZv8gvEiwX5  
	6) https://play.golang.org/p/EK8Ljk6hJ4J  
	7) https://play.golang.org/p/EK8Ljk6hJ4J  
	8) https://play.golang.org/p/VK4RPghvXCQ  
	9) https://play.golang.org/p/Q3p_lNFX8W5  
	10) https://play.golang.org/p/4qVqVHPBANL  

>> Capítulos - Agrupamento de dados.  
	- 8  
	- 9  

# Anotações  
	- Arrays:  
		○ Em GO utilizamos o slice para representar o "Array". O slice tem como elemento "subjacente" o Array nativo. Porém, o Slice ele não tem um limite prédefinido de elementos. Já o Array tem um tamanho finito quando é declarado.  
			§ Ex.: slice := []int{1, 2, 3, 4, 5} vs array := [5]{1, 2, 3, 4, 5}  
		○ Para deletar o valor de uma slice é fatiando a fatia.  
			§ Ex.: {bolo, café, chá, pao}, se eu quiser tirar o pao do cardápio eu sobrescrevo os elementos com append selecionando a faixa de [0:3] ficando {bolo, café, chá}. O 3 não é incluso.  
		○ comma ok idiom, ex.: example, ok := example["joao"]; if example, ok; !ok ; { fmt.Println("not exist")}  
		○ map : ex.:  
		```  
		example := map[string]int {
						"Joao": 35,
						"Jose": 45,
					      }
		```  
		○ Para adicionar elementos em map apenas atribua. Ex.: example["Maria"] = 58
		○ Podemos o range com map para percorrer o key, value pairs. Ex.: for key, value := range example {fmt.Println(key, value)}, blablabla
		○ Podemos deletar um item, ex.: delete(example, "Joao')

# Exercícios Realizados  
	1) https://play.golang.org/p/VgX53rV5cJN  
	2) https://play.golang.org/p/AaHNGHN_3lY  
	3) https://play.golang.org/p/MwAsZmfCIkL  
	4) https://play.golang.org/p/vbai2XlaYh6  
	5) https://play.golang.org/p/nmnIt0aoxoc  
	6) https://play.golang.org/p/r0vuPZDDqft  
	https://play.golang.org/p/qmsN2tkMSTw - com for  
	7) https://play.golang.org/p/onZAKbgcGCF  
	8) https://play.golang.org/p/_Js0ZGaC4pW  
	9) https://play.golang.org/p/BpSpJ7rmuU9  
	10) https://play.golang.org/p/7-q8na1CJOT - deletando um elemento do map  


>> Capítulos - Structs  
	- 10  
	- 11  

# Anotações  
Structs são um tipo de estruturas de dados. Então serve para definir uma estrutura com vários tipos diferentes de dados, se for o caso.  
Ex.:  
struct  
------------------------------------------------------------------------  
```
	type cliente struct {
		nome string,
		cpf int,
		maiordezoitoanos bool,
	}
		
	primeiroCliente = cliente{"Gabriel", 9999999999, true}
	segundoCliente = cliente{
		nome: "Malutrom", 
		cpf: 9999999999,
		maiordezoitoanos: true
	}
	fmt.Println(primeiroCliente)  
	fmt.Println(segundoCliente)  
```
------------------------------------------------------------------------  


# Exercícios realizados  
	1) https://play.golang.org/p/RUT4vTiJO3P  
	2) https://play.golang.org/p/KhcUkksLb6p - Inacabado... Continuar  
	https://play.golang.org/p/lndO2sFQ1n3 - Mais coerente  
	3) https://play.golang.org/p/162Z7VQsusX  
	4) https://play.golang.org/p/lSRILesNrHf  
	
	

>> Capítulos - Funções  
	- 12  
	- 13  


# Anotações  
Ver mais sobre:  
	funções anônimas: Funções auto contidas, normalmente já passa um valor no ato da função.  
	closure: Declarações fora do escopo, usando o valor fora do escopo dentro de um return de uma outra função 
	callbak: Função que recebe um retorno de uma outra função.  
Retornar aos vídeos sobre esses assuntos.  

# Exercícios realizados  
	1) https://play.golang.org/p/0ZAc_PJxRZ1  
	2) https://play.golang.org/p/05p5w0Imxc4  
	3) https://play.golang.org/p/5QfihiuMD5m  
	4) https://play.golang.org/p/_5X0gvrqY_p  
	5) https://play.golang.org/p/TM9gEWYlps7  
	6) https://play.golang.org/p/fPiYG2ANkyQ  
	7) https://play.golang.org/p/nrrNiffcWl7  
	8) https://play.golang.org/p/tGrUC4lY5MS  
	9) https://play.golang.org/p/TudU1-96jpf  
	10) https://play.golang.org/p/HXd3QWyeEV8  
	11) Tenho que fazer um vídeo sobre funções...  


Capítulos - Ponteiros  
	- 14  
	- 15  

# Anotações  
Direference: & <- para saber o endereço na memória e * para saber o valor representado de uma dado endereço da memória (&).  
Exemplo.: x := 10  
		  y := &x <- endereço na memória  
		  *y <- extrai o valor do endereço da memória.  
O Ponteiro é uma variável que armazena o endereço na memória.  
Existe uma economia de performance quando vc altera o valor diretamente no endereço na memória. Ao invés de fazer um cópia do
valor.  

# Exercícios realizados  
	1) https://play.golang.org/p/sdAqqDW1j4s  
	   https://play.golang.org/p/DUWBGM0smLi  
	2) https://play.golang.org/p/BJzqJ52q9ZC  


Capítulos - Aplicações  
	- 16  
	- 17  

# Anotações  

golang.org/pkg/  
godoc.org  

Método chaining é um conceito de ligar um médodo no outro. Ou seja, você pode criar uma instância de um objeto para acessar seus métodos.  
No entanto, ao invés de criar a instância. Você pode acessar o método diretamente.  
Exemplos.:  
# Acessando diretamente:  
	json.NewEncoder(os.Stdout).Encode(<entrada>)  

# Criando instância:  
	myinstance := json.NewEncoder(os.Stdout)  
	myinstance.Encode(<entrada>)  	

testes por minha conta:  
https://play.golang.com/p/-aZJWUoHLa0  

# Exercícios realizados  
	1) https://play.golang.org/p/jFFFzmhModr  
	   https://play.golang.org/p/hZs-Qb0vuDH <- Ellen Korbes  
	2) https://play.golang.org/p/svs900DUWO6  
	3) https://play.golang.org/p/oOpD_H7Ly4w <- sem method chaning  
	   https://play.golang.org/p/5zdi3rqQ3MG <- com method chaning  
	4) https://play.golang.org/p/Ig3g-RPkLtp <- sort(Ints|Strings)  
	5) https://play.golang.org/p/G_4QFTfaxVV <- estou aqui.. Incompleto.  
	   https://play.golang.org/p/iADkwOYU1JI <- Done!  

	
Capítulos - Concorrência e Ambiente de desenvolvimento  
	- 18  
	- 19  
	- 20  

# Anotações - Concorrência e Ambiente de desenvolvimento  

. Concorrência não tem relação com paralelismo. São coisas diferentes.  
. A concorrência é o design, a forma, o conceito, a forma de pensar, no código que será escrito. Dito isso, a gente escreve código para 
  agir de forma concorrente. Então, faz todo sentido instruir o seu programa através de funções agir de forma concorrente.  
. O paralelismo surgirá dependendo da arquitetura de processamento do computador. Ou seja, é algo encapsulado. Que nós não temos "acesso"
  o runtime da linguagem se encarrega de pegar suas funções escritas de forma concorrente "executá-las" paralelamente dependendo dos núcleos
  de processadores.  

Aqui entra o conceito das goroutines. Pode-se entender goroutines como uma espécie de threads em seu conceito. No entanto, existe uma
diferença em sua execução das goroutines pelo runtime do golang.

>>threads
https://pt.wikipedia.org/wiki/Thread_(computa%C3%A7%C3%A3o)  


Apesar de podermos utilizar quantas goroutines eu quiser e deixar que o runtime do GO assuma a responsabilidade desta gerência. Eu preciso
tornar a execução de forma "syncrona". Ou seja, enviar as goroutines para execução e ter uma forma de esperar pelas suas tarefas serem
concluídas.  
Exemplo.: Tenho várias funções que precisam ser executadas de forma concorrente. Nesse caso, aciono o go func para todas elas e imponho uma
forma de tratar com sync.WaitGroup. `Ou seja, envia todos os foguetes, não sei o que eles estão fazendo. Mais espere pela conclusão de suas
tarefas.``  

>>Data races - Condição de corrida ou concorrência.  

É preciso definir um mutex ou exclusão mútua como definição. Pois quando estamos contruindo um programa que utiliza-se de funções  
concorrentes. Espera-se que tratemos o uso de memória compartilhada, garantindo que a goroutine e/ou thread não acessem e alterem valores  
diretamente no endereço de memória de forma sincronizada. O que poderá retornar um erro inesperado. Ou seja, Condição de corrida.  

Para tratar essa condição normalmente utiliza-se o mutex. Conforme falado acima. Ele funciona como uma espécie de lock obrigando a próxima
thread aguardar na fila.  

# Anotações - Ambiente de desenvolvimento

	. Compilação cruzada.
		Em GO podemos compilar o programa para N sistemas operacionais. Em Go isso é possível.
		Utilizamos:
			- GOOS
			- GOARCH

	. packages.
		Organizando seu programa em packages.  
		Artigo: https://rakyll.org/style-packages/  
		- Podemos importar packages ou referenciá-los em outros package para usar as funções.  
		  Isso é muito útil para legebilidade do código, visão limpa e organização. Além do conceito DRY para reuso de código.  


# Exercícios realizados  
	1) https://147896@github.com/147896/estudando-golang.git:src/20_exercicios-ninja-9/01/main.go  
	2) https://147896@github.com/147896/estudando-golang.git:src/20_exercicios-ninja-9/02/main.go  
	3) https://147896@github.com/147896/estudando-golang.git:src/20_exercicios-ninja-9/03/main.go  
	4) https://147896@github.com/147896/estudando-golang.git:src/20_exercicios-ninja-9/04/main.go  
	5) https://147896@github.com/147896/estudando-golang.git:src/20_exercicios-ninja-9/05/main.go  
	6) https://147896@github.com/147896/estudando-golang.git:src/20_exercicios-ninja-9/06/main.go  

>> Minhas brincadeiras, tentando usar para um contexto prático..  
https://stackoverflow.com/questions/14668850/list-directory-in-go  <- Isso aqui é um bom começo...  

- A ideia aqui é criar uma função que liste o conteúdo de um dado diretório de forma concorrente. Ou seja, para listar um diretório eu não preciso esperar o listar de forma sincrona. Eu coloco isso dentro de uma goroutine function informando um número de wait groups que é exatamente o length da lista retornada.  
- Outra coisa, é que depois que tivermos um resultado esperado no statement acima. Como estaremos focando em um caso específico. No caso o terraform init -reconfigure e o terraform plan. Precisamos iterar sobre a lista de diretórios acima e rodar esses comandos do terraform também de forma assíncrona usando goroutines.  


Capítulos - Canais (Channel)  
	- 21  

# Anotações - ...  
	. Canais é uma forma de você trocar informações intra goroutines. É um meio de comunicação entre goroutines para possibilitar a transmissão de informações.  
	
	**Lembrando que a func main também é uma goroutine.**  

	. Canais bidirecionais (send and receiver).  
	Consultar.: https://stackoverflow.com/questions/13596186/whats-the-point-of-one-way-channels-in-go  
	send e receiver são tipos diferentes. Ou seja, existem macanismos para checar os tipos protegendo, de certa forma, a escrita em um canal que é do tipo leitura. E vice-versa.  
	<-chan == receive  
	chan<- == send  

	. Quando declaro um canal send e um receive uma delas obrigatoriamente precisa ser uma goroutine. Ou seja, precisam concorrer.  
	. Range e close em canais normalmente são usados. Servem para iterar sobre itens que irão escrever sobre o canal e o close é utilizado após a conclusão do loop para informar para o canal que não existem mais itens a serem enviados.  

	. Exemplos utilizando o https://play.golang.com  
	# iterando sobre um channel: https://play.golang.com/p/rTaHevkMXMr  
	# uso do select (similar ao switch case só que para canais): 
		1. https://play.golang.com/p/wMNdJByPBx-  
		2. https://play.golang.com/p/HRgToy4umnZ  
		3. https://play.golang.com/p/DRuYhJNLf00  <- o select pode tratar multicanais para enviar e/ou receber.
	

# Exercícios realizados  
