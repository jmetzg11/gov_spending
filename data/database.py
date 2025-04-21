from sqlalchemy import create_engine, Column, Integer, String, Float
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker


Base = declarative_base()

class AgencyBudget(Base):
    __tablename__ = 'agency_budget'
    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    agency = Column(String(255))
    budget = Column(Float)


class ForeignAid(Base):
    __tablename__ = 'foreign_aid'
    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    country = Column(String(255))
    amount = Column(Float)
    year = Column(Integer)
    lat = Column(Float)
    lng = Column(Float)


class FunctionSpending(Base):
    __tablename__ = 'function_spending'
    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    year = Column(Integer)
    name = Column(String(255))
    amount = Column(Float)

engine = create_engine('sqlite://data.db')

Base.metadata.create_all(bind=engine)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
