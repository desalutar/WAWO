FROM python:3.12-alpine

WORKDIR /app/frontend

RUN apk add --no-cache build-base

COPY . .

RUN pip install --no-cache-dir flask grpcio grpcio-tools

ENV PYTHONPATH=/app/frontend

EXPOSE 5500

CMD ["python", "app.py"]


ENV PYTHONPATH=/app/frontend
