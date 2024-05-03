program function_call_with_two_params_string;

function myfunction(msg, name: string) : string;
begin
    myfunction := msg + ' ' + name + '!';
end;

begin
    writeln('Hello from main!');
    writeln(myfunction('Hello', 'Alex'));
end.