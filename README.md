<h1>Go Expert</h1>
<h5>
Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
</h5>

<p><i><b>Você precisará nos entregar dois sistemas em Go:</b></i></p>
<li><i>client.go</i></li>
<li><i>server.go</i></li>

<h5>Os requisitos para cumprir este desafio são:</h5>
<p>O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.</p>
<p>O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: <a href="https://economia.awesomeapi.com.br/json/last/USD-BRL">https://economia.awesomeapi.com.br/json/last/USD-BRL</a> e em seguida deverá retornar no formato JSON o resultado para o cliente.</p>
<p>Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.</p>
<p>O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.</p>
<p>Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.</p>
<p>O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}</p>
<p>O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.</p>