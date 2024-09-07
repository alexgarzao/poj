# POJ (Pascal running on the JVM)

O objetivo do projeto é rodar um _subset_ de [Pascal](https://en.wikipedia.org/wiki/Pascal_(programming_language)) na [JVM](https://en.wikipedia.org/wiki/Java_virtual_machine). Para tal, iremos criar um [compilador](https://en.wikipedia.org/wiki/Compiler) que entende este _subset_ de Pascal e gere o [_assembly_ Java](https://www.eg.bucknell.edu/~cs360/java-assembler/examples.html) equivalente. De posse deste _assembly_, será utilizado um [montador _assembly_](https://pt.wikipedia.org/wiki/Linguagem_assembly#Montador) Java para gerar o _bytecode_ em um arquivo [class](https://en.wikipedia.org/wiki/Java_class_file), possibilitando assim que este arquivo seja executado na JVM.

Qual a motivação para o projeto? Aprender temas relacionados a compiladores. E, para tal, nada melhor do que criar um, por mais simples que seja.

POJ é um projeto com escopo simples (na medida do possível) e de fácil entendimento. O foco é ser utilizado para estudos.

# Requisitos

- Git
- Go 1.22
- Make
- ANTLR 4.13.1: Gerador de parsers (pacote está versionado junto com este projeto na pasta parser)
- JAVA: Runtime do Java (necessário para executar o ANTLR e para executar o arquivo class (bytecode java) gerado)
- JASM 0.7.0: Assemblador Java assembly (necessário para gerar o arquivo class (instruções sobre como baixar [aqui](https://github.com/roscopeco/jasm)))

# Como gerar o executável do POJ

```
# Baixar o repositório do projeto
git clone git@github.com:alexgarzao/poj.git

# Executar os testes (opcional)
make test

# Construir o binário
make build
```

Após isso, o binário do POJ estará na pasta bin.

# Passos para compilar o hello_world.pas

Segue abaixo o passo-a-passo para compilar o "Hello world!".

```
# Executar o POJ para gerar o assembly Java (arquivo jasm)
./bin/poj ./tests/valid_pascal_programs/hello_world

# Executar o JASM (java assembler) para gerar o executável Java (arquivo class)
jasm hello_world.jasm

# Executar o arquivo class com a JVM
java hello_world
```

Uma forma mais enxuta é utilizar o make:

```
make compile-and-run-example program=hello_world
```

# Exemplos de programas aceitos pelo POJ

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


# Recursos implementados do Pascal

O POJ implementa um subconjunto das funcionalidades existentes em compiladores de Pascal modernos. Este subconjunto compreende:

- Declaração e uso de variáveis dos tipos string, int e bool
- Entrada e saída de dados (terminal)
- Instruções de controle: If/Else, For, Repeat e While
- Operações aritméticas com precedência de operadores
- Declaração e uso de procedures e funções
- Uso de funções recursivas

# Documentação do projeto

Maiores informações sobre o funcionamento e a arquitetura do projeto podem ser vistas nesta série de publicações [aqui](https://dev.to/alexgarzao/series/26440). Além disso, temos a [visão macro](docs/visao_macro.md) bem como as ["internals"](docs/internals.md) do projeto.

# Licença

POJ utiliza a [licença Apache](LICENSE). 
