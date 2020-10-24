from fastapi import FastAPI

from purse.monobank.routers import monobank_router


app = FastAPI()
app.include_router(monobank_router)
