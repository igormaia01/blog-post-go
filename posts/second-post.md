---
title: 'Estratégias de Evicção de Cache'
author: 'Igor'
date: '2025-09-24'
tags: ['pessoal', 'blog', 'cache', 'estratégias']
excerpt: 'Uma visão geral das estratégias de evicção de cache e suas implicações.'
slug: 'estrategias-eviccao-cache'
status: 'published'
---

# Estratégias de Evicção de Cache

Este artigo explora as principais estratégias de evicção de cache utilizadas em sistemas computacionais, suas vantagens, desvantagens e casos de uso específicos.

## 1. Least Recently Used (LRU)

Quando o cache atinge o limite de tamanho, o item que foi acessado há mais tempo é removido para liberar memória e armazenar o item mais recente, mantendo o controle da ordem dos itens mais acessados.

![LRU](/images/LRU.png 'LRU Cache')

### Vantagens

- **Implementação simples**: Fácil de entender a lógica
- **Uso eficiente do cache**: Em muitos cenários pode ser eficiente se o uso mais recente indicar acesso futuro
- **Adaptabilidade**: Muito útil para diferentes tipos de aplicações, incluindo banco de dados, cache web e sistemas de arquivos

### Desvantagens

- **Ordenação rígida**: Esta abordagem assume que a ordem de acesso dos itens reflete na utilidade futura de um item, e em alguns casos isso não é verdade
- **Problemas de cold start**: Quando um cache é inicialmente populado, LRU pode não ter performance otimizada pois requer dados históricos suficientes para tomar decisões informadas de evicção
- **Overhead de memória**: Implementar LRU frequentemente requer memória adicional para armazenar timestamps ou manter ordem de acesso, o que pode impactar o consumo geral de memória do sistema

### Casos de Uso

- **Cache web**: Para páginas e recursos frequentemente acessados
- **Banco de dados**: Cache de consultas e resultados
- **Sistemas de arquivos**: Cache de arquivos recentemente utilizados

## 2. Least Frequently Used (LFU)

As entradas acessadas com menor frequência são removidas primeiro. Baseia-se na ideia de que os itens mais frequentemente acessados serão relevantes no futuro, então quando estiver cheio, remove o item menos acessado em termos de número de vezes que cada item é acessado.

![LFU](/images/LFU.png 'LFU Cache')

### Vantagens

- **Adaptabilidade a padrões de acesso variados**: Eficaz em cenários onde alguns itens são acessados com pouca frequência mas ainda são relevantes
- **Otimizado para tendências de longo prazo**: Funciona bem quando itens são melhor caracterizados por sua frequência geral
- **Baixo overhead de memória**: LFU não precisa manter timestamps, economizando memória

### Desvantagens

- **Sensível ao acesso inicial**: LFU pode não ter boa performance no início da criação do cache, pois depende de padrões de acesso históricos
- **Dificuldade em lidar com mudanças nos padrões de acesso**: LFU pode ter dificuldades em cenários onde os padrões de acesso mudam frequentemente, mantendo itens que não são mais relevantes
- **Complexidade dos contadores de frequência**: Implementar contadores de frequência pode adicionar complexidade

### Casos de Uso

- **Cache de consultas de banco de dados**: Pode ser usado para cachear resultados frequentemente acessados
- **Tabelas de roteamento de rede**: Útil para cachear informações de roteamento em redes

## 3. First In First Out (FIFO)

O primeiro item que entra no cache será o primeiro item a ser removido, seguindo o princípio de fila.

![FIFO](/images/FIFO.png 'FIFO Cache')

### Vantagens

- **Implementação simples**: Escolha fácil para cenários onde a simplicidade é prioridade
- **Comportamento previsível**: Sabemos o que vai acontecer devido à ordem de entrada
- **Eficiente em memória**: Não precisa salvar metadados adicionais

### Desvantagens

- **Falta de adaptabilidade**: Devido à ordem de remoção dos itens, o FIFO não se adapta à importância dos itens

### Casos de Uso

- **Agendamento de tarefas**: No agendamento de tarefas, FIFO pode ser empregado para determinar a ordem em que processos ou tarefas são executados
- **Filas de mensagens**: FIFO garante que mensagens sejam tratadas na ordem que são recebidas em sistemas de filas de mensagem. Em aplicações que dependem de comunicação baseada em mensagens, isso é essencial para preservar a ordem dos processos
- **Cache para streaming**: Para algumas aplicações de streaming onde preservar a ordem dos dados é crucial, FIFO pode ser apropriado. Por exemplo, FIFO garante que frames sejam exibidos na ordem correta em um cache de streaming de vídeo

## 4. Random Replacement

Random Replacement é uma política de evicção de cache onde, quando o cache está cheio e um novo item precisa ser armazenado, um item existente escolhido aleatoriamente é removido para abrir espaço. Diferente de algumas políticas determinísticas como LRU (Least Recently Used) ou FIFO (First-In-First-Out), que têm critérios específicos para selecionar itens a serem removidos, Random Replacement simplesmente seleciona um item aleatoriamente.

![Random-Replacement](/images/random-replacement.png 'Random Replacement Cache')

### Exemplo

Considere um cache com três slots e os seguintes dados:

1. Item A
2. Item B
3. Item C

Agora, se o cache está cheio e um novo item, Item D, precisa ser armazenado, Random Replacement pode escolher remover o Item B, resultando em:

1. Item A
2. Item D
3. Item C

A seleção do Item B para remoção é inteiramente aleatória nesta política, tornando-a uma estratégia direta mas menos previsível comparada a outras.

### Vantagens

- **Simplicidade**: Random replacement é uma estratégia direta e fácil de implementar. Não requer rastreamento complexo ou análise de padrões de acesso
- **Evita vieses**: Como random replacement não depende de padrões de uso históricos, evita potenciais vieses que podem surgir em políticas mais determinísticas
- **Baixo overhead**: O algoritmo envolve overhead computacional mínimo, tornando-o eficiente em termos de requisitos de processamento

### Desvantagens

- **Performance subótima**: Random replacement pode levar a performance subótima do cache comparado a políticas mais sofisticadas. Não considera os padrões de uso reais ou a probabilidade de acessos futuros
- **Sem adaptabilidade**: Falta adaptabilidade a padrões de acesso em mudança. Outras políticas de evicção, como LRU ou LFU, consideram o comportamento histórico dos itens e se adaptam a padrões em evolução, potencialmente fornecendo melhor performance do cache ao longo do tempo
- **Possibilidade de taxas de hit ruins**: A natureza aleatória da evicção pode resultar em taxas de hit ruins, onde itens frequentemente acessados são removidos sem intenção, levando a mais cache misses

### Casos de Uso

1. **Ambientes de cache não-críticos**:

   - Em cenários onde o impacto de cache misses é mínimo ou onde o cache é empregado para propósitos não-críticos, como armazenamento temporário de dados não-essenciais, random replacement pode ser suficiente

2. **Simulação e testes**:

   - Em situações de teste e ambientes de simulação onde simplicidade e conveniência de uso são mais importantes que políticas de evicção complexas, random replacement é útil

3. **Sistemas com recursos limitados**:
   - Em ambientes com recursos limitados, onde recursos computacionais são escassos, o baixo overhead do random replacement pode ser vantajoso

## Conclusão

Em conclusão, políticas de evicção de cache desempenham um papel crucial no design de sistemas, impactando a eficiência e performance dos mecanismos de cache. A escolha de uma política de evicção depende das características específicas e requisitos do sistema. Enquanto políticas mais simples como Random Replacement oferecem facilidade de implementação e baixo overhead, estratégias mais sofisticadas como Least Recently Used (LRU) ou Least Frequently Used (LFU) levam em conta padrões de acesso históricos, levando a melhor adaptação a cargas de trabalho em mudança.

A escolha da estratégia ideal deve considerar fatores como:

- **Padrões de acesso esperados**
- **Recursos computacionais disponíveis**
- **Criticidade da aplicação**
- **Complexidade de implementação aceitável**

Compreender essas estratégias permite aos desenvolvedores tomar decisões informadas sobre qual política de evicção melhor se adequa às necessidades específicas de seus sistemas.
