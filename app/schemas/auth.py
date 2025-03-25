from pydantic import BaseModel, Field, EmailStr


class UserCreate(BaseModel):
    email: EmailStr = Field(..., min_length=1, max_length=100)
    username: str = Field(..., min_length=1, max_length=100)
    password: str = Field(
        ...,
        min_length=8, 
        max_length=100
        pattern=r"^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$",
        description="Password must contain at least one uppercase, one lowercase, one digit, and one special character."
    )

class UserLogin(BaseModel):
    email: EmailStr = Field(..., min_length=1, max_length=100)
    password: str = Field(..., min_length=8, max_length=100)
