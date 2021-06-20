from unittest import TestCase

from starlette.testclient import TestClient

from purse.monobank import routers
from purse.monobank.models import WebHookItem, WebHookData, Transaction


class MonobankTest(TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        cls.client = TestClient(routers.monobank_router)

    def test_monobank_endpoint(self):
        fake_transaction = {
            "id": "ZuHWzqkKGVo=",
            "time": 1554466347,
            "description": "Покупка щастя",
            "mcc": 7997,
            "hold": False,
            "amount": -95000,
            "operationAmount": -95000,
            "currencyCode": 980,
            "commissionRate": 0,
            "cashbackAmount": 19000,
            "balance": 10050000
        }
        response = self.client.post(
            'monobank',
            json=WebHookItem(
                type='',
                data=WebHookData(
                    account='',
                    statementItem=Transaction.construct(**fake_transaction),
                ),
            ).dict(),
        )
        assert response.status_code == 200
