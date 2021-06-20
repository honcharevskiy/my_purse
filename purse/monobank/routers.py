from logging import getLogger

from fastapi import APIRouter

from purse.monobank.models import WebHookItem


monobank_router = APIRouter()

logger = getLogger(__name__)


@monobank_router.post('/monobank')
def monobank_handler(item: WebHookItem):
    """Monobank WebHook handler."""
    logger.info('Receive: %s', item)
    return {'Hello': 'word'}
