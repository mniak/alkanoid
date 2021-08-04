# Rotina de transações
Cada portador de cartão (cliente) possui uma conta com seus dados.
A cada operação realizada pelo cliente uma transação é criada e associada à sua respectiva conta.

Cada transação possui um tipo (compra a vista, compra parcelada, saque ou pagamento), um valor e uma data de criação.

Transações de tipo _​compra e saque_  são registradas com _​valor negativo​_, enquanto transações de **​pagamento​** são registradas com **​valor positivo​**.

## Estrutura de dados
Segue abaixo uma estrutura de dados ​sugerida ​(fique a vontade para criar seu próprio modelo​)​:

### `Accounts`
| Account_ID | Document_Number |
|------------|-----------------|
| 1          | 12345678900     |

### `OperationsTypes`
| OperationType_ID | Description0     |
|------------------|------------------|
| 1                | COMPRA A VISTA   |
| 2                | COMPRA PARCELADA |
| 3                | SAQUE            |
| 4                | PAGAMENTO        |

### `Transactions`
| Transaction_ID | Account_ID | OperationType_ID | Amount | EventDate                   |
|----------------|------------|------------------|--------|-----------------------------|
| 1              | 1          | 1                | -50.0  | 2020-01-01T10:32:07.7199222 |
| 2              | 1          | 1                | -23.5  | 2020-01-01T10:48:12.2135875 |
| 3              | 1          | 1                | -18.7  | 2020-01-02T19:01:23.1458543 |
| 4              | 1          | 4                | 60.0   | 2020-01-05T09:34:18.5893223 | 

Na tabela de `​Transactions​`, a coluna `​Amount` guarda o valor da transação e a coluna
`EventDate​` guarda o momento em que ocorreu a transação.

## Endpoints
Desenvolva os endpoints abaixo considerando as regras de negócio mencionadas
anteriormente:

#### **`POST​`** `/accounts `​(criação de uma conta)
Request Body:
```
{
    "document_number": "12345678900"
}
```

#### **`GET`​** `/accounts/:accountId` ​(consulta de informações de uma conta)
Response Body:
```
{
    "account_id": 1,
    "document_number": "12345678900"
}
```

#### **`POST`** `/transactions` ​(criação de uma transação)
Request Body:
```
{
    "account_id": 1,
    "operation_type_id": 4,
    "amount": 123.45
}
```