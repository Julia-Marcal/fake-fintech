## Serviços

O projeto utilizará **microsserviços** para consumir várias **APIs de investimentos e economia**. Cada microsserviço será responsável por tarefas específicas, como buscar dados das APIs, processá-los e entregá-los ao frontend ou a outros serviços.

## APIs que Serão Utilizadas

As seguintes APIs serão integradas aos microsserviços para fornecer dados para o projeto:

### 1. **B3 API (Bolsa de Valores Brasileira)**  
   - **Descrição**: Fornece acesso a dados da bolsa de valores brasileira, incluindo dados de mercado em tempo real e históricos, ações corporativas e instrumentos financeiros.  
   - **Casos de Uso**: Plataformas de negociação de ações, ferramentas de gestão de portfólio e aplicações de análise de mercado.  
   - **Funcionalidades**: Dados em tempo real, dados históricos, ações corporativas e detalhes de instrumentos financeiros.  
   - **Documentação**: [Documentação da B3 API](https://developers.b3.com.br/apis)

### 2. **Dados de Mercado API**  
   - **Descrição**: Oferece dados financeiros e econômicos de empresas brasileiras, incluindo indicadores financeiros, tendências de mercado e fundamentos das empresas.  
   - **Casos de Uso**: Análise financeira, pesquisa de investimentos e previsão econômica.  
   - **Funcionalidades**: Indicadores financeiros, fundamentos das empresas e tendências de mercado.  
   - **Documentação**: [Documentação da Dados de Mercado API](https://www.dadosdemercado.com.br/api/docs/empresas/indicadores-financeiros)

### 3. **Bavest API**  
   - **Descrição**: Fornece dados financeiros, incluindo preços de ações, fundamentos das empresas e dados alternativos. Projetada para aplicações fintech e plataformas de investimento.  
   - **Casos de Uso**: Plataformas de investimento, robôs assessores e ferramentas de análise financeira.  
   - **Funcionalidades**: Dados de ações em tempo real e históricos, fundamentos das empresas e dados alternativos.  
   - **Documentação**: [Documentação da Bavest API](https://docs.bavest.co/)

### 4. **Trading Economics API**  
   - **Descrição**: Fornece acesso a uma ampla gama de indicadores econômicos, incluindo PIB, inflação, desemprego e dados comerciais de mais de 196 países.  
   - **Casos de Uso**: Pesquisa econômica, análise de políticas e previsão de mercado.  
   - **Funcionalidades**: Dados econômicos em tempo real e históricos, cobertura global e solicitações personalizáveis.  
   - **Documentação**: [Documentação da Trading Economics API](https://docs.tradingeconomics.com/indicadores/snapshot/)
