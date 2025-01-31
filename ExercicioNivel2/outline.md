## Fundamentos da Programação

### Tipo booleano

- Agora vamos explorar os tipos de maneira mais detalhada.
- O tipo bool é um tipo binário, que só pode conter um dos dois valores: true e false. (Verdadeiro ou falso, sim ou não, zero ou um, etc.)
- Sempre que você ver operadores relacionais (==, <=, >=, !=, <, >), o resultado da expressão será um valor booleano.
- Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.

### Como os computadores funcionam

- Isso é importante pois daqui pra frente vamos falar de ints, bytes, e etc.
- Não é necessário um conhecimento a fundo mas é importante ter uma idéia de como as coisas funcionam por trás dos panos.
- ASCII: <https://en.wikipedia.org/wiki/ASCII>

### Tipos numéricos

- int vs. float: Números inteiros vs. números com frações.
- Integers:
    - Números inteiros
    - int & uint → “implementation-specific sizes”
    - Todos os tipos numéricos são distintos, exceto:
        - byte = uint8
        - rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
    - Tipos são únicos
        - Go é uma linguagem estática
        - int e int32 não são a mesma coisa
        - Para "misturá-los" é necessário conversão
    - Regra geral: use somente int
- Floating point:
    - Números racionais ou reais
    - Regra geral: use somente float64

### Tipo string (cadeias de caracteres)

- Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
    - %v(qualquer valorem um formato natural) %T(tipo de qualquer valor)
    - Raw string literals => são strings apresentadas em sua forma bruta sem formatação, para isso colocamaos entre clases `` .
    - Conversão para slice of bytes: []byte(x)
    - %#U(Unicode format), %#x(hexadecimal)<https://pkg.go.dev/fmt>
    - <https://blog.golang.org/strings>

### Sistemas numéricos

- Base-10: decimal, 0–9
- Base-2: binário, 0–1
- Base-16: hexadecimal, 0–f


### Constantes

- São valores imutáveis.
- Podem ser tipadas ou não:
    - const hello = "Bom dia"
    - const hello string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
    - Ex. qual o tipo de 42? int? uint? float64?
    - Ou seja, é uma flexibilidade conveniente.

### Iota

- É um identificador especial pré-declarado integrado que simplifica a definição de constantes de inclemento, ela contém um valor inteiro.
- É normalmente usada dentro de declarações constantes para gerar uma série de valores relacionados, incrementando em 1 para cada constante subsequente.Ex:
```
  const (a = iota
         b = iota
         c = iota
  )
```
```
  func main(){
      fmt.Println(a, b, c)
  }
```
- Na prática.
    - iota,
    - iota + 1,
    - a = iota b c
    - reinicia em cada const
    -  _ pode ser usado para ignorar uma variável


### Deslocamento de bits

- O deslocamento de bits em Go é a operação de mover dígitos binários para a direita ou para a esquerda.
- Na prática:
 ```
  func main(){
       a := 10     // 1010 em binário
       result := a << 1  // 10100 em binário ou 20 em decimal
       fmt.Println("a << 1:", result) //Output: 20
  }
```
-  Explicação:
```
        1010  (a em binário)
        << 1  (desloca todos os bits para a esquerda)
        ----
        10100  (resultado, equivalente a 20 em decimal)
```
     
- Saiba mais:
- <https://splice.com/blog/iota-elegant-constants-golang/>
- <https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827>

- Fim da sessão!
