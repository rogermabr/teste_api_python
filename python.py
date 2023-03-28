import json
import requests

# VAI REALIZAR A LEITURA DO ARQUIVO EM JSON
with open('/mnt/c/Users/roger/Documents/cadastro.json', 'r') as f:
    payload = json.load(f)

# REALIZARÁ A VALIDAÇÃO DO ARQUIVO POR NOME E SOBRENOME
if 'firstName' not in payload or len(payload['firstName']) < 2:
    print('Campo firstName inválido')
    exit()
if 'lastName' not in payload or len(payload['lastName']) < 2:
    print('Campo lastName inválido')
    exit()
if 'cpf' not in payload or len(payload['cpf']) != 11:
    print('Campo cpf inválido')
    exit()

# CRIANDO O CAMPO COM NOME E SOBRENOME.
payload['fullName'] = payload['firstName'] + ' ' + payload['lastName']

# REMOVE OS CAMPOS NOME E SOBRENOME
del payload['firstName']
del payload['lastName']

# REMOVE PONTO E TRAÇO DO CPF.
payload['cpf'] = payload['cpf'].replace('.', '').replace('-', '')

# ENVIAR O PAYLOAD PARA A URL
# url = 'url_do_webservice'
# response = requests.post(url, json=payload)
# print(response.status_code)