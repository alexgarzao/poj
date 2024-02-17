# POJ (Pascal running on the JVM)

O objetivo do projeto é rodar um _subset_ de [Pascal](https://en.wikipedia.org/wiki/Pascal_(programming_language)) na [JVM](https://en.wikipedia.org/wiki/Java_virtual_machine). Para tal, iremos criar um [compilador](https://en.wikipedia.org/wiki/Compiler) que entende este _subset_ de Pascal e gere o [_assembly_ Java](https://www.eg.bucknell.edu/~cs360/java-assembler/examples.html) equivalente. De posse deste _assembly_, será utilizado um [montador _assembly_](https://pt.wikipedia.org/wiki/Linguagem_assembly#Montador) Java para gerar o _bytecode_ em um arquivo [class](https://en.wikipedia.org/wiki/Java_class_file), possibilitando assim que este arquivo seja executado na JVM.

Qual a motivação para o projeto? Aprender temas relacionados a compiladores. E, para tal, nada melhor do que criar um, por mais simples que seja.

POJ é um projeto com escopo simples (na medida do possível) e de fácil entendimento. O foco é ser utilizado para estudos.

# Dependências

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

```
# Executar o POJ para gerar o assembly Java (arquivo jasm)
./bin/poj ./examples/hello_world

# Executar o JASM (java assembler) para gerar o executável Java (arquivo class)
jasm hello_world.jasm

# Executar o arquivo class com a JVM
java hello_world
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

# Licença

POJ utiliza a [licença Apache](LICENSE). 
