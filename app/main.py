from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from models import Base, engine
from routers import auth_router

Base.metadata.create_all(bind=engine)

app = FastAPI()

# CORS settings
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(auth_router, prefix="/auth", tags=["auth"])

@app.get("/ping")
def ping():
    return {"status": "OK"}