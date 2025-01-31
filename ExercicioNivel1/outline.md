### Hello world!

- Estrutura básica: 
    - package => em Go, packages são uma forma de agrupar arquivos qu contenham código de um mesmo escopo, Ex: arquivos como structs, funções, variáveis, etc..
    - package main => pacote principal
    - func main: é onde a execução começa e termina, só pode haver um por pasta.
    - import => são os pacotes nativos de go que importamos para usar nos nossos projetos. Ex: import "fmt"
- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Go não aceita variáveis não utilizadas, se declarou tem que usar ou vai dar erro.

### Operador curto de declaração

- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas 
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks(blocos de código)
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores => Operadores em programação são símbolos que representam ações específicas que podem ser executadas em expressões. Eles são usados para 
    incrementar, decrementar, comparar e avaliar dados.Os operadores podem ser aritméticos, lógicos ou relacionais.(+  -  *  /)
    - operandos => são os valores sobre os quais os operadores atuam. Os operandos podem ser: Constantes, Variáveis, Resultados de funções. 
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões 
    - expressão -> qualquer coisa que "produz um resultado" Ex: 2 + 2
    - scope (abrangência) => o escopo de uma variável é definido pela sua posição no código-fonte. A palavra "escopo" refere-se à parte do programa em que a 
     variável pode ser utilizada. 
       - package-level scope => escopo global
        - Variáveis globais
          - Variáveis declaradas fora de qualquer função
          - Podem ser acessadas por qualquer função no pacote
        - Variáveis locais 
          - São definidas dentro de uma função
          - Os parâmetros formais de uma função também são tratados como variáveis locais
       -Escopos de variáveis em Go
          - O escopo de dentro pode acessar o de fora, mas não o contrário 
          - A variável é "visível" dentro do seu escopo 
       - Lição principal:
          -  := utilizado pra criar novas variáveis, dentro de code blocks
          -  = para atribuir valores a variáveis já existentes

### A palavra-chave var

- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar

### Explorando tipos

- Tipos em Go são extremamente importantes.
- Tipos em Go são estáticos.
- Ao declarar uma variável para conter valores de um certo tipo, essa variável só poderá conter valores desse tipo.
- O tipo pode ser deduzido pelo compilador:
    - x := 10
    - var y = "Eu amo Coca cola"
- Ou pode ser declarado especificamente:
    - var w string = "isso é uma string"
    - var z int = 15
    - na declaração var z int com package scope, atribuição z = 15 no codeblock (somente)
- Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
    - int, string, bool
- Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
    - slice, array, struct, map
- O ato de definir, criar, estruturar tipos compostos chama-se composição.

### Valor zero

- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero?
- Os zeros:
    - ints: 0
    - floats: 0.0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.

### O pacote fmt

- Setup: strings, ints, bools.
- Strings: interpreted string literals vs. raw string literals.
    - Rune literals.
    - Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte. 
- Format printing: documentação.
    - Grupo #1: Print -> standard out
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error)
        - func Printf(format string, a ...interface{}) (n int, err error)
            - Format verbs. (%v %T)
    - Grupo #2: Print -> string, pode ser usado como variável
        - func Sprint(a ...interface{}) string
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string
    - Grupo #3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

### Criando seu próprio tipo

- Tipos em Go são extremamente importantes. tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
- Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
- Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
- type hotdog int → var b hotdog (main hotdog)
- Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog. 

### Conversão, não coerção

- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. 
