from pydantic import BaseModel, Field, EmailStr, validator


class UserCreate(BaseModel):
    email: EmailStr = Field(..., min_length=1, max_length=100)
    username: str = Field(..., min_length=1, max_length=100)
    password: str = Field(
        ...,
        min_length=8, 
        max_length=100,
        description="Password must contain at least one uppercase, one lowercase, one digit, and one special character."
    )
    
    @validator("password")
    def validate_password(cls, value):
        if not any(c.islower() for c in value):
            raise ValueError("Password must contain at least one lowercase letter.")
        if not any(c.isupper() for c in value):
            raise ValueError("Password must contain at least one uppercase letter.")
        if not any(c.isdigit() for c in value):
            raise ValueError("Password must contain at least one digit.")
        if not any(c in "@$!%*?&" for c in value):
            raise ValueError("Password must contain at least one special character (@$!%*?&).")
        return value

class UserLogin(BaseModel):
    email: EmailStr = Field(..., min_length=1, max_length=100)
    password: str = Field(..., min_length=8, max_length=100)
