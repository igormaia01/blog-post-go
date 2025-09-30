# Admin Dashboard e Sistema de Autenticação

## Visão Geral

O sistema agora possui um painel administrativo completo com autenticação e métricas de posts.

## Funcionalidades Implementadas

### 🔐 Sistema de Autenticação

- Login seguro com usuário e senha armazenados no arquivo `.env`
- Sessões gerenciadas com tokens seguros
- Expiração automática de sessões após 24 horas
- Limpeza automática de sessões expiradas a cada hora
- Middleware de autenticação protegendo rotas administrativas

### 📊 Dashboard com Métricas

O dashboard exibe:

- **Total de Posts**: Quantidade total de posts no blog
- **Posts Publicados**: Quantidade de posts com status "published"
- **Posts em Rascunho**: Quantidade de posts com status "draft"
- **Total de Visualizações**: Soma de todas as visualizações de posts
- **Total de Compartilhamentos**: Soma de todos os compartilhamentos
- **Visualizações de Hoje**: Visualizações registradas hoje
- **Compartilhamentos de Hoje**: Compartilhamentos registrados hoje

### 📈 Métricas por Post

Cada post exibe individualmente:

- **Visualizações**: Incrementadas automaticamente ao acessar o post
- **Compartilhamentos Totais**: Soma de todos os compartilhamentos
- **Compartilhamentos por Plataforma**:
  - Facebook
  - Twitter
  - LinkedIn

### 🔄 Rastreamento de Compartilhamentos

- Rastreamento automático quando o usuário clica nos botões de compartilhamento
- API endpoint `/api/track-share` para registrar compartilhamentos
- Suporte para múltiplas plataformas sociais

## Como Usar

### 1. Configurar Credenciais do Admin

Edite o arquivo `configs/app.env` para definir suas credenciais:

```env
ADMIN_USERNAME=seu_usuario
ADMIN_PASSWORD=sua_senha_segura
ADMIN_SECRET=chave-secreta-aleatoria-para-tokens
```

### 2. Acessar o Painel Administrativo

1. Inicie o servidor:

```bash
go run cmd/server/main.go
```

2. Acesse o painel de login:

```
http://localhost:3100/admin/login
```

3. Faça login com as credenciais configuradas no `.env`

4. Após o login, você será redirecionado para o dashboard:

```
http://localhost:3100/admin
```

### 3. Visualizar Métricas

No dashboard, você verá:

- Cards com estatísticas gerais no topo
- Lista de posts com métricas detalhadas de cada um
- Status de publicação de cada post

### 4. Rastreamento Automático

As métricas são rastreadas automaticamente:

- **Visualizações**: Incrementadas quando um usuário acessa um post
- **Compartilhamentos**: Registrados quando o usuário clica em um botão de compartilhamento

### 5. Fazer Logout

Clique no botão "Logout" no cabeçalho do admin para encerrar a sessão.

## Estrutura dos Serviços

### AuthService (`internal/services/auth_service.go`)

- Gerencia autenticação de usuários
- Cria e valida tokens de sessão
- Limpa sessões expiradas automaticamente

### MetricsService (`internal/services/metrics_service.go`)

- Rastreia visualizações de posts
- Rastreia compartilhamentos por plataforma
- Fornece estatísticas agregadas
- Thread-safe para acesso concorrente

### Middleware de Autenticação (`internal/middleware/auth.go`)

- Protege rotas administrativas
- Valida tokens de sessão
- Redireciona para login se não autenticado

## Rotas da API

### Públicas

- `GET /admin/login` - Página de login
- `POST /admin/login` - Submissão do formulário de login

### Protegidas (requerem autenticação)

- `GET /admin` - Dashboard principal
- `GET /admin/logout` - Logout
- `GET /admin/posts` - Lista de posts
- `GET /admin/posts/new` - Criar novo post
- `GET /admin/posts/:id/edit` - Editar post
- `POST /admin/posts` - Salvar novo post
- `POST /admin/posts/:id` - Atualizar post
- `DELETE /admin/posts/:id` - Deletar post
- `GET /admin/tags` - Lista de tags
- `GET /admin/settings` - Configurações

### API de Métricas

- `POST /api/track-share` - Rastrear compartilhamento
  ```json
  {
    "post_id": 1,
    "platform": "twitter"
  }
  ```

## Segurança

### Boas Práticas Implementadas

- ✅ Senhas armazenadas em variáveis de ambiente
- ✅ Tokens de sessão gerados com criptografia segura
- ✅ Cookies HTTPOnly para prevenir XSS
- ✅ Expiração automática de sessões
- ✅ Middleware de autenticação em todas as rotas administrativas

### Recomendações para Produção

1. **Usar HTTPS**: Sempre use HTTPS em produção
2. **Senhas Fortes**: Use senhas complexas e únicas
3. **Secret Key**: Gere uma chave secreta longa e aleatória
4. **Rate Limiting**: Implemente rate limiting no login
5. **2FA**: Considere adicionar autenticação de dois fatores
6. **Logs**: Mantenha logs de tentativas de login

## Modelos de Dados

### PostMetrics

```go
type PostMetrics struct {
    PostID          int
    ViewCount       int
    ShareCount      int
    FacebookShares  int
    TwitterShares   int
    LinkedInShares  int
    LastViewedAt    time.Time
    LastSharedAt    time.Time
}
```

### DashboardStats

```go
type DashboardStats struct {
    TotalPosts      int
    PublishedPosts  int
    DraftPosts      int
    ArchivedPosts   int
    TotalViews      int
    TotalShares     int
    TodayViews      int
    TodayShares     int
}
```

### Session

```go
type Session struct {
    Token     string
    Username  string
    CreatedAt time.Time
    ExpiresAt time.Time
}
```

## Troubleshooting

### Não consigo fazer login

1. Verifique se as credenciais no `.env` estão corretas
2. Verifique se o servidor está lendo o arquivo `.env`
3. Verifique os logs do servidor para erros

### Métricas não estão sendo registradas

1. Verifique se o JavaScript está habilitado no navegador
2. Abra o console do navegador para ver erros
3. Verifique se a rota `/api/track-share` está acessível

### Sessão expira muito rápido

- As sessões expiram após 24 horas
- Para alterar, modifique a duração em `auth_service.go`:

```go
ExpiresAt: time.Now().Add(24 * time.Hour), // Altere aqui
```

## Próximas Melhorias Sugeridas

- [ ] Persistir métricas em banco de dados
- [ ] Adicionar gráficos de visualizações ao longo do tempo
- [ ] Exportar relatórios de métricas
- [ ] Adicionar mais plataformas de compartilhamento
- [ ] Sistema de notificações por email
- [ ] Backup automático de dados
- [ ] Interface para gerenciar múltiplos usuários admin
