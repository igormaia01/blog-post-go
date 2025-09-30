#!/bin/bash

# Script de teste para o sistema de autenticação e métricas

echo "==================================="
echo "🧪 Teste do Sistema de Dashboard"
echo "==================================="
echo ""

BASE_URL="http://localhost:3100"

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}1. Testando endpoint público (home)${NC}"
curl -s -o /dev/null -w "Status: %{http_code}\n" $BASE_URL/
echo ""

echo -e "${YELLOW}2. Testando acesso ao admin sem autenticação (deve redirecionar)${NC}"
curl -s -I $BASE_URL/admin | grep -i "location\|302\|301"
echo ""

echo -e "${YELLOW}3. Tentando fazer login${NC}"
echo "Enviando credenciais: admin / admin123"
COOKIE=$(curl -s -c - -X POST $BASE_URL/admin/login \
  -d "username=admin&password=admin123" \
  | grep "admin_session" | awk '{print $7}')

if [ ! -z "$COOKIE" ]; then
  echo -e "${GREEN}✓ Login bem-sucedido! Token recebido.${NC}"
  echo "Token (primeiros 20 chars): ${COOKIE:0:20}..."
else
  echo -e "${RED}✗ Falha no login${NC}"
  exit 1
fi
echo ""

echo -e "${YELLOW}4. Acessando dashboard com autenticação${NC}"
curl -s -b "admin_session=$COOKIE" $BASE_URL/admin \
  | grep -o "<title>.*</title>" || echo "Dashboard acessível"
echo ""

echo -e "${YELLOW}5. Simulando visualização de post (ID: 1)${NC}"
curl -s -o /dev/null -w "Status: %{http_code}\n" $BASE_URL/post/1
echo ""

echo -e "${YELLOW}6. Simulando compartilhamento no Twitter (Post ID: 1)${NC}"
curl -s -X POST $BASE_URL/api/track-share \
  -H "Content-Type: application/json" \
  -d '{"post_id":1,"platform":"twitter"}' | jq '.'
echo ""

echo -e "${YELLOW}7. Simulando compartilhamento no Facebook (Post ID: 1)${NC}"
curl -s -X POST $BASE_URL/api/track-share \
  -H "Content-Type: application/json" \
  -d '{"post_id":1,"platform":"facebook"}' | jq '.'
echo ""

echo -e "${YELLOW}8. Simulando mais visualizações${NC}"
for i in {1..5}; do
  echo -n "View $i... "
  curl -s -o /dev/null -w "%{http_code}\n" $BASE_URL/post/1
done
echo ""

echo -e "${YELLOW}9. Testando logout${NC}"
curl -s -b "admin_session=$COOKIE" $BASE_URL/admin/logout \
  -I | grep -i "location\|302\|301"
echo ""

echo -e "${GREEN}==================================="
echo "✓ Testes concluídos!"
echo "===================================${NC}"
echo ""
echo "Acesse o dashboard em: $BASE_URL/admin/login"
echo "Credenciais: admin / admin123"
echo ""
echo "Os posts devem mostrar:"
echo "  - 6 visualizações no Post ID 1"
echo "  - 1 compartilhamento no Twitter"
echo "  - 1 compartilhamento no Facebook"
