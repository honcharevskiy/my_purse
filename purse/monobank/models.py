from typing import TypedDict

from pydantic import BaseModel


class Transaction(BaseModel):
    """Monobank transaction fields."""
    id: str
    time: int
    description: str
    mcc: int
    hold: bool
    amount: int
    operationAmount: int
    currencyCode: int
    commissionRate: int
    cashbackAmount: int
    balance: int


class WebHookData(BaseModel):
    account: str
    statementItem: Transaction


class WebHookItem(BaseModel):
    """Description of monobank fields from docs.

    Docs url: https://api.monobank.ua/docs/
    """
    type: str
    data: WebHookData
