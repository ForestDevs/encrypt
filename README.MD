```bash
 docker build -t app_encrypt_decrypt ./

 docker run --rm -d --name app_encrypt_decrypt \
  -e APP_LOG_LEVEL=2 \
  -e APP_PORT=8876 \
  -e APP_PROXY_PORT=8866 \
  -e CRYPTO_KEY=thisis32bitlongpassphraseimusing \
  -p 8081:8081 app_encrypt_decrypt
```