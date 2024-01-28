Resumo do Stress Tester:

O Stress Tester é uma ferramenta desenvolvida em Go que tem como objetivo realizar testes de estresse em uma aplicação web específica. O usuário fornece a URL alvo, o número desejado de requisições e o nível de concorrência desejado através de flags de linha de comando. As requisições são enviadas em paralelo para simular uma carga intensiva e verificar como a aplicação se comporta sob pressão.

Explicação Detalhada:

O Stress Tester foi projetado para permitir que os usuários avaliem o desempenho e a robustez de uma aplicação web ao simular uma carga significativa de tráfego. Ele é implementado em Go e utiliza concorrência para enviar múltiplas requisições simultaneamente.

A interface de linha de comando do Stress Tester é configurada através de três flags principais:

-url (obrigatória): Define a URL da aplicação que será testada. É a única flag obrigatória e indica para onde as requisições serão enviadas.

-requests (opcional): Especifica o número total de requisições que o teste de estresse realizará. Se não for fornecido, o padrão é 1.

-concurrency (opcional): Define o nível de concorrência, ou seja, quantas requisições serão executadas simultaneamente. Se não for fornecido, o padrão é 1.

Exemplo de uso:

bash
Copy code
go run stress_tester.go -url https://example.com -requests 1000 -concurrency 50
No exemplo acima, o Stress Tester enviará 1000 requisições para a URL https://example.com com uma concorrência de 50 requisições simultâneas.

A ferramenta utiliza a biblioteca nativa de concorrência do Go para criar goroutines que realizam as requisições. Após o término do teste, são exibidas estatísticas relevantes, como o tempo total decorrido, média de tempo por requisição e status de cada requisição.

O Stress Tester fornece aos desenvolvedores uma maneira rápida e eficaz de avaliar o desempenho de suas aplicações sob diferentes condições de carga, identificando possíveis gargalos e melhorando a confiabilidade do sistema em produção.
