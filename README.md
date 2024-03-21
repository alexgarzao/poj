# POJ (Pascal running on the JVM)

O objetivo do projeto é rodar um _subset_ de [Pascal](https://en.wikipedia.org/wiki/Pascal_(programming_language)) na [JVM](https://en.wikipedia.org/wiki/Java_virtual_machine). Para tal, iremos criar um [compilador](https://en.wikipedia.org/wiki/Compiler) que entende este _subset_ de Pascal e gere o [_assembly_ Java](https://www.eg.bucknell.edu/~cs360/java-assembler/examples.html) equivalente. De posse deste _assembly_, será utilizado um [montador _assembly_](https://pt.wikipedia.org/wiki/Linguagem_assembly#Montador) Java para gerar o _bytecode_ em um arquivo [class](https://en.wikipedia.org/wiki/Java_class_file), possibilitando assim que este arquivo seja executado na JVM.

Qual a motivação para o projeto? Aprender temas relacionados a compiladores. E, para tal, nada melhor do que criar um, por mais simples que seja.

POJ é um projeto com escopo simples (na medida do possível) e de fácil entendimento. O foco é ser utilizado para estudos.

# Requisitos

- Git
- Go 1.22
- Make
- ANTLR 4.13.1: pacote está versionado junto com este projeto na pasta parser
- JAVA: para poder executar o ANTLR e para executar o arquivo class (bytecode java) gerado
- JASM 0.7.0: instruções sobre como baixar [aqui](https://github.com/roscopeco/jasm)

# Como gerar o executável do POJ

```
# Baixar o repositório do projeto
git clone git@github.com:alexgarzao/poj.git

# Executar os testes (opcional)
make test

# Construir o binário
make build
```

Após isso, o binário do POJ estará em na pasta bin.

# Passos para compilar o hello_world.pas

Segue abaixo o passo-a-passo para compilar o "Hello world!".

```
# Executar o POJ para gerar o assembly Java (arquivo jasm)
./bin/poj ./tests/pascal_programs/hello_world

# Executar o JASM (java assembler) para gerar o executável Java (arquivo class)
jasm hello_world.jasm

# Executar o arquivo class com a JVM
java hello_world
```

Uma forma mais enxuta é utilizar o make:

```
make compile-and-run-example program=hello_world
```

# Exemplos de programas aceitos

Abaixo é possível ver o clássico “Hello world!” em Pascal:

```
program Hello;
begin
  writeln ('Hello world!');
end.
```

Abaixo temos o cálculo do fatorial, de forma recursiva:

```
program fatorial;

var numero : integer;

function fatorial(n : integer) : integer;
begin
    if n<0 then fatorial := 0
    else begin
        if n<=1 then fatorial := 1
        else fatorial := n * fatorial(n-1);
    end;
end;

begin
    write('Introduza numero inteiro: ');
    readln(numero);
    writeln;
    writeln('O fatorial de ', numero, ' e: ', fatorial(numero));
end.
```

Abaixo temos um exemplo com entrada e saída de dados:

```
program NameAndAge;
var
  MyName: String;
  MyAge : Byte;
begin
  Write('What is your name? '); Readln(MyName);
  Write('How old are you? '); Readln(MyAge);
  Writeln;
  Writeln('Hello ', MyName);
  Writeln('You are ', MyAge, ' years old');
end.
```

No momento, apenas o "Hello world!" pode ser compilado com o POJ. Para quem tiver interesse de ver a execução destes outros programas, sugiro utilizar o [Free Pascal Compiler](https://www.freepascal.org/). Nas próximas semanas poderemos utilizar o POJ com todos estes exemplos :-)

# Ajustes e melhorias

O projeto ainda está em desenvolvimento e as próximas atualizações serão voltadas nas seguintes tarefas:

- [x] Definir qual gerador de parser será utilizado
- [x] Encontrar uma gramática de Pascal pronta no formato do gerador de parsers (ANTLR)
- [x] Ajustar a gramática para reconhecer o subset de Pascal esperado
- [x] No parser identificar a definição de procedimentos Pascal para determinar o bloco Pascal principal
- [x] Encontrar um montador Java assembly (JASM)
- [x] POJ gerar um código assembly Java válido para o assemblador
- [x] Criar o README inicial
- [ ] Declaração e uso de variáveis string
- [ ] Saída de dados (terminal)
- [ ] Instrução If/Else
- [ ] Declaração e uso de variáveis inteiras
- [ ] Operações aritméticas sem precedência de operadores
- [ ] Instrução For
- [ ] Instrução Repeat
- [ ] Instrução While
- [ ] Operações aritméticas com precedência de operadores
- [ ] Entrada de dados (terminal)
- [ ] Declaração e uso de variáveis booleanas
- [ ] Declaração e uso de variáveis de ponto flutuante (Real)
- [ ] Declaração e uso de procedures
- [ ] Declaração e uso de funções
- [ ] Uso de funções recursivas

# Licença

POJ utiliza a [licença Apache](LICENSE). 
